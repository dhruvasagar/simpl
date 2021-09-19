package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDBPath() string {
	boltPath := os.Getenv("BOLT_PATH")
	if boltPath == "" {
		boltPath = "pay-later-go.db"
	}
	return boltPath
}

// DB encapsulates database connection instance
type DB struct {
	db *gorm.DB
}

// New returns the database connection
func New() *DB {
	dbPath := getDBPath()
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Merchant{})
	db.AutoMigrate(&Transaction{})

	return &DB{db: db}
}
