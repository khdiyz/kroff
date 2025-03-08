package repository

import (
	"encoding/json"
	"fmt"
	"kroff/pkg/models"

	sq "github.com/Masterminds/squirrel"
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

func (r *ProductPostgres) GetAllProductsPublic(lang string, options models.FilterOptions) ([]models.ProductPublic, int, error) {
	products := []models.ProductPublic{}

	query := sq.Select("id", "category_id", "name->>"+fmt.Sprintf("'%s'", lang)+" as name", "code", "price", "photo").
		From("products").
		Where(sq.Expr("true"))

	countQuery := sq.Select("count(id)").From("products").Where(sq.Expr("true"))

	filters := options.Filters

	if categoryId, ok := filters["category_id"]; ok {
		if categoryId.(int64) != 0 {
			query = query.Where(sq.Eq{"category_id": categoryId})
			countQuery = countQuery.Where(sq.Eq{"category_id": categoryId})
		}
	}

	if options.SortBy != "" {
		order := "ASC"
		if options.Order == "desc" {
			order = "DESC"
		}
		query = query.OrderBy(fmt.Sprintf("%s %s", options.SortBy, order))
	} else {
		query = query.OrderBy("created_at DESC") // Default sorting
	}

	if options.Limit > 0 {
		offset := (options.Page - 1) * options.Limit
		query = query.Limit(uint64(options.Limit)).Offset(uint64(offset))
	}

	sqlQuery, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, 0, err
	}

	fmt.Println(sqlQuery, args)

	rows, err := r.db.Queryx(sqlQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductPublic
		err := rows.Scan(&product.ID, &product.CategoryId, &product.Name, &product.Code, &product.Price, &product.Photo)
		if err != nil {
			return nil, 0, err
		}

		products = append(products, product)
	}

	countSql, countArgs, err := countQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, 0, err
	}

	var totalCount int
	err = r.db.Get(&totalCount, countSql, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}
