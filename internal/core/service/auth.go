package service

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-learn-news-portal/internal/core/entity"
	"go-learn-news-portal/internal/framework/primary/rest/request"
	"go-learn-news-portal/internal/framework/primary/rest/response"
	"go-learn-news-portal/internal/framework/secondary/repository"
	"go-learn-news-portal/library/v1/convert"
	"go-learn-news-portal/library/v1/handling"
	"go-learn-news-portal/library/v1/middleware"
)

type Auth interface {
	Login(ctx context.Context, req *request.Login) (*response.Login, error)
}

func NewAuth(userRepo repository.User) Auth {
	return &authCore{userRepo: userRepo}
}

type authCore struct {
	userRepo repository.User
}

func (c *authCore) Login(ctx context.Context, req *request.Login) (*response.Login, error) {
	e := &entity.User{
		Email: req.Email,
	}

	err := c.userRepo.FindUserByEmail(ctx, e)
	if err != nil {
		return nil, err
	}

	err = convert.CheckPasswordHash(req.Password, e.Password)
	if err != nil {
		return nil, handling.NewHttpError400(errors.New("password is invalid"))
	}

	claims := &middleware.Payload{
		ID: e.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer: "learn-news-portal",
		},
	}

	token, err := middleware.GenerateAccessToken(claims)
	if err != nil {
		return nil, err
	}

	refreshToken, err := middleware.GenerateRefreshToken(&middleware.RefreshPayload{
		Payload: *claims,
	})
	if err != nil {
		return nil, err
	}

	return &response.Login{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}
