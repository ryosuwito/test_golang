package handlers

import (
	"log"
	"net/http"

	"test.com/repositories"
	"test.com/structs"

	"github.com/gin-gonic/gin"
)

// UserHandler struct
type UserHandler struct {
	Repository repositories.RepositoryInterface
}

// CreateUser method
func (u *UserHandler) CreateUser(c *gin.Context) {
	var user structs.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.Repository.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// GetUsers method
func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.Repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// ProductHandler struct
type ProductHandler struct {
	Repository repositories.RepositoryInterface
}

// CreateProduct method
func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var product structs.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.Repository.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

// GetProducts method
func (p *ProductHandler) GetProducts(c *gin.Context) {
	products, err := p.Repository.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GoRoutineHandler struct
type GoRoutineHandler struct{}

// StartGoRoutine method
func (g *GoRoutineHandler) StartGoRoutine(c *gin.Context) {
	// Start go routine to fetch data from external API
	go func() {
		data, err := fetchDataFromAPI()
		if err != nil {
			log.Printf("Failed to fetch data from API: %v", err)
			return
		}

		log.Printf("Data fetched successfully: %v", data)
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Go routine started successfully"})
}

// fetchDataFromAPI menunjukkan bagaimana Go routine bisa membantu menangani komputasi yang rumit.
// Dalam hal ini, urutan Fibonacci dikalkulasikan dengan pendekatan rekursif.
func fetchDataFromAPI() (interface{}, error) {
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	result := fibonacci(45)
	return result, nil
}
