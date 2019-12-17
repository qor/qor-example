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
		cfg.prefix = "/graphql"
	}
	if cfg.Version == "" {
		cfg.Version = "0.0.1-SNAPSHOT"
	}
	return &App{Config: cfg}
}

// Config home config struct
type Config struct {
	Prefix string
	Version string
}

// ConfigureApplication configure appliction
func (app App) ConfigureApplication(application *application.Application) {
	Graphql := admin.New(&qor.Config{DB: db.DB})

//	Graphql.AddResource(&products.Product{})
//	Graphql.AddResource(&orders.Orders{})
//	Graphql.AddResource(&users.User{})
//	Graphql.AddResource(&products.Category{})

	app.Router.Mount(app.Config.Prefix, Graphql.NewServerMux(app.COnfig.Prefix))
}/
