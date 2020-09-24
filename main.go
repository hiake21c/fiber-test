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

func setupRoutes(router *fiber.App) {

	v1 := router.Group("/api/v1")
	{
		v1.Get("/product", product.GetAllProduct)
		v1.Get("/product/:id", product.GetProduct)
		v1.Post("/product", product.SaveProduct)
		v1.Delete("/product/:id", product.DeleteProduct)
		v1.Put("/product/:id", product.UpdateProduct)
	}

	router.Get("/", homeHandler)
}

func main() {

	router := fiber.New()
	initDatabase()
	setupRoutes(router)
	router.Listen(":3000")

}
