package postgres

import (
	"context"
	"fmt"

	"github.com/gogapopp/L0/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func New(config *config.Config) (*repository, error) {
	const op = "repositry.postgres.New"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDb)

	ctx := context.Background()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) MigrateUp(config *config.Config) error {
	const op = "repositry.postgres.MigrateUp"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDb)

	m, err := migrate.New(config.MogrationDir, dsn)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *repository) MigrateDown(config *config.Config) error {
	const op = "repositry.postgres.MigrateDown"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDb)

	m, err := migrate.New(config.MogrationDir, dsn)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
