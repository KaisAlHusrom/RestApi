package Routes

import (
	"github.com/KaisAlHusrom/RestApi/Controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	//Users Routers
	users := app.Group("/users")
	users.Post("/addUser", Controllers.AddUser)
	users.Get("/getUsers", Controllers.ShowUsers)
	users.Get("/getUsersWithProducts", Controllers.GetUsersWithProducts)
	users.Get("/getUser/:user_id", Controllers.GetUserById)
	users.Get("/getUserProducts/:user_id", Controllers.GetUserProductsById)
	users.Delete("/deleteUser/:user_id", Controllers.DeleteUserById)
	users.Put("/updateUser/:user_id", Controllers.UpdateUserById)

	// Products Routers
	products := app.Group("/products")
	products.Post("/addProduct", Controllers.AddProduct)
	products.Get("/getProducts", Controllers.GetProducts)
	products.Get("/getProduct/:product_id", Controllers.GetProductById)
	products.Put("/updateProduct/:product_id", Controllers.UpdateProductById)
	products.Delete("/deleteProduct/:product_id", Controllers.DeleteProductById)
}
