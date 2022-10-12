package storage

import (
	"context"
	"database/sql"
	"real-estate/core/entities"
)

// The EstateStorage Repository
type EstateStorage interface {
	CreateEstates(estates []entities.Estate) (err error)
	GetEstates(mode, city, estateType string) ([]entities.Estate, error)
}

// The Estate storage Service
type Estate struct {
	db  *sql.DB
	ctx context.Context
}

// Get todo lists
func (t *Estate) CreateEstates(estates []entities.Estate) (err error) {
	tx, err := t.db.Begin()
	err = tx.Query()
	if err != nil {

	}
	err = tx.Commit()


	return err
}

// Insert a new todo
func (t *Estate) GetEstates(mode, city, estateType string) ([]entities.Estate, error) {
	return nil, nil
}

// Constructor
func New(ctx context.Context, db *sql.DB) EstateStorage {
	db.

	return &Estate{db: db, ctx: ctx}
}
