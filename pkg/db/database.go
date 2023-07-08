package db

import (
	"MAXPUMP1/pkg/domain/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=ashiq password=ashiq123 dbname=maxpump port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("eror creating database")
		return nil
	}

	db.AutoMigrate(entity.User{})

	DB = db

	return db
}
