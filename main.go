package main

import (
	"fmt"
	"tokobelanja-golang/conf"
	"tokobelanja-golang/controller"
	"tokobelanja-golang/middleware"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"tokobelanja-golang/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := conf.InitDB()

	// repository
	userRepository := repository.NewUserRepository(db)
	transactionRepository := repository.NewTransactionHistoryRepository(db)
	productRepository := repository.NewProductRepository(db)
	categoriesRepository := repository.NewCategoryRepository(db)

	// service
	userService := service.NewUserService(userRepository)
	transactionService := service.NewTransactionHistoryService(transactionRepository, productRepository, userRepository)
	categoriesService := service.NewCategoryService(categoriesRepository)
	productService := service.NewProductService(productRepository)

	// controller
	userController := controller.NewUserController(userService)
	transactionController := controller.NewTransactionHistoryController(transactionService, userService, productService)
	categoriesController := controller.NewCategoryController(categoriesService, userService)
	productController := controller.NewProductController(productService, userService)

	router := gin.Default()

	// routing
	// user
	router.POST("users/register", userController.RegisterUser)
	router.POST("users/login", userController.Login)
	router.POST("users/topup", middleware.AuthMiddleware(), userController.TopUpSaldo)

	// transaction
	router.POST("transactions", middleware.AuthMiddleware(), transactionController.NewTransaction)
	router.POST("transactions/my-transactions", middleware.AuthMiddleware(), transactionController.GetMyTransaction)
	router.POST("transactions/user-transactions", middleware.AuthMiddleware(), transactionController.GetUserTransaction)

	// categories
	router.POST("categories", middleware.AuthMiddleware(), categoriesController.CreateCategory)
	router.GET("categories", middleware.AuthMiddleware(), categoriesController.GetAllCategory)
	router.PATCH("categories/:id", middleware.AuthMiddleware(), categoriesController.UpdateCategory)
	router.DELETE("categories/:id", middleware.AuthMiddleware(), categoriesController.DeleteCategory)

	// product
	router.POST("products", middleware.AuthMiddleware(), productController.CreateProduct)
	router.PUT("products/:id", middleware.AuthMiddleware(), productController.UpdateProduct)
	router.GET("products", middleware.AuthMiddleware(), productController.GetAllProduct)
	router.DELETE("products/:id", middleware.AuthMiddleware(), productController.DeleteProduct)

	router.Run()

}
