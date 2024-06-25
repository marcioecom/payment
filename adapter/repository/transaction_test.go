package repository

import (
	"os"
	"testing"

	"github.com/marcioecom/payment/adapter/repository/fixture"
	"github.com/marcioecom/payment/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDBInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionsRepositoryDB(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
