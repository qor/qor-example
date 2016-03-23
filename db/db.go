package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	if config.Config.DB.Adapter == "mysql" {
		if config.Config.DB.Host == "localhost" {
			DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?parseTime=True&loc=Local&charset=utf8", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Name))
		} else {
			DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local&charset=utf8", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Name))
		}
	} else if config.Config.DB.Adapter == "postgres" {
		if config.Config.DB.Host == "localhost" {
			DB, err = gorm.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Name))
		} else {
			DB, err = gorm.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Name, config.Config.DB.Host, config.Config.DB.Port))
		}
	} else if config.Config.DB.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", config.Config.DB.Name)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if debug := os.Getenv("DEBUG"); len(debug) > 0 {
			DB.LogMode(false)
		} else {
			DB.LogMode(config.Config.DB.Debug)
		}
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
