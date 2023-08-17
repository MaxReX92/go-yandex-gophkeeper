package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/db"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// PostgresDBServiceConfig contains required configuration for postgres db service.
type PostgresDBServiceConfig interface {
	ConnectionString() string
}

type service struct {
	conn *sql.DB
}

func NewDBService(ctx context.Context, conf PostgresDBServiceConfig) (*service, error) {
	conn, err := sql.Open("pgx", conf.ConnectionString())
	if err != nil {
		return nil, logger.WrapError("open db connection", err)
	}

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, logger.WrapError("ping db connection", err)
	}

	return &service{
		conn: conn,
	}, err
}

func (d *service) CallInTransaction(ctx context.Context, action func(context.Context, db.Executor) error) error {
	_, err := d.CallInTransactionResult(ctx, func(ctx context.Context, executor db.Executor) ([]any, error) {
		return nil, action(ctx, executor)
	})

	return err
}

func (d *service) CallInTransactionResult(ctx context.Context, action func(context.Context, db.Executor) ([]any, error)) ([]any, error) {
	tx, err := d.conn.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return nil, logger.WrapError("begin transaction in postgresql database", err)
	}

	executor := newExecutor(tx)
	result, err := action(ctx, executor)
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			logger.ErrorFormat("failed to rollback transaction: %v", rollbackError)
		}

		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, logger.WrapError("commit transaction", err)
	}

	return result, nil
}
