package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/qor-example/config"
)

var DB *gorm.DB

func init() {
	var err error
	var db gorm.DB

	if config.Config.DB.Adapter == "mysql" {
		db, err = gorm.Open("mysql", "qor:qor@/qor_bookstore?parseTime=True&loc=Local")
	} else if config.Config.DB.Adapter == "postgres" {
		db, err = gorm.Open("postgres", "user=qor password=qor dbname=qor_bookstore sslmode=disable")
	} else {
		db, err = gorm.Open("sqlite3", config.Config.DB.Name)
	}

	if err == nil {
		DB = &db
	} else {
		panic(err)
	}
}
