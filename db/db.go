package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Name))
	} else if config.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}
		Publish = publish.New(DB.Set("l10n:mode", "unscoped"))

		l10n.RegisterCallbacks(DB)
		sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
		media_library.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}
