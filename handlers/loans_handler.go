package handlers

import (
	"biblioapp/database"
	"biblioapp/models"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetLoans(c echo.Context) error {
	var loans []models.Loan
	if err := database.DB.Preload("Book").Preload("User").Find(&loans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, loan := range loans {
		fmt.Printf("Loan ID: %d, Book: %+v, User: %+v\n", loan.ID, loan.Book, loan.User)
	}
	return c.JSON(http.StatusOK, loans)
}

func AddLoan(c echo.Context) error {
	var loan models.Loan
	if err := c.Bind(&loan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var existingLoan models.Loan
	var user models.User
	var book models.Book

	// Check if user exists
	if err := database.DB.First(&user, loan.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, "User not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check if book exists
	if err := database.DB.First(&book, loan.BookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, "Book not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check if a loan already exists
	if err := database.DB.Where("book_id = ? AND user_id = ? AND return_date = ?", loan.BookID, loan.UserID, loan.ReturnDate).First(&existingLoan).Error; err == nil {
		return c.JSON(http.StatusBadRequest, "A loan already exists for the given book, user, and return date")
	} else if err != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := database.DB.Create(&loan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, loan)
}

func UpdateLoan(c echo.Context) error {
	var loan models.Loan
	if err := c.Bind(&loan); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	if err := database.DB.Save(&loan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c echo.Context) error {
	loanId := c.Param("id")
	if err := database.DB.Where("id = ?", loanId).Delete(&models.Loan{}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Loan deleted")
}

func GetUsersLoans(c echo.Context) error {
	userID := c.Param("id")
	var loans []models.Loan
	if err := database.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, loans)
}
