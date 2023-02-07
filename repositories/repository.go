package repositories

import (
	"test.com/structs"

	"github.com/jinzhu/gorm"
)

// Repository interface
type RepositoryInterface interface {
	CreateUser(user structs.User) error
	GetUsers() ([]structs.User, error)
	CreateProduct(product structs.Product) error
	GetProducts() ([]structs.Product, error)
}

// DbHandler struct
type DbHandler struct {
	Db *gorm.DB
}

// Repository struct
type Repository struct {
	DbHandler *DbHandler
}

// CreateUser method
func (r *Repository) CreateUser(user structs.User) error {
	return r.DbHandler.Db.Create(&user).Error
}

// GetUsers method
func (r *Repository) GetUsers() ([]structs.User, error) {
	var users []structs.User
	if err := r.DbHandler.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// CreateProduct method
func (r *Repository) CreateProduct(product structs.Product) error {
	return r.DbHandler.Db.Create(&product).Error
}

// GetProducts method
func (r *Repository) GetProducts() ([]structs.Product, error) {
	var products []structs.Product
	if err := r.DbHandler.Db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// NewUserRepository returns a new User repository
func NewUserRepository(dbHandler DbHandler) RepositoryInterface {
	return &Repository{
		DbHandler: &dbHandler,
	}
}

// NewProductRepository returns a new Product repository
func NewProductRepository(dbHandler DbHandler) RepositoryInterface {
	return &Repository{
		DbHandler: &dbHandler,
	}
}
