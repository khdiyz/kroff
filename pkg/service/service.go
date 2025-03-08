package service

import (
	"context"
	"io"
	"kroff/config"
	"kroff/pkg/models"
	"kroff/pkg/repository"
	"kroff/pkg/storage"
)

type Services struct {
	Category
	Authorization
	File
	Product
}

func NewServices(repo *repository.Repository, storage *storage.Storage, cfg *config.Config) *Services {
	return &Services{
		Category:      NewCategoryService(repo),
		File:          NewFileService(storage, cfg),
		Authorization: NewAuthService(repo, cfg),
		Product:       NewProductService(repo),
	}
}

type Authorization interface {
	Login(request models.LoginRequest) (string, error)
	ParseToken(token string) (*jwtCustomClaim, error)
}

type Category interface {
	CreateCategory(category models.CreateCategory) (int64, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int64) (models.Category, error)
	UpdateCategory(category models.UpdateCategory) error
	DeleteCategory(id int64) error
	GetAllCategoriesPublic(lang string) ([]models.CategoryPublic, error)
}

type File interface {
	UploadFile(ctx context.Context, file io.Reader, fileSize int64, contentType string) (string, error)
	UploadWithName(ctx context.Context, file io.Reader, fileSize int64, contentType string, fileName string) error
}

type Product interface {
	CreateProduct(product models.CreateProduct) (int64, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int64) (models.Product, error)
	UpdateProduct(product models.UpdateProduct) error
	DeleteProduct(id int64) error
	GetAllProductsPublic(lang string, options models.FilterOptions) ([]models.ProductPublic, int, error)
}
