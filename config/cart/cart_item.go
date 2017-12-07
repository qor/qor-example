package cart

import (
	"reflect"
)

type CartItem struct {
	ProductID uint `form:"id" json:"id" binding:"required"`
	Quantity  uint `form:"qty" json:"qty" binding:"required"`
}

func (moduleItem *CartItem) Bind(ptr interface{}) error {
	var (
		typ = reflect.TypeOf(ptr).Elem()
		val = reflect.ValueOf(ptr).Elem()

		modItVal = reflect.ValueOf(moduleItem)
	)

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)

		if !structField.CanSet() {
			continue
		}

		// structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get("cartitem")

		if inputFieldName == "" {
			continue
		}

		val.Field(i).Set(reflect.Indirect(modItVal).FieldByName(inputFieldName))
	}
	return nil
}
