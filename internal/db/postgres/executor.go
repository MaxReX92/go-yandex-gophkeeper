package postgres

import (
	"context"
	"database/sql"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/db"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type dbUser struct {
	id            sql.NullString
	password      sql.NullString
	name          sql.NullString
	personalToken sql.NullString
}

type dbSecret struct {
	id         sql.NullString
	userID     sql.NullString
	secretType sql.NullString
	content    sql.RawBytes
}

type dbExecutor struct {
	tx *sql.Tx
}

func newExecutor(tx *sql.Tx) *dbExecutor {
	return &dbExecutor{tx: tx}
}

func (d *dbExecutor) AddUser(ctx context.Context, id string, username string, password string, personalToken string) error {
	command := "INSERT INTO users VALUES ($1, $2, $3)"
	_, err := d.tx.ExecContext(ctx, command, username, password, personalToken)
	if err != nil {
		return logger.WrapError("call add user query", err)
	}

	return nil
}

func (d *dbExecutor) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	command := "SELECT u.id, u.username, u.password, u.personalToken" +
		"FROM users u " +
		"WHERE u.username = $1"

	rows, err := d.tx.QueryContext(ctx, command, username)
	if err != nil {
		return nil, logger.WrapError("call get user query", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, rows.Err()
	}

	var user dbUser
	err = rows.Scan(&user.id, &user.name, &user.password, &user.personalToken)
	if err != nil {
		return nil, logger.WrapError("scan rows", err)
	}

	err = rows.Err()
	if err != nil {
		return nil, logger.WrapError("get rows", err)
	}

	if !user.id.Valid {
		return nil, logger.WrapError("read user id", db.ErrInvalidDBValue)
	}
	if !user.name.Valid {
		return nil, logger.WrapError("read user name", db.ErrInvalidDBValue)
	}
	if !user.password.Valid {
		return nil, logger.WrapError("read user password", db.ErrInvalidDBValue)
	}
	if !user.personalToken.Valid {
		return nil, logger.WrapError("read user personalToken", db.ErrInvalidDBValue)
	}

	return &model.User{
		Identity:      user.id.String,
		Name:          user.name.String,
		Password:      user.password.String,
		PersonalToken: user.personalToken.String,
	}, nil
}

func (e *dbExecutor) AddSecret(ctx context.Context, userID string, secret *generated.Secret) error {
	typeName, err := e.toTypeName(secret.Type)
	if err != nil {
		return logger.WrapError("convert secret type", err)
	}

	command := "INSERT INTO secret VALUES ($1, $2, $3, $4)"
	_, err = e.tx.ExecContext(ctx, command, secret.Identity, userID, typeName, secret.Content)
	if err != nil {
		return logger.WrapError("call add user query", err)
	}

	return nil
}

func (e *dbExecutor) ChangeSecret(ctx context.Context, userID string, secret *generated.Secret) error {
	command := "UPDATE secret " +
		"SET content = $1 " +
		"WHERE id = $2 and userID = $3"
	_, err := e.tx.ExecContext(ctx, command, secret.Content, secret.Identity, userID)
	if err != nil {
		return logger.WrapError("call update user query", err)
	}

	return nil
}

func (e *dbExecutor) GetAllSecrets(ctx context.Context, userID string) ([]*generated.Secret, error) {
	command := "SELECT s.id, s.userID, s.type, s.content " +
		"FROM secret s " +
		"WHERE s.userID = $1"

	rows, err := e.tx.QueryContext(ctx, command, userID)
	if err != nil {
		return nil, logger.WrapError("call get all secrets query", err)
	}
	defer rows.Close()

	dbSecrets := []dbSecret{}
	for rows.Next() {
		var secret dbSecret
		err = rows.Scan(&secret.id, &secret.userID, &secret.secretType, &secret.content)
		if err != nil {
			return nil, logger.WrapError("scan rows", err)
		}

		dbSecrets = append(dbSecrets, secret)
	}

	err = rows.Err()
	if err != nil {
		return nil, logger.WrapError("get rows", err)
	}

	secretsLen := len(dbSecrets)
	secrets := make([]*generated.Secret, secretsLen)
	for i := 0; i < secretsLen; i++ {
		dSecret := dbSecrets[i]

		if !dSecret.id.Valid {
			return nil, logger.WrapError("read secret id", db.ErrInvalidDBValue)
		}
		if !dSecret.userID.Valid {
			return nil, logger.WrapError("read secret user id", db.ErrInvalidDBValue)
		}
		if !dSecret.secretType.Valid {
			return nil, logger.WrapError("read secret user id", db.ErrInvalidDBValue)
		}

		secretType, err := e.toType(dSecret.secretType.String)
		if err != nil {
			return nil, logger.WrapError("convert type", err)
		}

		secrets[i] = &generated.Secret{
			Identity: dSecret.id.String,
			Type:     secretType,
			Content:  dSecret.content,
		}
	}

	return secrets, nil
}

func (e *dbExecutor) RemoveSecret(ctx context.Context, userID string, secret *generated.Secret) error {
	command := "DELETE from secret " +
		"WHERE id = $1 and userID = $2"
	_, err := e.tx.ExecContext(ctx, command, secret.Identity, userID)
	if err != nil {
		return logger.WrapError("call delete user query", err)
	}

	return nil
}

func (e *dbExecutor) toTypeName(secretType generated.SecretType) (string, error) {
	switch secretType {
	case generated.SecretType_BINARY:
		return "binary", nil
	case generated.SecretType_CARD:
		return "card", nil
	case generated.SecretType_CREDENTIAL:
		return "cred", nil
	case generated.SecretType_NOTE:
		return "note", nil
	default:
		return "", model.ErrUnknownType
	}
}

func (e *dbExecutor) toType(secretType string) (generated.SecretType, error) {
	switch secretType {
	case "binary":
		return generated.SecretType_BINARY, nil
	case "card":
		return generated.SecretType_CARD, nil
	case "cred":
		return generated.SecretType_CREDENTIAL, nil
	case "note":
		return generated.SecretType_NOTE, nil
	default:
		return -1, model.ErrUnknownType
	}
}
