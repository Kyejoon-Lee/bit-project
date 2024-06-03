package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	Login(ctx context.Context, claims jwt.Claims) (string, error)
}
