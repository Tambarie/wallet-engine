package dto

type User struct {
	ID           string `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty" binding:"required"`
	LastName     string `json:"last_name,omitempty" binding:"required"`
	Email        string `json:"email,omitempty" binding:"required"`
	Password     string `json:"password,omitempty" binding:"required" validate:"min=4"`
	HashPassword string `json:"-"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}
