package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/n-guitar/go_restapi_fiber_sqlite/pkg/book"
	"github.com/n-guitar/go_restapi_fiber_sqlite/pkg/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func apiTop(c *fiber.Ctx) error {
	return c.SendString("apiTop")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", apiTop)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() *gorm.DB {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	// Migrate the schema
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
	return database.DBConn
}

func main() {
	app := fiber.New()
	db := initDatabase()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	app.Get("/", apiTop)
	setupRoutes(app)

	app.Listen(":3000")

	// Close
	defer sqlDB.Close()
}
