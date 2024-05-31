package entgo

import (
	"context"
	"errors"
	"fmt"

	"bit-project/gateway/ent"

	log "github.com/sirupsen/logrus"
)

const (
	useTransactionKey = "useTx"
	transactionKey    = "tx"
)

type TxHandler struct {
	client *ent.Client
}

func NewTxHandler(client *ent.Client) *TxHandler {
	return &TxHandler{client}
}

func (r *TxHandler) GetExecutor(ctx context.Context) (*ent.Client, error) {
	return r.client, nil
}

func (r *TxHandler) BeginTx(ctx context.Context) (context.Context, context.CancelFunc, error) {
	ctxWithUseTx := context.WithValue(ctx, useTransactionKey, true)
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return ctxWithUseTx, nil, fmt.Errorf("starting a transaction: %w", err)
	}
	ctxWithTx := context.WithValue(ctxWithUseTx, transactionKey, tx)
	ctxWithCancel, cancelFunc := context.WithCancel(ctxWithTx)
	return ctxWithCancel, func() {
		err = tx.Rollback()
		if err != nil {
			log.Warnf("rollback is failed - %v", err)
		}
		cancelFunc()
	}, nil
}

func (r *TxHandler) Commit(ctx context.Context) error {
	useTx := ctx.Value(useTransactionKey).(bool)
	if !useTx {
		return nil
	} else {
		tx, err := r.client.Tx(ctx)
		if err != nil {
			return errors.New("get transaction from context fail")
		}
		// Commit the transaction.
		return tx.Commit()
	}
}
