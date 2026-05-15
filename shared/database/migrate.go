package database

import (
	"fmt"
	"sharpnr-engine/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // ← blank import, registers the driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // ← blank import, registers file source

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {

	if err := createSchemas(db); err != nil {
		return err
	}

	migrations := []func(*gorm.DB) error{

		migrateCore,
		migrateOrgProperty,
		migrateLegalEntity,
		migrateEmployee,
		migrateHotel,
		migrateRestaurant,
		migrateBar,
		migrateKitchen,
		migrateHospitalityMenu,
		migrateAccounting,
		migrateSetupWizard,
		migrateInvoicingEngine,
		migrateHospitalityOrderEngine,
		migrateSeroEngine,
		mingrateCustomerEngine,
		migrateInvoicingEngine,
		migratePaymentsEngine,
		migrateInventoryEngine,
		migrateOrgSuppliers,
	}

	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			return err
		}
	}

	return nil
}

func MigrateNewDB(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
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

