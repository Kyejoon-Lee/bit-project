package service

import (
	"bit-project/gateway/internal/domain"
	"bit-project/gateway/internal/repository"
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

func (s *userService) Login() {

}
