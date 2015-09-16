package db

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grengojbo/qor-example/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
	var conStr string

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		conStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local&charset=utf8", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
		db, err = gorm.Open("mysql", conStr)
	} else if config.Config.DB.Adapter == "postgres" {
		conStr = fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name)
		db, err = gorm.Open("postgres", conStr)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		DB = &db
		Publish = publish.New(DB)
		config.Config.I18n = i18n.New(database.New(DB))

		l10n.RegisterCallbacks(DB)
		sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
