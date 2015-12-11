package admin

import (
	"github.com/qor/exchange"
	"github.com/qor/qor-example/app/models"
)

var ProductExchange *exchange.Resource

func init() {
	ProductExchange = exchange.NewResource(&models.Product{}, exchange.Config{PrimaryField: "Code"})
	ProductExchange.Meta(exchange.Meta{Name: "Code"})
	ProductExchange.Meta(exchange.Meta{Name: "Name"})
	ProductExchange.Meta(exchange.Meta{Name: "Price"})
}
