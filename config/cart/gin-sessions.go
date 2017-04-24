package cart

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
)

type GinGonicSession struct {
	Session sessions.Session
}

func (gcs GinGonicSession) Restore() (map[uint]*CartItem, error) {
	var list map[uint]*CartItem

	session := gcs.Session
	data := session.Get("__meta_gin_cart")

	if data == nil {
		list = make(map[uint]*CartItem)
		return list, nil
	} else {
		encoded := data.(string)
		if err := json.Unmarshal([]byte(encoded), &list); err != nil {
			return list, err
		}
		return list, nil
	}
}

func (gcs GinGonicSession) Save(data map[uint]*CartItem) error {
	encoded, err := json.Marshal(data)

	if err != nil {
		return err
	}

	session := gcs.Session
	session.Set("__meta_gin_cart", string(encoded))
	session.Save()

	return nil
}
