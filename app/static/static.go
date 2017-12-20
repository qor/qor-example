package static

import (
	"net/http"
	"strings"

	"github.com/qor/qor-example/config/application"
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
	Prefixs []string
	Handler http.Handler
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	for _, prefix := range app.Config.Prefixs {
		application.Router.Mount("/"+strings.TrimPrefix(prefix, "/"), app.Config.Handler)
	}
}
