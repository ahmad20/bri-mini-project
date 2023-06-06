package database

import (
	"fmt"
	"log"

	"github.com/ahmad20/bri-mini-project/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(host, user, password, dbName, port string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Set connection pool settings if needed
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db

	migrateDatabase()

	log.Println("Database connected successfully")
	return nil
}

func migrateDatabase() {
	err := DB.AutoMigrate(
		&entities.Account{},
		&entities.Customer{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
