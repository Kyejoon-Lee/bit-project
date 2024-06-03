package module

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"bit-project/gateway/config"
	"bit-project/gateway/db/postgresql"
	"bit-project/gateway/internal/domain"
	"bit-project/gateway/internal/repository"
	"bit-project/gateway/internal/repository/entgo"
	"bit-project/gateway/internal/service"
)

var (
	serverContextOnce      sync.Once
	serverContext          context.Context
	serverCancelFunc       context.CancelFunc
	txHandlerOnce          sync.Once
	txHandlerInstance      *entgo.TxHandler
	userRepositoryOnce     sync.Once
	userRepositoryInstance repository.UserRepository
	userServiceOnce        sync.Once
	userServiceInstance    domain.UserService
)

func ServerContext() (context.Context, func()) {
	serverContextOnce.Do(func() {
		serverContext, serverCancelFunc = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	})
	return serverContext, serverCancelFunc
}
func TxHandler() *entgo.TxHandler {
	txHandlerOnce.Do(func() {
		txHandlerInstance = entgo.NewTxHandler(postgresql.Open(Config()))
	})
	return txHandlerInstance
}

func UserService() domain.UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = service.NewUserService(
			TxHandler(),
			UserRepository(),
		)
	})
	return userServiceInstance
}
func UserRepository() repository.UserRepository {
	userRepositoryOnce.Do(func() {
		userRepositoryInstance = entgo.NewUserRepository(TxHandler())
	})
	return userRepositoryInstance
}

func Config() *config.Config {
	return config.GetConfig()
}
