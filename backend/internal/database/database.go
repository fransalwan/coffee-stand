package database

import (
	"fmt"
	"log"
	"os"

	"backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system env")
	}

	// Format DSN PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("✅ Database connected successfully!")

	// Panggil Auto-Migrate dan Seeder
	RunMigrations()
	SeedDatabase()
}

// RunMigrations menangani pembuatan/penyesuaian tabel
func RunMigrations() {
	// Tambahkan semua model lu di sini. Nanti kalau ada model Customer/Order, tinggal masukin ke slice ini.
	err := DB.AutoMigrate(
		&models.MenuItem{},
	)

	if err != nil {
		log.Fatal("❌ Failed to auto-migrate database:", err)
	}
	log.Println("✅ Database migrated successfully!")
}
