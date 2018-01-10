package stores

import (
	"strings"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/models/stores"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
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
func (app App) ConfigureApplication(application *application.Application) {
	app.ConfigureAdmin(application.Admin)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	// Add Store
	store := Admin.AddResource(&stores.Store{}, &admin.Config{Menu: []string{"Store Management"}})
	store.Meta(&admin.Meta{Name: "Owner", Type: "single_edit"})
	store.AddValidator(&resource.Validator{
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if meta := metaValues.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "Name can't be blank")
				}
			}
			return nil
		},
	})
}
