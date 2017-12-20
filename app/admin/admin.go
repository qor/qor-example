package admin

import (
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/help"
	"github.com/qor/i18n/exchange_actions"
	"github.com/qor/media/asset_manager"
	"github.com/qor/media_library"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/models/settings"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

// AssetManager asset manager
var AssetManager *admin.Resource

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
	Admin := application.Admin

	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// Add Translations
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: 1})

	// Add Worker
	Worker := SetupWorker(Admin)
	exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)
	Admin.AddResource(Worker, &admin.Config{Menu: []string{"Site Management"}})

	// Add Setting
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Shop Setting", Singleton: true})

	SetupNotification(Admin)
	SetupSEO(Admin)
	SetupDashboard(Admin)
}
