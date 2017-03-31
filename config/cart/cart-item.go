package cart

type CartItem struct {
	SizeVariationID uint `form:"sizevariation" json:"sizevariation"`
	Quantity        uint `form:"qty" json:"qty"`
}
