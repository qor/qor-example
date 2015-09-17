package seeds

import (
	"math/rand"
	"path/filepath"
	"time"

	"github.com/jinzhu/configor"
	"github.com/manveru/faker"
	"github.com/qor/qor-example/db"
)

var Fake *faker.Faker

var Seeds = struct {
	Categories []struct {
		Name string
	}
	Colors []struct {
		Name string
		Code string
	}
	Sizes []struct {
		Name string
		Code string
	}
	Products []struct {
		CategoryName    string
		Name            string
		NameWithSlug    string
		Code            string
		Price           float32
		Description     string
		MadeCountry     string
		ColorVariations []struct {
			ColorName string
			Images    []struct {
				URL string
			}
		}
		SizeVariations []struct {
			SizeName string
		}
	}
	Stores []struct {
		Name      string
		Phone     string
		Email     string
		Country   string
		Zip       string
		City      string
		Region    string
		Address   string
		Latitude  float64
		Longitude float64
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

func TruncateTables(tables ...interface{}) {
	for _, table := range tables {
		if err := db.DB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		if err := db.Publish.DraftDB().DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		db.DB.AutoMigrate(table)
		db.Publish.AutoMigrate(table)
	}
}
