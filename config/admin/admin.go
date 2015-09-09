package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
)

var Admin *admin.Admin
var Countries = []string{"China", "Japan", "USA"}

func init() {
	Admin = admin.New(&qor.Config{DB: db.Publish.DraftDB()})
	Admin.SetAuth(Auth{})

	product := Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})
	product.Meta(&admin.Meta{Name: "MadeCountry", Type: "select_one", Collection: Countries})
	product.Meta(&admin.Meta{Name: "Description", Type: "rich_editor", Resource: Admin.AddResource(&admin.AssetManager{}, &admin.Config{Invisible: true})})
	product.IndexAttrs("-ColorVariations")
	for _, country := range Countries {
		var country = country
		product.Scope(&admin.Scope{
			Name:  country,
			Group: "Made Country",
			Handle: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
				return db.Where("made_country = ?", country)
			},
		})
	}

	Admin.AddResource(&models.Color{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Size{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})

	Admin.AddResource(&models.Order{}, &admin.Config{Menu: []string{"Order Management"}})

	store := Admin.AddResource(&models.Store{}, &admin.Config{Menu: []string{"Store Management"}})
	store.IndexAttrs("-Latitude", "-Longitude")

	Admin.AddResource(config.Config.I18n, &admin.Config{Menu: []string{"Site Management"}})

	user := Admin.AddResource(&models.User{})
	user.IndexAttrs("ID", "Email", "Name", "Gender", "Role")

	Admin.AddResource(db.Publish)
}
