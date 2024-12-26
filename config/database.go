package config

import (
	"fmt"
	"log"
	"os"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func ConnectionDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	userDB := os.Getenv("DB_USER")
	passwordDB := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	nameDB := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, userDB, passwordDB, nameDB, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "riors_",
			SingularTable: true,
		},
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("Failed to connected database: %v", err)
	}

	db.AutoMigrate(&domain.Category{}, &domain.Product{}, &domain.DetailProduct{})
	db.AutoMigrate(&domain.User{}, &domain.Address{})
	db.AutoMigrate(&domain.Employee{})

	DB = db
}
