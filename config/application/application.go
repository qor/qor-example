package application

import (
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// MicroAppInterface micro app interface
type MicroAppInterface interface {
	ConfigureApplication(*Application)
}

// Application main application
type Application struct {
	*Config
}

// Config application config
type Config struct {
	Router *chi.Mux
	Admin  *admin.Admin
	DB     *gorm.DB
}

// New new application
func New(cfg *Config) *Application {
	return &Application{
		Config: cfg,
	}
}

// Use mount router into micro app
func (application *Application) Use(app MicroAppInterface) {
	app.ConfigureApplication(application)
}
