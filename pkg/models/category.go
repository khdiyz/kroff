package models

type Category struct {
	ID    int64           `json:"id"`
	Name  NameTranslation `json:"name"`
	Photo string          `json:"photo"`
}

type CategoryPublic struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Photo        string `json:"photo"`
	ProductCount int64  `json:"product_count"`
}

type CreateCategory struct {
	Name  NameTranslation `json:"name"`
	Photo string          `json:"photo"`
}

type UpdateCategory struct {
	ID    int64           `json:"id"`
	Name  NameTranslation `json:"name"`
	Photo string          `json:"photo"`
}
