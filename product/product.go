package product

import (
	"fiber-test/database"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/utils"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func GetAllProduct(c *fiber.Ctx) error {

	db := database.Db
	var products []Product
	db.Find(&products)

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.Db
	var product Product
	db.Find(&product, id)
	return c.JSON(product)
}

func SaveProduct(c *fiber.Ctx) error {
	db := database.Db
	var product Product
	product.Code = "CODE" + utils.UUID()
	product.Price = 10
	db.Create(&product)
	return c.JSON(&product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Db

	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		return c.Status(500).SendString("No Book Found with ID")
	}
	db.Delete(&product)
	return c.SendString("product Successfully deleted")
}
