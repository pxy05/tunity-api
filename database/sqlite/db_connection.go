package sqlite

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("tunity_test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
		return nil, err
	}
	
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	
	log.Println("Database connected successfully")
	return db, nil
}

func init() {
	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
