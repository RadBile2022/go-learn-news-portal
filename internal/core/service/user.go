package service

import (
	"context"
	"errors"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/internal/framework/primary/rests/request"
	"go-learn-news-portal/internal/framework/primary/rests/response"
	"go-learn-news-portal/internal/framework/secondary/repository"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/handling"
	"go-learn-news-portal/library/v1/middleware"
)

type User interface {
	FindUserByID(ctx context.Context) (*response.User, error)
	FindUserByEmail(ctx context.Context, req *request.FindUserByEmail) (*response.User, error)
	UpdatePassword(ctx context.Context, req *request.UpdatePassword) error
}

type userCore struct {
	userRepo repository.User
}

// UpdatePassword implements UserService.
func (c *userCore) UpdatePassword(ctx context.Context, req *request.UpdatePassword) error {

	if req.NewPassword != req.ConfirmPassword {
		return handling.NewHttpError400(errors.New("password and confirm password not match"))
	}

	// find and check password
	e := &entity.User{ID: middleware.GetUserIDFromContext(ctx)}

	err := c.userRepo.FindUserByID(ctx, e)
	if err != nil {
		return err
	}

	err = convert.CheckPasswordHash(req.CurrentPassword, e.Password)
	if err != nil {
		return handling.NewHttpError401(errors.New("invalid password"))
	}

	// if same password, replace old password
	password, err := convert.HashPassword(req.NewPassword)
	e.Password = password
	if err = c.userRepo.UpdatePassword(ctx, e); err != nil {
		return err
	}

	return nil
}

// FindUserByID implements UserService.
func (c *userCore) FindUserByID(ctx context.Context) (*response.User, error) {
	e := &entity.User{ID: middleware.GetUserIDFromContext(ctx)}

	if err := c.userRepo.FindUserByID(ctx, e); err != nil {
		return nil, err
	}

	return response.FromEntityUser(e), nil
}

func (c *userCore) FindUserByEmail(ctx context.Context, req *request.FindUserByEmail) (*response.User, error) {
	e := &entity.User{Email: req.Email}

	if err := c.userRepo.FindUserByEmail(ctx, e); err != nil {
		return nil, err
	}

	return response.FromEntityUser(e), nil
}

func NewUser(userRepo repository.User) User {
	return &userCore{userRepo: userRepo}
}
