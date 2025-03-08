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
