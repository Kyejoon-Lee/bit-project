package repository

import "context"

type TxHandler interface {
	BeginTx(ctx context.Context) (context.Context, context.CancelFunc, error)
	Commit(ctx context.Context) error
}
