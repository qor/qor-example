package migrations

import (
	"github.com/qor/activity"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/db"
	"github.com/qor/transition"
)

func init() {
	AutoMigrate(&media_library.AssetManager{})

	AutoMigrate(&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Category{}, &models.Collection{})

	AutoMigrate(&models.Address{})

	AutoMigrate(&models.Order{}, &models.OrderItem{})

	AutoMigrate(&models.Store{})

	AutoMigrate(&models.Setting{})

	AutoMigrate(&models.User{})

	AutoMigrate(&models.SEOSetting{})

	AutoMigrate(&transition.StateChangeLog{})

	AutoMigrate(&activity.QorActivity{})

	AutoMigrate(&admin.QorWidgetSetting{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)

		if publish.IsPublishableModel(value) {
			db.Publish.AutoMigrate(value)
		}
	}
}
