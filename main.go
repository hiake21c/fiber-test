package main

import (
	"fiber-test/database"
	"fiber-test/product"
	"fmt"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("hello, world!, Welcome home!")
}

func initDatabase() {
	var err error
	database.Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	//Migrate the schema
	database.Db.AutoMigrate(&product.Product{})
	fmt.Println("Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", homeHandler)
	app.Get("/api/v1/product", product.GetAllProduct)
	app.Get("/api/v1/product/:id", product.GetProduct)
	app.Post("/api/v1/product", product.SaveProduct)
	app.Delete("/api/v1/product/:id", product.DeleteProduct)
	app.Put("/api/v1/product/:id", product.SaveProduct)
}

func main() {

	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(":3000")

}
