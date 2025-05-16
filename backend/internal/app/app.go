package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/itsjayeshrathi/recallify/internal/store"
	"github.com/uptrace/bun"
)

type Application struct {
	logger *log.Logger
	DB     *bun.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	err = store.Migrate(context.Background(), pgDB)

	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		logger: logger,
		DB:     pgDB,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is availabe\n")
}
