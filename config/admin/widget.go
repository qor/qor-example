package admin

import (
	"github.com/qor/qor"
	"github.com/qor/qor-example/db"
	"github.com/qor/widget"
)

var Widget *widget.WidgetInstance

func init() {
	Widget = widget.New(&qor.Config{DB: db.DB})
	Widget.RegisterWidget(&widget.Widget{
		Name:     "Banner",
		Template: "banner",
		Context:  func(context *widget.Context, setting interface{}) *widget.Context { return context },
	})
}
