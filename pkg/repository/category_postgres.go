package repository

import (
	"encoding/json"
	"kroff/pkg/models"

	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategory(category models.CreateCategory) (int64, error) {
	name, err := json.Marshal(category.Name)
	if err != nil {
		return 0, err
	}

	query := `
	INSERT INTO categories (name, photo)
	VALUES ($1, $2)
	RETURNING id
	`
	var id int64
	err = r.db.QueryRow(query, name, category.Photo).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CategoryPostgres) GetAllCategories() ([]models.Category, error) {
	query := `
	SELECT id, name, photo FROM categories
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var (
			category models.Category
			name     []byte
		)
		err := rows.Scan(&category.ID, &name, &category.Photo)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(name, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryPostgres) GetCategoryByID(id int64) (models.Category, error) {
	query := `
	SELECT id, name, photo FROM categories WHERE id = $1
	`
	var (
		category models.Category
		name     []byte
	)

	err := r.db.QueryRow(query, id).Scan(&category.ID, &name, &category.Photo)
	if err != nil {
		return models.Category{}, err
	}

	err = json.Unmarshal(name, &category.Name)
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (r *CategoryPostgres) UpdateCategory(category models.UpdateCategory) error {
	query := `
	UPDATE categories SET name = $1, photo = $2 WHERE id = $3
	`
	name, err := json.Marshal(category.Name)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, name, category.Photo, category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryPostgres) DeleteCategory(id int64) error {
	query := `
	DELETE FROM categories WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryPostgres) GetAllCategoriesPublic(lang string) ([]models.CategoryPublic, error) {
	query := `
	SELECT 
		c.id, 
		c.name->>$1 as name, 
		c.photo,
		COUNT(p.id) as product_count
	FROM categories c
	LEFT JOIN products p ON c.id = p.category_id
	GROUP BY c.id
	ORDER BY c.id;`

	rows, err := r.db.Query(query, lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.CategoryPublic, 0)
	for rows.Next() {
		var category models.CategoryPublic
		err := rows.Scan(&category.ID, &category.Name, &category.Photo, &category.ProductCount)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
