package admin

import (
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
)

var Admin *admin.Admin

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB})

	product := Admin.AddResource(&models.Product{}, &admin.Config{Menu: []string{"Product Management"}})

	product.Meta(&admin.Meta{Name: "MadeCountry", Type: "select_one", Collection: []string{"China", "Japan", "USA"}})
	product.Meta(&admin.Meta{Name: "Description", Type: "rich_editor", Resource: Admin.AddResource(&admin.AssetManager{}, &admin.Config{Invisible: true})})

	Admin.AddResource(&models.Color{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Size{}, &admin.Config{Menu: []string{"Product Management"}})
	Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}})

	Admin.AddResource(&models.Order{}, &admin.Config{Menu: []string{"Order Management"}})

	Admin.AddResource(&models.Store{}, &admin.Config{Menu: []string{"Store Management"}})

	Admin.AddResource(config.Config.I18n, &admin.Config{Menu: []string{"Site Management"}})

	Admin.AddResource(&models.User{})

	Admin.AddResource(db.Publish)
}
