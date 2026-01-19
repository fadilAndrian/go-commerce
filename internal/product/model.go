package product

import "time"

type Product struct {
	ID        int64
	Name      string
	Price     int
	Stock     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductRequest struct {
	Name  string `validate:"required,min=3"`
	Price int    `validate:"required,gt=0"`
	Stock int    `validate:"required,gte=0"`
}
