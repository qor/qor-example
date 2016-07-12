package admin

import (
	"html/template"

	"github.com/qor/admin"
)

func initFuncMap() {
	Admin.RegisterFuncMap("render_latest_order", renderLatestOrder)
}

func renderLatestOrder(context *admin.Context) template.HTML {
	var orderContext = context.NewResourceContext("Order")
	orderContext.Searcher.Pagination.PerPage = 5
	// orderContext.SetDB(orderContext.GetDB().Where("state in (?)", []string{"paid"}))

	if orders, err := orderContext.FindMany(); err == nil {
		return orderContext.Render("index/table", orders)
	}
	return template.HTML("")
}
