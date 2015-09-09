// +build ignore

package main

import (
	"log"
	"path/filepath"

	"github.com/jinzhu/configor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

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
}{}

var Tables = []interface{}{
	&models.Category{}, &models.Color{}, &models.Size{},
}

func main() {
	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}

	truncateTables()
	createRecords()
}

func truncateTables() {
	for _, table := range Tables {
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

func createRecords() {
	createCategories()
	createColors()
	createSizes()
}

func createCategories() {
	for _, c := range Seeds.Categories {
		category := &models.Category{}
		category.Name = c.Name
		if err := db.DB.Create(&category).Error; err != nil {
			log.Fatalf("create category (%v) failure, got err %v", category, err)
		}
	}
}

func createColors() {
	for _, c := range Seeds.Colors {
		color := &models.Color{}
		color.Name = c.Name
		color.Code = c.Code
		if err := db.DB.Create(&color).Error; err != nil {
			log.Fatalf("create color (%v) failure, got err %v", color, err)
		}
	}
}

func createSizes() {
	for _, s := range Seeds.Sizes {
		size := &models.Size{}
		size.Name = s.Name
		size.Code = s.Code
		if err := db.DB.Create(&size).Error; err != nil {
			log.Fatalf("create size (%v) failure, got err %v", size, err)
		}
	}
}
