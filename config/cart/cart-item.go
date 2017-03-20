package cart

type CartItem struct {
	SizeVariationID uint    `form:"sizevariation" json:"sizevariation"`
	Quantity        uint    `form:"quantity" json:"quantity"`
	Price           float32 `form:"price" json:"price"`
}
