package database

import (
	"fmt"
	"sharpnr-engine/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // ← blank import, registers the driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // ← blank import, registers file source
)

func MigrateDB(cfg config.DatabaseConfig) error {
	// Must use the same sslmode as ConnectDB — a migrator that quietly falls
	// back to an unencrypted connection is the same leak, just at startup.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode,
	)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return fmt.Errorf("failed to init migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
