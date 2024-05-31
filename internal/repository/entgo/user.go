package entgo

import (
	"context"
	"errors"

	"bit-project/gateway/ent"
	"bit-project/gateway/ent/user"
	"bit-project/gateway/internal/repository"
)

type userRepository struct {
	*TxHandler
}

func NewUserRepository(handler *TxHandler) repository.UserRepository {
	return &userRepository{handler}
}

// Create creates a new user only when the server is initialized.
func (r *userRepository) Create(ctx context.Context, user ent.User) error {
	Mutex.Lock()
	defer Mutex.Unlock()

	tx, err := r.GetExecutor(ctx)
	if err != nil {
		return err
	}

	isExist, err := tx.User.Query().Exist(ctx)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("user already exists")
	}

	_, err = tx.User.
		Create().
		SetEmail(user.Email).
		SetKakaoSub(user.KakaoSub).
		SetName(user.Name).
		SetLastLoginDate(user.LastLoginDate).
		SetRefreshToken(user.RefreshToken).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Exist(ctx context.Context, aud string) (bool, error) {
	Mutex.Lock()
	defer Mutex.Unlock()

	tx, err := r.GetExecutor(ctx)
	if err != nil {
		return false, err
	}

	exists, err := tx.User.Query().
		Where(user.KakaoSubEQ(aud)).
		Exist(ctx)
	if err != nil {
		return false, err
	}
	return exists, nil
}
