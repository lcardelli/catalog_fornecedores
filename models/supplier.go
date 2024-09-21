package models

type Supplier struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Category string `json:"category"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	GoogleID string `json:"google_id"` // ID do Google
}


