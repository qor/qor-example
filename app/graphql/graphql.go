package graphql

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/config/db"
//        "github.com/qor/qor-example/lib/models/orders"
//        "github.com/qor/qor-example/lib/models/products"
//        "github.com/qor/qor-example/lib/models/users"
)

// New new home app
func New(cfg *Config) *App {
	if cfg.Prefix == "" {
		cfg.Prefix = "/api"
	}
	if cfg.Version == "" {
		cfg.Version = "0.0.1-SNAPSHOT"
	}
	return &App{Config: cfg}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix string
	Version string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	GraphQL := admin.New(&qor.Config{DB: db.DB})
/*
	Product := GraphQL.AddResource(&products.Product{})

        ColorVariationMeta := Product.Meta(&admin.Meta{Name: "ColorVariations"})
        ColorVariation := ColorVariationMeta.Resource
        ColorVariation.IndexAttrs("ID", "Color", "Images", "SizeVariations")
        ColorVariation.ShowAttrs("Color", "Images", "SizeVariations")

        SizeVariationMeta := ColorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
        SizeVariation := SizeVariationMeta.Resource
        SizeVariation.IndexAttrs("ID", "Size", "AvailableQuantity")
        SizeVariation.ShowAttrs("ID", "Size", "AvailableQuantity")

	GraphQL.AddResource(&orders.Order{})
	GraphQL.AddResource(&users.User{})
//	  User := GraphQL.AddResource(&users.User{})
//	  userOrders, _ := User.AddSubResource("Orders")
//	  userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})
	GraphQL.AddResource(&products.Category{})
*/
	application.Router.Mount(app.Config.Prefix, GraphQL.NewServeMux(app.Config.Prefix))
}
