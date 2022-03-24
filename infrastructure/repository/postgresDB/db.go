package postgresDB

import (
	"fmt"
	domain "github.com/Tambarie/online-store/domain/store"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Init() *gorm.DB {
	var dsn string
	dsn = os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = fmt.Sprintf("host=%v user=%v dbname=%v port=%v sslmode=%v TimeZone=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("TIMEZONE"))
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(domain.Products{})
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected Successfully...")
	return db
}
