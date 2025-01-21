package sqllite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/kareem717/k7-cbo/internal/storage"
	"github.com/kareem717/k7-cbo/internal/storage/sqllite/company"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// DB wraps the sql.DB connection
type repository struct {
	db *bun.DB
}

type config struct {
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type configOptFunc func(cfg *config)

func withMaxIdleConns(maxIdleConns int) configOptFunc {
	return func(cfg *config) {
		cfg.MaxIdleConns = maxIdleConns
	}
}

func withConnMaxLifetime(connMaxLifetime time.Duration) configOptFunc {
	return func(cfg *config) {
		cfg.ConnMaxLifetime = connMaxLifetime
	}
}

// NewRepository creates a repository implementation using sqlite
func NewRepository(opts ...configOptFunc) (storage.Repository, error) {
	cfg := &config{
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Hour,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	sqldb, err := sql.Open(sqliteshim.ShimName, "k7-cbo.db")
	if err != nil {
		panic(err)
	}

	sqldb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqldb.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	return &repository{db}, nil
}

func (r *repository) Company() storage.CompanyRepository {
	return company.NewCompanyRepository(r.db)
}

const (
	migrationsDir = "migrations"
)

func (r *repository) Migrate(ctx context.Context) error {
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(r.db.DB, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
func (r *repository) HealthCheck(ctx context.Context) error {
	return r.db.Ping()
}

func (r *repository) Shutdown(ctx context.Context) error {
	return r.db.Close()
}
