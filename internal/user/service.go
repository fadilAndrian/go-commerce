package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fadilAndrian/go-commerce/internal/helper"
)

type UserService struct {
	ur *UserRepo
}

func NewUserService(ur *UserRepo) *UserService {
	return &UserService{ur}
}

func (s *UserService) Register(ctx context.Context, request *RegisterUserRequest) error {
	existedUser, err := s.ur.FindByEmail(ctx, request.Email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if existedUser != nil {
		return errors.New("Email has been used")
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := &User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	return s.ur.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context, request *LoginUserRequest) (string, error) {
	user, err := s.ur.FindByEmail(ctx, request.Email)

	if errors.Is(err, sql.ErrNoRows) {
		return "", sql.ErrNoRows
	}

	if err != nil {
		return "", err
	}

	if err := helper.ValidatePassword(request.Password, user.Password); err != nil {
		return "", sql.ErrNoRows
	}

	token, err := helper.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) AuthProfile(ctx context.Context, id int64) (*User, error) {
	user, err := s.ur.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
