package product

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) List(c context.Context) ([]Product, error) {
	q := `
		SELECT id, name, price, created_at, updated_at, stock FROM products ORDER BY created_at DESC
	`

	rows, err := r.db.Query(c, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepo) Create(c context.Context, p *Product) error {
	q := `
		INSERT INTO products (name, price, stock)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(c, q, p.Name, p.Price, p.Stock)

	return err
}

func (r *ProductRepo) FindById(c context.Context, id int64) (*Product, error) {
	var product Product

	q := `
		SELECT id, name, price, stock FROM products WHERE id = $1
	`

	row := r.db.QueryRow(c, q, id)
	if err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) Update(c context.Context, p *Product) error {
	q := `
		UPDATE products
		SET name = $1,  price = $2, stock = $3, updated_at = NOW()
		WHERE id = $4
	`

	_, err := r.db.Exec(c, q, p.Name, p.Price, p.Stock, p.ID)

	return err
}

func (r *ProductRepo) Delete(c context.Context, p *Product) error {
	q := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(c, q, p.ID)

	return err
}
