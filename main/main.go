package main

import (
	"GoProject/bookManagement/handlers"
	"GoProject/bookManagement/models"
	_ "GoProject/bookManagement/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	models.InitDB("sqlite.db")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/books", handlers.GetBooks)
	e.GET("/books/:id", handlers.GetBook)
	e.POST("/books", handlers.CreateBook)
	e.PUT("/books/:id", handlers.UpdateBook)
	e.DELETE("/books/:id", handlers.DeleteBook)

	// Старт на сървъра
	log.Fatal(e.Start(":8080"))
}
