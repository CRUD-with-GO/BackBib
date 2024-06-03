package handlers

import (
	"biblioapp/database"
	"biblioapp/models"

	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error{
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func AddBooks(c echo.Context) error{
	var book models.Book
	c.Bind(&book)
	var existingBook  models.Book
	database.DB.Where("id = ?",book.ID).Find( &existingBook)
	if existingBook.ID!=0 {	
		return c.JSON(http.StatusBadRequest, "Book already exists")
		}
	database.DB.Create(&book)
	return c.JSON(http.StatusCreated, book)
}

func UpdateBook(c echo.Context) error {
    
    bookId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid book ID")
    }

    var book models.Book
    if err := c.Bind(&book); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid request payload")
    }

    book.ID = int(bookId)
    result := database.DB.Save(&book)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error.Error())
    }


    return c.JSON(http.StatusOK, book)
}


func DeleteBook(c echo.Context) error {
    bookId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid book ID")
    }


    result := database.DB.Delete(&models.Book{}, bookId)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error.Error())
    }


    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "Book not found")
    }

    return c.NoContent(http.StatusNoContent)
}

func GetBook(c echo.Context) error {
	isbn := c.Param("id")
	var book models.Book
	database.DB.Where("isbn = ?",isbn).Find(book)
	return c.JSON(http.StatusOK,book)

}