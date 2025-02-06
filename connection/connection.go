package connection

import (
	"context"
	"fmt"
	"os"
	"rental/app"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(ctx context.Context, log *logrus.Entry) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("Gagal membuka koneksi ke PostgreSQL: %v", err)
		return nil
	}

	err = db.AutoMigrate(
		&app.Customer{},
		&app.Car{},
		&app.Booking{},
		&app.Membership{},
		&app.Driver{},
		&app.DriverIncentive{},
		&app.BookingType{},
	)
	if err != nil {
		log.Errorf("Gagal melakukan migrasi: %v", err)
		return nil
	}

	log.Info("Koneksi ke PostgreSQL berhasil")
	return db
}
