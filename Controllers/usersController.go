package Controllers

import (
	"time"

	"github.com/KaisAlHusrom/RestApi/Config"
	"github.com/KaisAlHusrom/RestApi/Models"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

// inser User
func AddUser(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["user_name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User Name is required",
		})
	}

	if data["email"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "email is required",
		})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "password is required",
		})
	}

	user := Models.User{
		UserName:  data["user_name"],
		Email:     data["email"],
		Password:  data["password"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	Config.DB.Create(&user)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User Added Successfully",
		"data":    user,
	})
}

// GET users info only
func ShowUsers(c *fiber.Ctx) error {
	var users []Models.User
	// limit, _ := strconv.Atoi(c.Query("limit"))
	// skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	Config.DB.Select("user_id, user_name, email, password, created_at, updated_at").Find(&users).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Users List Api",
		"data":    users,
	})
}

// Get All users with all their products
func GetUsersWithProducts(c *fiber.Ctx) error {
	var users []Models.User
	// limit, _ := strconv.Atoi(c.Query("limit"))
	// skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	Config.DB.Select("*").Preload("Products").Find(&users).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Users List Api with products",
		"data":    users,
	})
}

// Get one user info only by id
func GetUserById(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	var user Models.User

	Config.DB.Select("user_id, user_name, email, updated_at, created_at").Where("user_id = ?", userId).First(&user)

	//if ther is not a user with given Id
	if user.UserId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "There is no user with id: " + userId,
		})
	}

	userData := make(map[string]interface{})
	userData["id"] = user.UserId
	userData["user_name"] = user.UserName
	userData["email"] = user.Email
	userData["created_at"] = user.CreatedAt
	userData["updated_at"] = user.UpdatedAt

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User with id: " + userId,
		"data":    userData,
	})
}

// Get user info with his products by id
func GetUserProductsById(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	var user Models.User

	// Get the user by ID
	err := Config.DB.Preload("Products").First(&user, userId).Error
	if err != nil {
		// Handle the error if the user is not found
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "There is no user with id: " + userId,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get user with his products successfully with id: " + userId,
		"data":    user,
	})
}

// Delete users with belong products
func DeleteUserById(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	var user Models.User
	var product Models.Product

	Config.DB.Where("user_id=?", userId).First(&user)
	if user.UserId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "There is no user with id: " + userId,
		})
	}

	err := Config.DB.Transaction(func(tx *gorm.DB) error {
		// Delete the products associated with the user
		err := tx.Delete(&product, "user_id = ?", userId).Error
		if err != nil {
			// Handle the error if products deletion fails
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Delete Products Failed",
			})
		}

		// Delete the user
		if err := tx.Delete(&user, userId).Error; err != nil {
			// Handle the error if user deletion fails
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Delete User Failed",
			})
		}

		return nil
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Delete Failed",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User with id: " + userId + " with his products deleted successfully",
	})
}

// update users info
func UpdateUserById(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	var user Models.User

	Config.DB.Find(&user, "user_id=?", userId)
	if user.UserId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "There is no user with id: " + userId,
		})
	}

	var updatedUser Models.User

	err := c.BodyParser(&updatedUser)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	if updatedUser.UserName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User Name is required",
		})
	}

	if updatedUser.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "email is required",
		})
	}

	if updatedUser.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "password is required",
		})
	}

	user.UserName = updatedUser.UserName
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	user.UpdatedAt = time.Time{}

	Config.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User with id: " + userId + " updated successfully",
		"data":    user,
	})
}
