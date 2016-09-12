package admin

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/l10n"
	"github.com/qor/media_library"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/resource"
	"github.com/qor/sorting"
	"github.com/qor/widget"
)

var Widgets *widget.Widgets

type QorWidgetSetting struct {
	widget.QorWidgetSetting
	l10n.Locale
	// publish.Status
	// DeletedAt *time.Time
}

func initWidgets() {
	if Widgets == nil {
		Widgets = widget.New(&widget.Config{DB: db.DB})
		Widgets.WidgetSettingResource = Admin.AddResource(&QorWidgetSetting{}, &admin.Config{Menu: []string{"Site Management"}, Priority: 3})
		Widgets.WidgetSettingResource.Meta(&admin.Meta{Name: "Name", FormattedValuer: func(widget interface{}, ctx *qor.Context) interface{} {
			setting := widget.(*QorWidgetSetting)
			return template.HTML(`<img src="/images/Widget-` + setting.WidgetType + `.png" width="80" height="35" style="margin-right: 12px;"/><span>` + setting.Name + `</span>`)
		}})

		Widgets.RegisterScope(&widget.Scope{
			Name: "From Google",
			Visible: func(context *widget.Context) bool {
				if request, ok := context.Get("Request"); ok {
					_, ok := request.(*http.Request).URL.Query()["from_google"]
					return ok
				}
				return false
			},
		})

		Admin.AddResource(Widgets)

		// Top Banner
		type bannerArgument struct {
			Title           string
			ButtonTitle     string
			Link            string
			BackgroundImage media_library.FileSystem
			Logo            media_library.FileSystem
		}

		Widgets.RegisterWidget(&widget.Widget{
			Name:      "NormalBanner",
			Templates: []string{"banner", "banner2"},
			Setting:   Admin.NewResource(&bannerArgument{}),
			Context: func(context *widget.Context, setting interface{}) *widget.Context {
				context.Options["Setting"] = setting
				return context
			},
		})

		type slideImage struct {
			Title string
			Image media_library.FileSystem
		}

		type slideShowArgument struct {
			SlideImages []slideImage
		}
		slideShowResource := Admin.NewResource(&slideShowArgument{})
		slideShowResource.AddProcessor(func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if slides, ok := value.(*slideShowArgument); ok {
				for _, slide := range slides.SlideImages {
					if slide.Title == "" {
						return errors.New("slide title is blank")
					}
				}
			}
			return nil
		})

		Widgets.RegisterWidget(&widget.Widget{
			Name:      "SlideShow",
			Templates: []string{"slideshow"},
			Setting:   slideShowResource,
			Context: func(context *widget.Context, setting interface{}) *widget.Context {
				context.Options["Setting"] = setting
				return context
			},
		})

		Widgets.RegisterWidgetsGroup(&widget.WidgetsGroup{
			Name:    "Banner",
			Widgets: []string{"NormalBanner", "SlideShow"},
		})

		// selected Products
		type selectedProductsArgument struct {
			Products       []string
			ProductsSorter sorting.SortableCollection
		}
		selectedProductsResource := Admin.NewResource(&selectedProductsArgument{})
		selectedProductsResource.Meta(&admin.Meta{Name: "Products", Type: "select_many", Collection: func(value interface{}, context *qor.Context) [][]string {
			var collectionValues [][]string
			var products []*models.Product
			context.GetDB().Find(&products)
			for _, product := range products {
				collectionValues = append(collectionValues, []string{fmt.Sprintf("%v", product.ID), product.Name})
			}
			return collectionValues
		}})

		Widgets.RegisterWidget(&widget.Widget{
			Name:      "Products",
			Templates: []string{"products"},
			Setting:   selectedProductsResource,
			Context: func(context *widget.Context, setting interface{}) *widget.Context {
				if setting != nil {
					var products []*models.Product
					context.GetDB().Limit(9).Preload("ColorVariations").Where("id IN (?)", setting.(*selectedProductsArgument).Products).Find(&products)
					setting.(*selectedProductsArgument).ProductsSorter.Sort(&products)
					context.Options["Products"] = products
				}
				return context
			},
		})
	}
}
