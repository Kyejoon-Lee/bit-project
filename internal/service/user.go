package service

import (
	"context"

	"bit-project/gateway/internal/domain"
	"bit-project/gateway/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type userService struct {
	txHandler      repository.TxHandler
	userRepository repository.UserRepository
}

func NewUserService(
	txHandler repository.TxHandler,
	userRepository repository.UserRepository,
) domain.UserService {
	return &userService{
		txHandler:      txHandler,
		userRepository: userRepository,
	}
}

func (s *userService) Login(ctx context.Context, claims jwt.Claims) (string, error) {
	return "", nil

}
