package admin

import (
	"github.com/qor/action_bar"
	"github.com/qor/qor-example/config/application"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

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
	AddNotification(application.Admin)
}
