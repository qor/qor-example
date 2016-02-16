package migrations

import (
	"log"

	"github.com/qor/activity"
	"github.com/qor/media_library"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/publish"
	"github.com/qor/transition"
)

var Admin *admin.Admin

func init() {
	log.Println("Start migration ...")
	AutoMigrate(&media_library.AssetManager{})

	log.Println("model: Unit")
	AutoMigrate(&models.Unit{})

	log.Println("model: Product, ColorVariation, ColorVariationImage, SizeVariation")
	AutoMigrate(&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	log.Println("model: Color, Size, Category, Collection")
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Category{}, &models.Collection{})

	log.Println("model: Address")
	AutoMigrate(&models.Address{})

	log.Println("model: Order, OrderItem")
	AutoMigrate(&models.Order{}, &models.OrderItem{})

	log.Println("model: Newsletter")
	AutoMigrate(&models.Newsletter{})

	log.Println("model: Store")
	AutoMigrate(&models.Store{})

	log.Println("model: Setting")
	AutoMigrate(&models.Setting{})

	log.Println("model: Role, Language, Phone")
	AutoMigrate(&models.Role{}, &models.Language{}, models.Phone{})
	log.Println("model: Organization")
	AutoMigrate(&models.Organization{})
	log.Println("model: User")
	AutoMigrate(&models.User{})

	log.Println("model: Car")
	AutoMigrate(&models.Car{})

	log.Println("model: Seo")
	AutoMigrate(&models.Seo{})

	log.Println("model: StateChangeLog")
	AutoMigrate(&transition.StateChangeLog{})

	log.Println("model: QorActivity")
	AutoMigrate(&activity.QorActivity{})

	log.Println("Finish migration :)")
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)

		if publish.IsPublishableModel(value) {
			db.Publish.AutoMigrate(value)
		}
	}
}
