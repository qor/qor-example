package admin

import (
	"github.com/qor/qor-example/db"
	"github.com/qor/widget"
)

var Widget *widget.WidgetInstance

func init() {
	Widget = widget.New(&widget.Config{DB: db.DB})
	Admin.AddResource(Widget)
	type bannerArgument struct {
		Title    string
		SubTitle string
	}

	Widget.RegisterWidget(&widget.Widget{
		Name:     "Banner",
		Template: "banner",
		Setting:  Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				argument := setting.(*bannerArgument)
				context.Options["Title"] = argument.Title
				context.Options["SubTitle"] = argument.SubTitle
			}
			return context
		},
	})
}
