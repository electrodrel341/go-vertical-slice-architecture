package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
)

type DbConn struct {
	DbPool *sql.DB
}

// InitPool initializes the database connection pool and runs migrations.
func InitPool(config *DbConfig, logger zerolog.Logger) *DbConn {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.dbHost, config.dbPort, config.dbUser, config.dbPassword, config.dbName)

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		logger.Panic().Err(err).Msg("failed to open database connection")
	}

	if err := db.Ping(); err != nil {
		logger.Panic().Err(err).Msg("failed to ping database")
	}

	if config.dbMigrationPath != nil {
		if err := runMigration(db, *config.dbMigrationPath, logger); err != nil {
			logger.Panic().Err(err).Msg("failed to run database migration")
		}
	}

	logger.Info().Str("source", "database").Msg("Database connected successfully")

	return &DbConn{DbPool: db}
}

func runMigration(db *sql.DB, path string, logger zerolog.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", path),
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	logger.Info().Str("source", "migration").Msg("Running database migrations")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	logger.Info().Str("source", "migration").Msg("Database migrations applied successfully")

	return nil
}
