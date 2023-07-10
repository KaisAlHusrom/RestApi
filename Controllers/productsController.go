package Controllers

import (
	"fmt"
	"time"

	"github.com/KaisAlHusrom/RestApi/Config"
	"github.com/KaisAlHusrom/RestApi/Models"
	"github.com/gofiber/fiber/v2"
)

// insert Product
func AddProduct(c *fiber.Ctx) error {
	var product Models.Product

	err := c.BodyParser(&product)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	var user Models.User

	if err := Config.DB.Where("user_id = ?", product.UserID).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find user",
		})
	}

	if product.ProductName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Product Name is required",
		})
	}

	if product.Price == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Price is required",
		})
	}

	if product.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User is not found",
		})
	}

	Config.DB.Create(&product)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product Added Successfully",
		"data":    product,
	})
}

// Get Products
func GetProducts(c *fiber.Ctx) error {
	var products []Models.Product
	var count int64

	Config.DB.Select("*").Find(&products).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Products List Api",
		"data":    products,
	})
}

// Get Product with id
func GetProductById(c *fiber.Ctx) error {
	productId := c.Params("product_id")

	var product Models.Product
	Config.DB.Find(&product, "product_id=?", productId)
	if product.ProductId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("There is no product with id: %v", productId),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Get product successfully",
		"data":    product,
	})
}

// delete product by id
func DeleteProductById(c *fiber.Ctx) error {
	productId := c.Params("product_id")

	var product Models.Product
	if err := Config.DB.Where("product_id = ?", productId).First(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find Product",
		})
	}

	err := Config.DB.Delete(&product, "product_id = ?", productId).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Delete Product Faild",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product with id: " + productId + " deleted successfully",
	})
}

// update product by id
func UpdateProductById(c *fiber.Ctx) error {
	productId := c.Params("product_id")

	var product Models.Product

	//check if product not null
	Config.DB.Find(&product, "product_id=?", productId)
	if product.ProductId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("There is no product with id: %v", productId),
		})
	}

	var updatedProduct Models.Product

	err := c.BodyParser(&updatedProduct)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	var user Models.User

	if err := Config.DB.Where("user_id = ?", updatedProduct.UserID).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to find user",
		})
	}

	if updatedProduct.ProductName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Product Name is required",
		})
	}

	if updatedProduct.Price == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Price is required",
		})
	}

	if updatedProduct.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User is not found",
		})
	}
	product.ProductName = updatedProduct.ProductName
	product.Price = updatedProduct.Price
	product.UserID = updatedProduct.UserID
	product.UpdatedAt = time.Time{}

	Config.DB.Save(&product)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Product Updated Successfully",
		"data":    product,
	})
}
