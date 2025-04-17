package db_connection

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DbOpen() {
	dsn := ConnectionString

	var err error
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	migration(Db)
}

func CloseDb() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("failed to get database object: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("failed to close database: %v", err)
	}
	fmt.Println("Database connection closed")
}
