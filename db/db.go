package db

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor/l10n"
	"github.com/qor/qor/publish"
)

var (
	DB           gorm.DB
	Pub          *publish.Publish
	ProductionDB *gorm.DB
	StagingDB    *gorm.DB
)

func init() {
	var err error

	if os.Getenv("DB") == "mysql" {
		if DB, err = gorm.Open("mysql", "qor:qor@/qor_bookstore?parseTime=True&loc=Local"); err != nil {
			panic(err)
		}
	} else {
		if DB, err = gorm.Open("postgres", "user=qor password=qor dbname=qor_bookstore sslmode=disable"); err != nil {
			panic(err)
		}
	}

	DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.User{})
	DB.LogMode(true)

	Pub = publish.New(&DB)
	Pub.AutoMigrate(&models.Author{}, &models.Book{})

	StagingDB = Pub.DraftDB()         // Draft resources are saved here
	ProductionDB = Pub.ProductionDB() // Published resources are saved here

	l10n.Global = "en-US"
	l10n.RegisterCallbacks(&DB)
}
