package service

import (
	"kroff/pkg/models"
	"kroff/pkg/repository"
)

type ProductService struct {
	repo *repository.Repository
}

func NewProductService(repo *repository.Repository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product models.CreateProduct) (int64, error) {
	return s.repo.Product.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.Product.GetAllProducts()
}

func (s *ProductService) GetProductByID(id int64) (models.Product, error) {
	return s.repo.Product.GetProductByID(id)
}

func (s *ProductService) UpdateProduct(product models.UpdateProduct) error {
	return s.repo.Product.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int64) error {
	return s.repo.Product.DeleteProduct(id)
}