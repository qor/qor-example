package cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qor/session/manager"
)

type Cart struct {
	CartItems map[uint]*CartItem `json:"cart_items,omitempty"`
	w         http.ResponseWriter
	req       *http.Request
}

type mutator func(*CartItem, uint)

func (cart *Cart) Add(cartItem *CartItem) (*CartItem, error) {
	if cartItem.ProductID == 0 {
		return nil, nil
	}
	if item, ok := cart.CartItems[cartItem.ProductID]; ok {
		cartItem.Quantity = cartItem.Quantity + item.Quantity
	}
	cart.CartItems[cartItem.ProductID] = cartItem
	if encoded, err := json.Marshal(cart.CartItems); err != nil {
		panic(err)
	} else {
		manager.SessionManager.Add(cart.w, cart.req, "__meta_qor_cart", string(encoded))
	}

	return cart.CartItems[cartItem.ProductID], nil
}

func (cart *Cart) Edit(id uint, item *CartItem) error {
	if _, exists := cart.CartItems[id]; exists {
		cart.CartItems[id] = item

		if encoded, err := json.Marshal(cart.CartItems); err != nil {
			return err
		} else {
			return manager.SessionManager.Add(cart.w, cart.req, "__meta_qor_cart", string(encoded))
		}
	}
	return fmt.Errorf("Item is not exist")
}

func (cart *Cart) Remove(id uint) error {
	if _, exists := cart.CartItems[id]; exists {
		delete(cart.CartItems, id)

		if encoded, err := json.Marshal(cart.CartItems); err != nil {
			return err
		} else {
			return manager.SessionManager.Add(cart.w, cart.req, "__meta_qor_cart", string(encoded))
		}
	}
	return fmt.Errorf("Item is not exist")
}

func (cart *Cart) GetContent() map[uint]*CartItem {
	return cart.CartItems
}

func (cart *Cart) IsEmpty() bool {
	if len(cart.GetContent()) > 0 {
		return false
	} else {
		return true
	}
}

func (cart *Cart) Each(callback mutator) {
	for key, item := range cart.CartItems {
		callback(item, key)
	}
}

func (cart *Cart) GetItemsIDS() (itemIDs []uint) {
	itemIDs = make([]uint, 0, len(cart.GetContent()))
	cart.Each(func(item *CartItem, key uint) {
		itemIDs = append(itemIDs, key)
	})

	return
}

func (cart *Cart) EmptyCart() error {
	cart.CartItems = make(map[uint]*CartItem)

	if encoded, err := json.Marshal(cart.CartItems); err != nil {
		return err
	} else {
		return manager.SessionManager.Add(cart.w, cart.req, "__meta_qor_cart", string(encoded))
	}
}

func GetCart(w http.ResponseWriter, req *http.Request) (*Cart, error) {
	var list map[uint]*CartItem

	data := manager.SessionManager.Get(req, "__meta_qor_cart")

	if data == "" {
		list = make(map[uint]*CartItem)
	} else {
		encoded := data
		if err := json.Unmarshal([]byte(encoded), &list); err != nil {
			panic(err)
		}
	}

	bucket := &Cart{
		CartItems: list,
		w:         w,
		req:       req,
	}
	return bucket, nil
}
