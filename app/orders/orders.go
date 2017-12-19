package orders

import (
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/utils"
	"github.com/qor/render"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("orders")}, "app/orders/views")}

	utils.AddFuncMapMaker(controller.View)

	application.Router.Get("/cart", controller.Cart)
	application.Router.Put("/cart", controller.UpdateCart)
	application.Router.Put("/cart/checkout", controller.Checkout)
}
