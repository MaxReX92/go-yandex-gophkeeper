package postgres

import (
	"context"
	"database/sql"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/server/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
)

type PostgresDBStorageConfig interface {
	ConnectionString() string
}

type dbSecret struct {
	Identity     sql.NullString
	UserIdentity sql.NullString
	Content      sql.RawBytes
}

type dbStorage struct {
	conn *sql.DB
}

func NewDBStorage(ctx context.Context, conf PostgresDBStorageConfig) (*dbStorage, error) {
	conn, err := sql.Open("pgx", conf.ConnectionString())
	if err != nil {
		return nil, logger.WrapError("open db connection", err)
	}

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, logger.WrapError("ping db connection", err)
	}

	return &dbStorage{
		conn: conn,
	}, nil
}

func (d *dbStorage) AddSecret(ctx context.Context, userId string, secret *generated.Secret) error {
	return d.callInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		command := "INSERT INTO secret VALUES ($1, $2, $3)"
		_, err := tx.ExecContext(ctx, command, secret.Identity, userId, secret.Content)
		if err != nil {
			return logger.WrapError("call add user query", err)
		}

		return nil
	})
}

func (d *dbStorage) ChangeSecret(ctx context.Context, userId string, secret *generated.Secret) error {
	return d.callInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		command := "UPDATE secret " +
			"SET content = $1 " +
			"WHERE id = $2 and userId = $3"
		_, err := tx.ExecContext(ctx, command, secret.Content, secret.Identity, userId)
		if err != nil {
			return logger.WrapError("call update user query", err)
		}

		return nil
	})
}

func (d *dbStorage) GetAllSecrets(ctx context.Context, userId string) ([]*generated.Secret, error) {
	result, err := d.callInTransactionResult(ctx, func(ctx context.Context, tx *sql.Tx) ([]any, error) {
		command := "SELECT s.id, s.userId, s.content " +
			"FROM secret s " +
			"WHERE s.userId = $1"

		rows, err := tx.QueryContext(ctx, command, userId)
		if err != nil {
			return nil, logger.WrapError("call get car query", err)
		}
		defer rows.Close()

		if !rows.Next() {
			return nil, rows.Err()
		}

		secrets := []any{}
		for rows.Next() {
			var secret dbSecret
			err = rows.Scan(&secret.Identity, &secret.UserIdentity, &secret.Content)
			if err != nil {
				return nil, logger.WrapError("scan rows", err)
			}

			secrets = append(secrets, secret)
		}

		err = rows.Err()
		if err != nil {
			return nil, logger.WrapError("get rows", err)
		}

		return secrets, nil
	})

	if err != nil {
		return nil, logger.WrapError("get all user secrets", err)
	}

	secretsLen := len(result)
	secrets := make([]*generated.Secret, secretsLen)
	for i := 0; i < secretsLen; i++ {
		dSecret := result[i].(dbSecret)

		if !dSecret.Identity.Valid {
			return nil, logger.WrapError("read secret id", storage.ErrInvalidDBValue)
		}
		if !dSecret.UserIdentity.Valid {
			return nil, logger.WrapError("read secret user id", storage.ErrInvalidDBValue)
		}

		secrets[i] = &generated.Secret{
			Identity: dSecret.Identity.String,
			Content:  dSecret.Content,
		}
	}

	return secrets, nil
}

func (d *dbStorage) RemoveSecret(ctx context.Context, userId string, secret *generated.Secret) error {
	return d.callInTransaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
		command := "DELETE secret " +
			"WHERE id = $1 and userId = $2"
		_, err := tx.ExecContext(ctx, command, secret.Identity, userId)
		if err != nil {
			return logger.WrapError("call delete user query", err)
		}

		return nil
	})
}

func (d *dbStorage) callInTransaction(ctx context.Context, action func(context.Context, *sql.Tx) error) error {
	_, err := d.callInTransactionResult(ctx, func(ctx context.Context, tx *sql.Tx) ([]any, error) {
		return nil, action(ctx, tx)
	})

	return err
}

func (d *dbStorage) callInTransactionResult(ctx context.Context, action func(context.Context, *sql.Tx) ([]any, error)) ([]any, error) {
	tx, err := d.conn.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return nil, logger.WrapError("begin transaction in postgresql database", err)
	}

	result, err := action(ctx, tx)
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
