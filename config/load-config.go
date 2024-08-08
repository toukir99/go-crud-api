package config

import (
	"go-crud-api/db"
)

// LoadConfig initializes and returns the application configuration.
func LoadConfig() (*Config, error) {
    db, err := db.DBConnection()
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        DB: db,
    }
    return cfg, err
}
