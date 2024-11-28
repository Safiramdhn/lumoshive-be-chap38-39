package database

import (
	"errors"
	"fmt"
	"log"
	"lumoshive-be-chap38-39/config"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB initializes the database connection
func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	// Configure GORM's logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Set log level to Info
			IgnoreRecordNotFoundError: false,       // Log all errors
			Colorful:                  true,        // Enable colorful output
		},
	)

	// Build the PostgreSQL connection string
	dsn, err := makePostgresString(cfg)
	if err != nil {
		log.Printf("Error building connection string: %v\n", err)
		return nil, fmt.Errorf("failed to build database connection string: %w", err)
	}

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Printf("Error connecting to the database: %v\n", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error retrieving raw DB connection: %v\n", err)
		return nil, fmt.Errorf("failed to retrieve raw database connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Printf("Error testing database connection: %v\n", err)
		return nil, fmt.Errorf("database connection test failed: %w", err)
	}

	// Setup database (migrations and initialization)
	if err := setupDatabase(db); err != nil {
		log.Printf("Error setting up the database: %v\n", err)
		return nil, fmt.Errorf("failed to set up database: %w", err)
	}

	log.Println("Database connected and setup successfully.")
	return db, nil
}

// makePostgresString ensures the connection string is constructed correctly
func makePostgresString(cfg config.Config) (string, error) {
	if cfg.DB.DBUser == "" || cfg.DB.DBName == "" {
		return "", errors.New("missing database configuration (user or name)")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, cfg.DB.DBPassword, cfg.DB.DBName,
	), nil
}

// setupDatabase performs database migration and initialization
func setupDatabase(db *gorm.DB) error {
	if err := migrateDatabase(db); err != nil {
		log.Printf("Error migrating database: %v\n", err)
		return err
	}

	if err := initiateShippingData(db); err != nil {
		log.Printf("Error initializing shipping data: %v\n", err)
		return err
	}

	log.Println("Database setup completed successfully.")
	return nil
}
