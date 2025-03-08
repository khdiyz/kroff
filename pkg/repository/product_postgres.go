package repository

import (
	"encoding/json"
	"kroff/pkg/models"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(product models.CreateProduct) (int64, error) {
	name, err := json.Marshal(product.Name)
	if err != nil {
		return 0, err
	}

	query := `
	INSERT INTO products (category_id, name, code, price, photo)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	var id int64
	err = r.db.Get(&id, query, product.CategoryId, name, product.Code, product.Price, product.Photo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductPostgres) GetAllProducts() ([]models.Product, error) {
	query := `
	SELECT id, category_id, name, code, price, photo
	FROM products
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var (
			name    []byte
			product models.Product
		)

		err := rows.Scan(&product.ID, &product.CategoryId, &name, &product.Code, &product.Price, &product.Photo)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(name, &product.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductPostgres) GetProductByID(id int64) (models.Product, error) {
	query := `
	SELECT id, category_id, name, code, price, photo
	FROM products
	WHERE id = $1
	`

	var (
		product models.Product
		name    []byte
	)

	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.CategoryId, &name, &product.Code, &product.Price, &product.Photo)
	if err != nil {
		return models.Product{}, err
	}

	err = json.Unmarshal(name, &product.Name)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductPostgres) UpdateProduct(product models.UpdateProduct) error {
	query := `
	UPDATE products
	SET category_id = $1, name = $2, code = $3, price = $4, photo = $5
	WHERE id = $6
	`
	name, err := json.Marshal(product.Name)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, product.CategoryId, name, product.Code, product.Price, product.Photo, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductPostgres) DeleteProduct(id int64) error {
	query := `
	DELETE FROM products
	WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductPostgres) GetAllProductsCount() (int64, error) {
	query := `
	SELECT COUNT(*) FROM products
	`

	var count int64
	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
