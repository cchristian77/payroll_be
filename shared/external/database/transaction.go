package database

import (
	"context"
	"gorm.io/gorm"
)

type keys string

const txKey keys = "DBTRX"

// ConnFromContext retrieves the *gorm.DB transaction from the context or falls back to the provided default if available.
// This function is used to pass the database transaction and connection between layers.
func ConnFromContext(ctx context.Context, defaults ...*gorm.DB) (*gorm.DB, bool) {
	tx, ok := GetTxFromContext(ctx)
	if ok {
		return tx, ok
	}

	if len(defaults) == 0 {
		return nil, false
	}

	return defaults[0], true
}

// InitTx initializes a database transaction, then store it to the context.
func InitTx(ctx context.Context, db *gorm.DB) (context.Context, *gorm.DB) {
	dbTx := db.WithContext(ctx).Begin()
	return context.WithValue(ctx, txKey, dbTx), dbTx
}

// GetTxFromContext retrieves a database transaction from the provided context.
func GetTxFromContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	return tx, ok
}
