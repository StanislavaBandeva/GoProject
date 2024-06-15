package handlers

import (
	"GoProject/bookManagement/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetBooks(c echo.Context) error {
	books, err := models.GetBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to fetch books"})
	}
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid book ID"})
	}
	book, err := models.GetBook(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Book not found"})
	}
	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	err := models.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to create book"})
	}
	return c.JSON(http.StatusCreated, book)
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid book ID"})
	}
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	err = models.UpdateBook(id, book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to update book"})
	}
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid book ID"})
	}
	err = models.DeleteBook(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to delete book"})
	}
	return c.NoContent(http.StatusNoContent)
}
