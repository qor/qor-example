package migrations

import (
	"github.com/qor/activity"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/banner_editor"
	"github.com/qor/help"
	"github.com/qor/media/asset_manager"
	"github.com/qor/qor-example/app/blogs"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/app/orders"
	"github.com/qor/qor-example/app/products"
	"github.com/qor/qor-example/app/stores"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/seo"
	"github.com/qor/qor-example/db"
	"github.com/qor/transition"
)

func init() {
	AutoMigrate(&asset_manager.AssetManager{})

	AutoMigrate(&products.Product{}, &products.ProductVariation{}, &products.ProductImage{}, &products.ColorVariation{}, &products.ColorVariationImage{}, &models.SizeVariation{})
	AutoMigrate(&products.Color{}, &products.Size{}, &products.Material{}, &products.Category{}, &products.Collection{})

	AutoMigrate(&users.User{}, &users.Address{})

	AutoMigrate(&orders.Order{}, &orders.OrderItem{})

	AutoMigrate(&orders.DeliveryMethod{})

	AutoMigrate(&stores.Store{})

	AutoMigrate(&settings.Setting{}, &settings.MediaLibrary{})

	AutoMigrate(&transition.StateChangeLog{})

	AutoMigrate(&activity.QorActivity{})

	AutoMigrate(&admin.QorWidgetSetting{})

	AutoMigrate(&blogs.Page{}, &blogs.Article{})

	AutoMigrate(&seo.MySEOSetting{})

	AutoMigrate(&help.QorHelpEntry{})

	AutoMigrate(&auth_identity.AuthIdentity{})

	AutoMigrate(&banner_editor.QorBannerEditorSetting{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
