package application

import (
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/assetfs"
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
	Router  *chi.Mux
	AssetFS assetfs.Interface
	Admin   *admin.Admin
	DB      *gorm.DB
}

// New new application
func New(cfg *Config) *Application {
	if cfg.AssetFS == nil {
		cfg.AssetFS = assetfs.AssetFS()
	}

	return &Application{
		Config: cfg,
	}
}

// Use mount router into micro app
func (application *Application) Use(app MicroAppInterface) {
	app.ConfigureApplication(application)
}
