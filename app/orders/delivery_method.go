package orders

import (
	"github.com/jinzhu/gorm"
)

type DeliveryMethod struct {
	gorm.Model

	Name  string
	Price float32
}
