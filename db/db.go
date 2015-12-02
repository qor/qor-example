package db

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/media_library"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor/i18n"
	"github.com/qor/qor/i18n/backends/database"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
	"github.com/qor/qor/sorting"
	"github.com/qor/qor/validations"
)

var (
	DB      *gorm.DB
	Publish *publish.Publish
)

func init() {
	var err error
	var db gorm.DB

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Name))
	} else if config.Config.DB.Adapter == "postgres" {
		db, err = gorm.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		DB = &db
		// DB.LogMode(true)
		Publish = publish.New(DB)
		config.Config.I18n = i18n.New(database.New(DB))

		l10n.RegisterCallbacks(DB)
		sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
		media_library.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
