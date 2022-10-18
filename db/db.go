package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var defaultMigrationsDir = "../migrations"

type DB struct {
	*pgx.Conn
}

func New(ctx context.Context, connStr string) (*DB, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &DB{conn}, nil
}

func Migrate(ctx context.Context, connStr string) error {
	var migrationsDir string
	if md := os.Getenv("MIGRATIONS_DIR"); len(md) == 0 {
		migrationsDir = defaultMigrationsDir
	}

	conn, err := goose.OpenDBWithDriver("postgres", connStr)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		return err
	}

	if err := goose.Up(conn, migrationsDir); err != nil {
		return err
	}

	return nil
}
