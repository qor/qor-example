package admin

import (
	"github.com/qor/exchange"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"

	"github.com/qor/qor-example/app/models"
)

var ProductExchange = exchange.NewResource(&models.Product{}, exchange.Config{PrimaryField: "Code"})

func init() {
	ProductExchange.Meta(&exchange.Meta{Name: "Code"})
	ProductExchange.Meta(&exchange.Meta{Name: "Name"})
	ProductExchange.Meta(&exchange.Meta{Name: "Price"})

	ProductExchange.AddValidator(&resource.Validator{
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if utils.ToInt(metaValues.Get("Price").Value) < 100 {
				return validations.NewError(record, "Price", "price can't less than 100")
			}
			return nil
		},
	})
}
