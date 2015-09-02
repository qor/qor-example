package migrations

import (
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
)

var Admin *admin.Admin

func init() {
	db.DB.AutoMigrate(&models.Product{}, &models.ProductImage{}, &models.ColorVariation{}, &models.SizeVariation{})
	db.DB.AutoMigrate(&models.Color{}, &models.Size{}, &models.Category{})

	db.DB.AutoMigrate(&models.Address{})

	db.DB.AutoMigrate(&models.Order{}, &models.OrderItem{})

	db.DB.AutoMigrate(&models.Store{})

	db.DB.AutoMigrate(&models.User{})
}
