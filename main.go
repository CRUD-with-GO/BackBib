package main

import (
	"biblioapp/database"
	"biblioapp/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))


	database.InitDB();
	e.GET("/books",handlers.GetBooks)
	e.GET("/book/:id",handlers.GetBook)
	e.POST("/books", handlers.AddBooks)
	e.PUT("/updateBook/:id",handlers.UpdateBook)
	e.DELETE("/deleteBook/:id",handlers.DeleteBook)


	e.GET("/users",handlers.GetUsers)
	e.GET("/user/:id",handlers.GetUser)
	e.POST("/users", handlers.AddUser)
	e.PUT("/updateUser/:id",handlers.UpdateUser)
	e.DELETE("/deleteUser/:id",handlers.DeleteUser)


	e.GET("/loan",handlers.GetLoans)
	e.GET("/loans/:id",handlers.GetUsersLoans)
	e.POST("/loan", handlers.AddLoan)
	e.PUT("/updateLoan/:id",handlers.UpdateLoan)
	e.DELETE("/deleteLoan/:id",handlers.DeleteLoan)


	e.Logger.Fatal(e.Start(":8080"))
}