package service

import (
	"kroff/pkg/models"
	"kroff/pkg/repository"
	"kroff/utils/response"

	"google.golang.org/grpc/codes"
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

func (s *ProductService) GetAllProductsPublic(lang string, options models.FilterOptions) ([]models.ProductPublic, int, error) {
	lang = getLang(lang)

	products, totalCount, err := s.repo.Product.GetAllProductsPublic(lang, options)
	if err != nil {
		return nil, 0, response.ServiceError(err, codes.Internal)
	}

	return products, totalCount, nil
}
