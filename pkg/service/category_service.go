package service

import (
	"kroff/pkg/models"
	"kroff/pkg/repository"
	"kroff/utils/response"

	"google.golang.org/grpc/codes"
)

type CategoryService struct {
	repo *repository.Repository
}

func NewCategoryService(repo *repository.Repository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category models.CreateCategory) (int64, error) {
	return s.repo.Category.CreateCategory(category)
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.Category.GetAllCategories()
}

func (s *CategoryService) GetAllCategoriesPublic(lang string) ([]models.CategoryPublic, error) {
	lang = getLang(lang)

	categories, err := s.repo.Category.GetAllCategoriesPublic(lang)
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	allProductsCategoryName := models.NameTranslation{
		Uz: "Barcha tovarlar",
		Ru: "Все товары",
	}

	allProductsCount, err := s.repo.Product.GetAllProductsCount()
	if err != nil {
		return nil, response.ServiceError(err, codes.Internal)
	}

	allProductsCategory := models.CategoryPublic{
		ID:           0,
		Name:         allProductsCategoryName.Get(lang),
		Photo:        "",
		ProductCount: allProductsCount,
	}

	categories = append(categories, allProductsCategory)

	return categories, nil
}

func getLang(lang string) string {
	if lang == "ru" {
		return "ru"
	}

	return "uz"
}

func (s *CategoryService) GetCategoryByID(id int64) (models.Category, error) {
	return s.repo.Category.GetCategoryByID(id)
}

func (s *CategoryService) UpdateCategory(category models.UpdateCategory) error {
	return s.repo.Category.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id int64) error {
	return s.repo.Category.DeleteCategory(id)
}
