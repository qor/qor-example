// +build !bindatafs

package bindatafs

import "fmt"

var _bindata = map[string]interface{}{}

func Asset(name string) ([]byte, error) {
	return nil, fmt.Errorf("Asset %s not found", name)
}
