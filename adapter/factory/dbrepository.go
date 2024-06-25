package factory

import (
	"database/sql"

	repo "github.com/marcioecom/payment/adapter/repository"
	"github.com/marcioecom/payment/domain/repository"
)

type RepositoryDatabaseFactory struct {
	db *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{db: db}
}

func (r *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionsRepositoryDB(r.db)
}
