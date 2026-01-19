package user

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterUserRequest struct {
	Name     string `validate:"required,min=3"`
	Email    string `validate:"email,required"`
	Password string `validate:"required,min=6"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
