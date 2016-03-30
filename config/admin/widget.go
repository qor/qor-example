package admin

import (
	"fmt"
	"github.com/qor/admin"
	"github.com/qor/media_library"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/widget"
)

var Widget *widget.WidgetInstance

func init() {
	Widget = widget.New(&widget.Config{DB: db.DB})
	Admin.AddResource(Widget)

	// Top Banner
	type ImageStorage struct{ media_library.FileSystem }
	type bannerArgument struct {
		Title           string
		ButtonTitle     string
		Link            string
		BackgroundImage ImageStorage `sql:"type:varchar(4096)"`
		Logo            ImageStorage `sql:"type:varchar(4096)"`
	}

	Widget.RegisterWidget(&widget.Widget{
		Name:     "Banner",
		Template: "banner",
		Setting:  Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				argument := setting.(*bannerArgument)
				context.Options["Title"] = argument.Title
				context.Options["ButtonTitle"] = argument.ButtonTitle
				context.Options["Link"] = argument.Link
				context.Options["BackgroundUrl"] = argument.BackgroundImage.URL()
				context.Options["Logo"] = argument.Logo.URL()
			}
			return context
		},
	})

	// Feature Products
	type featureProductsArgument struct {
		Products []string
	}

	featureProductsResouce := Admin.NewResource(&featureProductsArgument{})
	featureProductsResouce.Meta(&admin.Meta{Name: "Products", Type: "select_many", Collection: func(value interface{}, context *qor.Context) [][]string {
		var collectionValues [][]string
		var products []*models.Product
		db.DB.Find(&products)
		for _, product := range products {
			collectionValues = append(collectionValues, []string{fmt.Sprintf("%v", product.ID), product.Name})
		}
		return collectionValues
	}})
	Widget.RegisterWidget(&widget.Widget{
		Name:     "Products",
		Template: "products",
		Setting:  featureProductsResouce,
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				var products []*models.Product
				db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Where("id IN (?)", setting.(*featureProductsArgument).Products).Find(&products)
				context.Options["Products"] = products
			}
			return context
		},
	})
}
