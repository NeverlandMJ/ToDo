package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/NeverlandMJ/ToDo/user-service/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Connect connects to database and migrates if necessary
func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"pgx",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB,
		),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate1: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", cfg.PostgresMigrationsPath),
		"postgres",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate2: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate3: %v", err)
	}

	return db, nil
}
