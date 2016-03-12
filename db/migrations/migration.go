package migrations

import (
	"log"

	"github.com/qor/activity"
	"github.com/qor/admin"
	"github.com/qor/media_library"
	"github.com/qor/publish"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/transition"
)

var Admin *admin.Admin

func init() {
	log.Println("Start migration ...")
	AutoMigrate(&media_library.AssetManager{})

	log.Println("model: Unit Ok")
	AutoMigrate(&models.Unit{})
	log.Println("model: Role Ok")
	AutoMigrate(&models.Role{})
	log.Println("model: Language, Phone Ok")
	AutoMigrate(&models.Language{}, &models.Phone{})
	log.Println("model: Category Ok")
	AutoMigrate(&models.Category{})

	log.Println("model: Product, ColorVariation, ColorVariationImage, SizeVariation")
	AutoMigrate(&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	log.Println("model: Color, Size, Collection")
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Collection{})

	log.Println("model: Address")
	AutoMigrate(&models.Address{})

	log.Println("model: Order, OrderItem")
	AutoMigrate(&models.Order{}, &models.OrderItem{})

	log.Println("model: Newsletter")
	AutoMigrate(&models.Newsletter{})

	log.Println("model: Store Ok")
	AutoMigrate(&models.Store{})

	log.Println("model: Setting")
	AutoMigrate(&models.Setting{})

	log.Println("model: Organization Ok")
	AutoMigrate(&models.Organization{})
	log.Println("model: User Ok")
	AutoMigrate(&models.User{})

	log.Println("model: Car Ok")
	AutoMigrate(&models.Car{})

	log.Println("model: SEOSetting Ok")
	AutoMigrate(&models.SEOSetting{})

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
