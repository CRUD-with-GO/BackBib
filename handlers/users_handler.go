package handlers

import (
	"biblioapp/database"
	"biblioapp/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(c echo.Context) error{
	var users []models.User
	database.DB.Preload("Loan").Find(&users)
	return c.JSON(http.StatusOK, users)
}

func AddUser(c echo.Context) error{
	var user models.User
	c.Bind(&user)
	var existingUser models.User
	database.DB.Where("id = ?",user.ID).Find( &existingUser)
	if existingUser.ID != 0 {
		return c.JSON(http.StatusBadRequest, "User already exists")
		}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated,user)
}

func UpdateUser(c echo.Context) error {
    userId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid user ID")
    }

    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request payload")
    }

    user.ID = int(userId)
    if err := database.DB.Save(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error{
	var user models.User
	userId := c.Param("id")
	database.DB.Where("id = ?", userId).Delete(&user)
	return c.JSON(http.StatusOK, GetUsers(c))
}


func GetUser(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    var user models.User
    if err := database.DB.Preload("Loans").Preload("Loans.Book").First(&user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.JSON(http.StatusNotFound, "User not found")
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, user)
}