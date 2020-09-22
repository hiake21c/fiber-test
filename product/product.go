package product

import (
	"fiber-test/database"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint64
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

	uuid := uuid.New()

	var product Product
	product.Code = uuid.String()
	product.Price = 10

	db.Save(&product)

	return c.JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var param = Product{}
	c.BodyParser(&param)

	id := c.Params("id")

	db := database.Db
	var product Product
	db.Find(&product, id)

	//product.Price, _ = strconv.ParseUint(price, 10 , 32)
	product.Price = param.Price
	db.Save(&product)
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
