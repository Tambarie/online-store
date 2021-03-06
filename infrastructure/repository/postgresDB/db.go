package postgresDB

import (
	"fmt"
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// Init Connecting to Postgres Database using Gorrm package
func Init() *gorm.DB {
	var dsn string
	dsn = os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Creates the Basket Table in the DB
	err = db.AutoMigrate(domain.Basket{})
	if err != nil {
		panic(err)
	}

	// Creates the Basket Products in the DB
	err = db.AutoMigrate(domain.Products{})
	if err != nil {
		panic(err)
	}

	helpers.LogEvent("INFO", fmt.Sprintf("Database Connected Successfully..."))
	return db
}
