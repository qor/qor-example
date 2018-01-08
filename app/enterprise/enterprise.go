// +build enterprise

package enterprise

import (
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
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {
	SetupPromotion(application.Admin)
	SetupMicrosite(application.Admin)
}
