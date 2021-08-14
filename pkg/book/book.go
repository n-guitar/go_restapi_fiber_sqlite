package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n-guitar/go_restapi_fiber_sqlite/pkg/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	// return c.SendString("All Books")
	db := database.DBConn
	var books []Book
	db.Find(&books)
	// books 配列をすばやく簡単に JSON 文字列にシリアル化して応答
	return c.JSON(books)
}

// ToDo 本がないときの処理
func GetBook(c *fiber.Ctx) error {
	// return c.SendString("Single Book")
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book Book
	book.Title = "1984"
	book.Author = "George Orwell"
	book.Rating = 5
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	// return c.SendString("Delete Book")
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&book)
	return c.SendString("Book Seccessfully Deleted")
}
