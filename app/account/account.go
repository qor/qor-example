package account

import (
	"github.com/go-chi/chi"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/config/auth"
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
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("account")}, "app/account/views")}

	utils.AddFuncMapMaker(controller.View)

	application.Router.Mount("/auth/", auth.Auth.NewServeMux())

	application.Router.With(auth.Authority.Authorize()).Route("/account", func(r chi.Router) {
		r.Get("/", controller.Orders)
		r.With(auth.Authority.Authorize("logged_in_half_hour")).Post("/add_user_credit", controller.AddCredit)
		r.Get("/profile", controller.Profile)
		r.Post("/profile", controller.Update)
	})
}
