package routes

import (
	"fahmi-wallet/controllers"

	"fahmi-wallet/middleware"

	"github.com/gin-gonic/gin"
)

var userController = controllers.UsersController{}
var authController = controllers.AuthController{}
var walletController = controllers.WalletController{}
var productController = controllers.ProductController{}

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// auths apis
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", userController.CreateUser)
		}

		// apis for user
		users := api.Group("/users")
		{
			users.GET("/", middleware.RoleMiddleware("admin"), userController.GetUsers)
			users.GET("/:id", middleware.RoleMiddleware("user", "admin"), userController.GetUser)
			users.PUT("/:id", middleware.RoleMiddleware("admin"), userController.UpdateUser)
			users.DELETE("/:id", middleware.RoleMiddleware("user", "admin"), userController.DeleteUser)
		}

		// apis for wallet
		wallet := api.Group("/wallet")
		{
			// apis for manage wallet
			wallet.POST("/:user_id/create", middleware.RoleMiddleware("user", "admin"), walletController.CreateWallet)
			wallet.GET("/:user_id", middleware.RoleMiddleware("user", "admin"), walletController.GetWallet)
			wallet.PUT("/:user_id/update", middleware.RoleMiddleware("user", "admin"), walletController.UpdateWallet)

			// apis for deposit and withdraw balance
			wallet.POST("/:user_id/deposit", middleware.RoleMiddleware("user", "admin"), walletController.Deposit)
			wallet.POST("/:user_id/withdraw", middleware.RoleMiddleware("user", "admin"), walletController.Withdraw)
		}

		// apis for product
		product := api.Group("/products")
		{
			product.POST("/", middleware.RoleMiddleware("admin"), productController.CreateProduct)
			product.GET("/", middleware.RoleMiddleware("admin", "user"), productController.GetAllProducts)
			product.GET("/:id", middleware.RoleMiddleware("admin", "user"), productController.GetProductByID)
			product.PUT("/:id", middleware.RoleMiddleware("admin"), productController.UpdateProduct)
			product.DELETE("/:id", middleware.RoleMiddleware("admin"), productController.DeleteProduct)
		}
	}
}
