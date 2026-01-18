package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	q := "SELECT id, email, password FROM users WHERE email = $1"

	if err := r.db.QueryRow(
		ctx,
		q,
		email,
	).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Create(ctx context.Context, u *User) error {
	q := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, q, u.Name, u.Email, u.Password)

	return err
}

func (r *UserRepo) FindById(ctx context.Context, id int64) (*User, error) {
	var user User

	q := `
		SELECT id, name, email FROM users WHERE id = $1
	`

	err := r.db.QueryRow(
		ctx,
		q,
		id,
	).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
