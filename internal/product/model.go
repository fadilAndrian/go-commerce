package product

import "time"

type Product struct {
	ID        int64
	Name      string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
