package cart

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Cart struct {
	CartItems map[uint]*CartItem
	storage   CartBucket
}

type mutator func(*CartItem, uint)

func (module *Cart) Add(cartItem *CartItem) *CartItem {
	fmt.Println(cartItem)

	if item, ok := module.CartItems[cartItem.SizeVariationID]; ok {
		cartItem.Quantity = cartItem.Quantity + item.Quantity
	}

	module.CartItems[cartItem.SizeVariationID] = cartItem
	module.storage.Save(module.CartItems)

	return module.CartItems[cartItem.SizeVariationID]
}

func (module *Cart) Remove(id uint) bool {
	if _, exists := module.CartItems[id]; exists {
		delete(module.CartItems, id)
		module.storage.Save(module.CartItems)
		return true
	}
	return false
}

func (module *Cart) GetContent() map[uint]*CartItem {
	return module.CartItems
}

func (module *Cart) IsEmpty() bool {
	if len(module.CartItems) > 0 {
		return false
	} else {
		return true
	}
}

func (module *Cart) Each(callback mutator) {
	for key, item := range module.CartItems {
		callback(item, key)
	}
	module.storage.Save(module.CartItems)
}

func GetCart(ctx *gin.Context) (*Cart, error) {
	storage := GinGonicSession{sessions.Default(ctx)}
	restored, _ := storage.Restore()
	bucket := &Cart{
		CartItems: restored,
		storage:   storage,
	}
	return bucket, nil
}
