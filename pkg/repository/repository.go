package repository

import (
	"kroff/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Category
	Product
}

type Category interface {
	CreateCategory(category models.CreateCategory) (int64, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int64) (models.Category, error)
	UpdateCategory(category models.UpdateCategory) error
	DeleteCategory(id int64) error
	GetAllCategoriesPublic(lang string) ([]models.CategoryPublic, error)
}

type Product interface {
	CreateProduct(product models.CreateProduct) (int64, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int64) (models.Product, error)
	UpdateProduct(product models.UpdateProduct) error
	DeleteProduct(id int64) error
	GetAllProductsCount() (int64, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category: NewCategoryPostgres(db),
		Product:  NewProductPostgres(db),
	}
}
