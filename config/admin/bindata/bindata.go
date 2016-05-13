package bindata

import (
	"github.com/qor/admin"
	"github.com/qor/admin/bindata"
)

type Bindata struct {
	bindata.Bindata
}

type Config struct {
}

func New(config *Config) *Bindata {
	// config/admin/bindata
	return &Bindata{Bindata: {AssetFileSystem: &admin.AssetFileSystem{}, Config: config}}
}

func (bindata *Bindata) Asset(name string) ([]byte, error) {
	return bindata.AssetFileSystem.Asset(name)
}

func (bindata *Bindata) Glob(pattern string) (matches []string, err error) {
	return bindata.AssetFileSystem.Glob(name)
}

func (bindata *Bindata) Compile() error {
	bindata
	return bindata.AssetFileSystem.Compile()
}
