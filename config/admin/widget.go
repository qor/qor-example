package admin

import (
	"github.com/qor/media_library"
	"github.com/qor/qor-example/db"
	"github.com/qor/widget"
)

var Widget *widget.WidgetInstance

func init() {
	Widget = widget.New(&widget.Config{DB: db.DB})
	Admin.AddResource(Widget)

	type ImageStorage struct{ media_library.FileSystem }
	type bannerArgument struct {
		Title           string
		ButtonTitle     string
		Link            string
		BackgroundImage ImageStorage `sql:"type:varchar(4096)"`
		Logo            ImageStorage `sql:"type:varchar(4096)"`
	}

	Widget.RegisterWidget(&widget.Widget{
		Name:     "Banner",
		Template: "banner",
		Setting:  Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				argument := setting.(*bannerArgument)
				context.Options["Title"] = argument.Title
				context.Options["ButtonTitle"] = argument.ButtonTitle
				context.Options["Link"] = argument.Link
				context.Options["BackgroundUrl"] = argument.BackgroundImage.URL()
				context.Options["Logo"] = argument.Logo.URL()
			}
			return context
		},
	})
}
