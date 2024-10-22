package database

import (
	"context"
	"fmt"
	"sync"

	"log/slog"

	"github.com/exaring/otelpgx"
	"github.com/felipeversiane/go-otel/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	once sync.Once
	db   DatabaseInterface
)

type database struct {
	db     *pgxpool.Pool
	config config.DatabaseConfig
}

type DatabaseInterface interface {
	GetDB() *pgxpool.Pool
	Ping(ctx context.Context) error
	Close()
}

func NewDatabaseConnection(ctx context.Context, config config.DatabaseConfig) DatabaseInterface {
	once.Do(func() {
		dsn := getConnectionString(config)
		poolConfig, parseErr := pgxpool.ParseConfig(dsn)
		if parseErr != nil {
			slog.Error("failed to parse database config", slog.String("dsn", dsn), slog.Any("error", parseErr))
			return
		}
		poolConfig.ConnConfig.Tracer = otelpgx.NewTracer()

		pool, poolErr := pgxpool.NewWithConfig(ctx, poolConfig)
		if poolErr != nil {
			slog.Error("failed to create database pool", slog.Any("error", poolErr))
			return
		}

		db = &database{
			db:     pool,
			config: config,
		}

		if pingErr := db.Ping(ctx); pingErr != nil {
			slog.Error("database ping failed", slog.Any("error", pingErr))
		}
	})

	return db
}

func (d *database) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *database) Close() {
	d.db.Close()
}

func (d *database) GetDB() *pgxpool.Pool {
	return d.db
}

func getConnectionString(config config.DatabaseConfig) string {
	user := config.User
	password := config.Password
	dbname := config.Name
	dbport := config.Port
	dbhost := config.Host

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", user, password, dbname, dbport, dbhost)
	return dsn
}
