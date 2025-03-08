package models

type NameTranslation struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
}

func (n *NameTranslation) Get(lang string) string {
	if lang == "ru" {
		return n.Ru
	}
	return n.Uz
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" default:"admin"`
	Password string `json:"password" binding:"required" default:"admin"`
}

type Pagination struct {
	Page       int `json:"page"  default:"1"`
	Limit      int `json:"limit" default:"10"`
	Offset     int `json:"-" default:"0"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
}

type FilterOptions struct {
	Filters map[string]any
	SortBy  string
	Order   string
	Page    int
	Limit   int
}