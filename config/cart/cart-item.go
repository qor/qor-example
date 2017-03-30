package cart

type CartItem struct {
	SizeVariationID uint `form:"sizevariation" json:"sizevariation"`
	Quantity        uint `form:"qty" json:"qty"`
}

type FullCartItem struct {
	CartItem
	MainImageURL string
	ProductName  string
	ColorName    string
	SizeName     string
	Price        float32
	Amount       float32
}
