package cart

import (
	"github.com/qor/qor-example/app/models"
)

type CartItem struct {
	SizeVariationID uint `form:"sizevariation" json:"sizevariation"`
	Quantity        uint `form:"quantity" json:"quantity"`
}

type fullCartItem struct {
	CartItem
	SizeVariation models.SizeVariation
}
