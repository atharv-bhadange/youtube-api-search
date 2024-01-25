package db

import (
	"fmt"
	"log"

	"github.com/atharv-bhadange/youtube-api-search/models"
	"github.com/atharv-bhadange/youtube-api-search/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DbConn is the connection handle for the database
	DbConn *gorm.DB
)

// Init initializes the database
func Init() error {
	var err error

	// Get database connection string
	dbConnectionString, err := utils.GetDbConnectionString()
	if err != nil {
		return fmt.Errorf("error while getting db connection string: %v", err)
	}

	// Connect to database
	DbConn, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error while connecting to database: %v", err)
	}

	// Migrate the schema
	err = DbConn.AutoMigrate(&models.VideoDetails{})
	if err != nil {
		return fmt.Errorf("error while migrating database: %v", err)
	}

	log.Println("Database migrated")

	return nil
}

// Close database connection
func Close() {
	sqlDB, err := DbConn.DB()
	if err != nil {
		log.Println("Error while closing database connection", err)
	}

	sqlDB.Close()
}
