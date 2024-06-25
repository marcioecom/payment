package repository

import (
	"database/sql"
	"time"

	"github.com/marcioecom/payment/domain/entity"
	"github.com/marcioecom/payment/domain/repository"
)

var _ repository.TransactionRepository = (*TransactionsRepositoryDB)(nil)

type TransactionsRepositoryDB struct {
	db *sql.DB
}

func NewTransactionsRepositoryDB(db *sql.DB) *TransactionsRepositoryDB {
	return &TransactionsRepositoryDB{db: db}
}

func (t *TransactionsRepositoryDB) Insert(id string, account string, amount float64, status entity.TransactionStatus, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		INSERT INTO transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		account,
		amount,
		status,
		errorMessage,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
