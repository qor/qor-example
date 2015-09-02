package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var db gorm.DB
	var err error

	if os.Getenv("DB") == "mysql" {
		if db, err = gorm.Open("mysql", "qor:qor@/qor_bookstore?parseTime=True&loc=Local"); err != nil {
			panic(err)
		}
	} else {
		if db, err = gorm.Open("postgres", "user=qor password=qor dbname=qor_bookstore sslmode=disable"); err != nil {
			panic(err)
		}
	}

	DB = &db
}
