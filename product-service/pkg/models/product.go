package models

type Product struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Stock     int    `json:"stock"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
