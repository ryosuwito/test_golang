package main

import (
	"github.com/gin-gonic/gin"

	"test.com/handlers"
	"test.com/repositories"
)

func setupRoutes(r *gin.Engine, dbHandler repositories.DbHandler) {
	userRepository := repositories.NewUserRepository(dbHandler)
	userHandler := handlers.UserHandler{Repository: userRepository}

	productRepository := repositories.NewProductRepository(dbHandler)
	productHandler := handlers.ProductHandler{Repository: productRepository}

	goRoutineHandler := handlers.GoRoutineHandler{}

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.GetProducts)
	r.GET("/start-go-routine", goRoutineHandler.StartGoRoutine)
}
