package product

import "github.com/jackc/pgx/v5/pgxpool"

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{db}
}
