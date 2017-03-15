package cart

type CartItem struct {
	SizeVariationID uint    `json:"sizevariation"`
	Quantity        uint    `json:"quantity"`
	Price           float32 `json: "price"`
}
