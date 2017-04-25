package admin

import (
	"html/template"

	"github.com/qor/admin"
)

func initFuncMap() {
	Admin.RegisterFuncMap("render_latest_order", renderLatestOrder)
	Admin.RegisterFuncMap("render_latest_products", renderLatestProduct)
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

func renderLatestProduct(context *admin.Context) template.HTML {
	var productContext = context.NewResourceContext("Product")
	productContext.Searcher.Pagination.PerPage = 5
	// productContext.SetDB(productContext.GetDB().Where("state in (?)", []string{"paid"}))

	if products, err := productContext.FindMany(); err == nil {
		return productContext.Render("index/table", products)
	}
	return template.HTML("")
}
