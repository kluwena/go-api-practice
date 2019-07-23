package data

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type key int

const (
	transactionKey key = 0
)

// Queryer represents the data commands interface
type Queryer interface {
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}

// Manager represents the manager to manage the data consistency
type Manager struct {
	db *sqlx.DB
}

func TransactionFromContext(ctx context.Context, queryer Queryer) Queryer {
	tx, ok := ctx.Value(transactionKey).(Queryer)
	if !ok {
		return queryer
	}

	return tx
}

// RunInTransaction runs the f with the transaction queryable inside the context
func (m *Manager) RunInTransaction(ctx context.Context, f func(context.Context) error) error {
	_, err := m.db.Beginx()
	if err != nil {
		return err
	}

	return nil
}

// NewManager creates a new manager
func NewManager(db *sqlx.DB) *Manager {
	return &Manager{
		db: db,
	}
}
