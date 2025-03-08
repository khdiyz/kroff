package models

type Product struct {
	ID         int64           `json:"id"`
	CategoryId int64           `json:"category_id"`
	Name       NameTranslation `json:"name"`
	Code       string          `json:"code"`
	Price      int64           `json:"price"`
	Photo      string          `json:"photo"`
}

type ProductPublic struct {
	ID         int64  `json:"id"`
	CategoryId int64  `json:"category_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Price      int64  `json:"price"`
	Photo      string `json:"photo"`
}

type CreateProduct struct {
	CategoryId int64           `json:"category_id"`
	Name       NameTranslation `json:"name"`
	Code       string          `json:"code"`
	Price      int64           `json:"price"`
	Photo      string          `json:"photo"`
}

type UpdateProduct struct {
	ID         int64           `json:"-"`
	CategoryId int64           `json:"category_id"`
	Name       NameTranslation `json:"name"`
	Code       string          `json:"code"`
	Price      int64           `json:"price"`
	Photo      string          `json:"photo"`
}
