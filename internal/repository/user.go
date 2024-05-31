package repository

import (
	"context"

	"bit-project/gateway/ent"
)

type UserRepository interface {
	Create(ctx context.Context, user ent.User) error
	Exist(ctx context.Context, aud string) (bool, error)
}
