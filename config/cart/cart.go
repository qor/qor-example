package cart

/* import (
	"fmt"
) */

type Cart struct {
	CartItems map[uint]*CartItem
	storage   CartBucket
}

/* func init() {
	fmt.Println("Initialize cart module")
	// ginsession = sessions.Sessions("mysession", store)
	fmt.Printf("sss %v\n", "asdasd")

	// var st CartBucket
} */

func (module *Cart) Add(id, quantity uint) *CartItem {
	if item, ok := module.CartItems[id]; ok {
		quantity = quantity + item.Quantity
	}

	module.CartItems[id] = &CartItem{
		SizeVariationID: id,
		Quantity:        quantity,
	}

	module.storage.Save(module.CartItems)

	return module.CartItems[id]
}

func (module *Cart) GetContent() map[uint]*CartItem {
	return module.CartItems
}

func GetCart(storage CartBucket) (*Cart, error) {
	restored, _ := storage.Restore()
	bucket := &Cart{
		CartItems: restored,
		storage:   storage,
	}
	return bucket, nil
}
