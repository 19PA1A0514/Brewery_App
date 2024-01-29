package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	ds := "root:kamesh3939A$@tcp(127.0.0.1:3306)/brewery?parseTime=true"
	var err error
	db, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
