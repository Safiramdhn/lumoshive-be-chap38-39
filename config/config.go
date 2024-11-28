package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	AppPort  string   `mapstructure:"port"`
	AppDebug bool     `mapstructure:"debug"`
	DB       DBConfig `mapstructure:"db"`
}

// DBConfig holds the database configuration
type DBConfig struct {
	DBHost     string `mapstructure:"host"`
	DBPort     string `mapstructure:"port"`
	DBUser     string `mapstructure:"user"`
	DBPassword string `mapstructure:"password"`
	DBName     string `mapstructure:"name"`
}

// LoadConfig initializes and reads configuration using Viper
func LoadConfig() (Config, error) {
	v := viper.New()

	// Configure Viper for environment variables and default values
	setupViper(v)

	// Attempt to read the config file
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Config file not found: %v. Using environment variables and defaults.", err)
	}

	// Unmarshal the configuration into the Config struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	log.Println("Configuration loaded successfully.")
	return config, nil
}

// setupViper configures Viper with default values and environment variables
func setupViper(v *viper.Viper) {
	// Set the config file and type
	v.SetConfigFile(".env")
	v.SetConfigType("dotenv")

	// Set default values
	v.SetDefault("port", "8080")
	v.SetDefault("debug", true)
	v.SetDefault("db.user", "postgres")
	v.SetDefault("db.password", "postgres")
	v.SetDefault("db.name", "shipping-ecommerce-dummy")
	v.SetDefault("db.host", "localhost")
	v.SetDefault("db.port", "5432")

	// Enable environment variable override
	v.AutomaticEnv()

	// Map environment variables explicitly (if keys are different)
	v.BindEnv("port", "PORT")
	v.BindEnv("debug", "DEBUG")
	v.BindEnv("db.name", "DATABASE_NAME")
	v.BindEnv("db.host", "DATABASE_HOST")
	v.BindEnv("db.user", "DATABASE_USERNAME")
	v.BindEnv("db.password", "DATABASE_PASSWORD")
	v.BindEnv("db.port", "DATABASE_PORT")
}
