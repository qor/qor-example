package db

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/l10n"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/qor/qor-example/config"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

var (
	DB      *gorm.DB
	Publish *publish.Publish
)

func init() {
	var err error
	var conStr string

	dbConfig := config.Config.DB
	if dbConfig.Adapter == "mysql" {
		if dbConfig.Host == "localhost" {
			conStr = fmt.Sprintf("%v:%v@/%v?parseTime=True&loc=Local&charset=utf8", dbConfig.User, dbConfig.Password, dbConfig.Name)
		} else {
			conStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local&charset=utf8", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
		}
		DB, err = gorm.Open("mysql", conStr)
	} else if dbConfig.Adapter == "postgres" {
		if dbConfig.Host == "localhost" {
			conStr = fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name)
		} else {
			conStr = fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Host, dbConfig.Port)
		}
		DB, err = gorm.Open("postgres", conStr)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		DB.LogMode(dbConfig.Debug)
		Publish = publish.New(DB)
		config.Config.I18n = i18n.New(database.New(DB))

		l10n.Global = config.Config.Locale
		l10n.RegisterCallbacks(DB)
		sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
		media_library.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
