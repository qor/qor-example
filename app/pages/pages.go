package pages

import (
	"fmt"

	"bitbucket.org/jinzhu/wcu/db"
	"github.com/qor/admin"
	"github.com/qor/page_builder"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config/application"
	"github.com/qor/qor-example/models/blogs"
	"github.com/qor/qor-example/utils"
	"github.com/qor/qor/resource"
	"github.com/qor/render"
	"github.com/qor/widget"
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
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("blog")}, "app/pages/views")}

	utils.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)
	application.Router.Get("/blog", controller.Index)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *Admin) {
	// Blog Management
	article := Admin.AddResource(&blogs.Article{}, &admin.Config{Menu: []string{"Blog Management"}})
	article.IndexAttrs("ID", "VersionName", "ScheduledStartAt", "ScheduledEndAt", "Author", "Title")

	// Setup pages
	PageBuilderWidgets := widget.New(&widget.Config{DB: db.DB})
	PageBuilderWidgets.WidgetSettingResource = Admin.NewResource(&QorWidgetSetting{}, &admin.Config{Name: "PageBuilderWidgets"})
	PageBuilderWidgets.WidgetSettingResource.NewAttrs(
		&admin.Section{
			Rows: [][]string{{"Kind"}, {"SerializableMeta"}},
		},
	)
	PageBuilderWidgets.WidgetSettingResource.AddProcessor(&resource.Processor{
		Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if widgetSetting, ok := value.(*QorWidgetSetting); ok {
				if widgetSetting.Name == "" {
					var count int
					context.GetDB().Set(admin.DisableCompositePrimaryKeyMode, "off").Model(&QorWidgetSetting{}).Count(&count)
					widgetSetting.Name = fmt.Sprintf("%v %v", utils.ToString(metaValues.Get("Kind").Value), count)
				}
			}
			return nil
		},
	})
	Admin.AddResource(PageBuilderWidgets)

	page := page_builder.New(&page_builder.Config{
		Admin:      Admin,
		PageModel:  &blogs.Page{},
		Containers: PageBuilderWidgets,
		// AdminConfig: &admin.Config{Name: "Campaign Pages or Builder", Menu: []string{"Sites & Campaign Pages"}, Priority: 2},
	})
	page.IndexAttrs("ID", "Title", "PublishLiveNow")
}
