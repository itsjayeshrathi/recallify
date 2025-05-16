package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/itsjayeshrathi/recallify/internal/models"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
)

func Open() (*bun.DB, error) {
	dsn := "postgres://postgres:2859@localhost:5432/recallify?sslmode=disable"
	dbsql, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("db: Open %w", err)
	}
	fmt.Println("Connected to database")

	db := bun.NewDB(dbsql, pgdialect.New())

	db.RegisterModel((*models.LinkTag)(nil))

	return db, nil
}

func Migrate(ctx context.Context, db *bun.DB) error {
	models := []any{
		(*models.User)(nil),
		(*models.Token)(nil),
		(*models.Tag)(nil),
		(*models.Link)(nil),
		(*models.LinkTag)(nil),
	}

	for _, model := range models {
		err := db.ResetModel(ctx, model)
		if err != nil {
			return err
		}
	}
	return nil
}
