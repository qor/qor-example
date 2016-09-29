package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/sorting"
)

type Store struct {
	gorm.Model
	StoreName string
	Owner     Owner
	Phone     string
	Email     string
	location.Location
	sorting.Sorting
}

type Owner struct {
	Name    string
	Contact string
	Email   string
}

func (owner *Owner) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, owner)
	case string:
		if v != "" {
			return owner.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (owner Owner) Value() (driver.Value, error) {
	return json.Marshal(owner)
}
