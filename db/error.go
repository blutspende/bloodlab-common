package db

import (
	"errors"

	pgx "github.com/jackc/pgconn"
	"github.com/lib/pq"
)

const (
	DuplicateKeyErrorCode        = pq.ErrorCode("23505")
	ForeignKeyViolationErrorCode = pq.ErrorCode("23503")
)

var (
	ErrBeginTransactionTransactionFailed    = errors.New("begin transaction failed")
	ErrCommitTransactionTransactionFailed   = errors.New("commit transaction failed")
	ErrRollbackTransactionTransactionFailed = errors.New("revert transaction failed")
	ErrCommitWithoutTransaction             = errors.New("invalid transaction, can not perform commit without transaction")
	ErrRollbackWithoutTransaction           = errors.New("invalid transaction, can not perform rollback without transaction")
)

func IsErrorCode(err error, errcode pq.ErrorCode) bool {
	pgErr, ok := err.(*pq.Error)
	if ok {
		return pgErr.Code == errcode
	}

	pgxErr, ok := err.(*pgx.PgError)
	if ok {
		currentCode := pq.ErrorCode(pgxErr.Code)
		return currentCode == errcode
	}

	return false
}

func TryCastErrorToPgError(err error) any {
	pgErr, ok := err.(*pq.Error)
	if ok {
		return pgErr
	}
	pgxErr, ok := err.(*pgx.PgError)
	if ok {
		return pgxErr
	}
	return err.Error()
}
