package feature

import (
	"path/filepath"

	// "github.com/apertoire/mlog"
	"github.com/jinzhu/configor"
	"github.com/qor/qor-example/db/seeds"
)

var (
	// fake = seeds.Fake
	// truncateTables = seeds.TruncateTables

	Seeds = seeds.Seeds
)

func Load() {
	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		// mlog.Error("feature init Fatal :(")
		mlog.Fatal(err)
	}
}

func CreateLanguages() {
	for _, c := range Seeds.Languages {
    fmt.Println(c)
		mlog.Trace("Languages: %v", c)
		// 	language := models.Language{}
		// 	language.Name = c.Name
		// 	if err := db.DB.Where(models.Language{Name: c.Name}).FirstOrCreate(&language).Error; err != nil {
		// 		mlog.Error(err)
		// 	}
		// 	//     if err := db.DB.Create(&language).Error; err != nil {
		// 	//       log.Fatalf("create language (%v) failure, got err %v", language, err)
		// 	//     }
	}
}
