package bindatafs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jteeuwen/go-bindata"
	"github.com/qor/admin"
)

type BindataFS struct {
	Path            string
	ViewPaths       []string
	AssetFileSystem admin.AssetFSInterface
}

var AssetFS *BindataFS

func init() {
	AssetFS = &BindataFS{AssetFileSystem: &admin.AssetFileSystem{}, Path: "config/admin/bindatafs"}
}

func (assetFS *BindataFS) RegisterPath(path string) error {
	assetFS.ViewPaths = append(assetFS.ViewPaths, path)
	return assetFS.AssetFileSystem.RegisterPath(path)
}

func (assetFS *BindataFS) Asset(name string) ([]byte, error) {
	name = strings.TrimPrefix(name, "/")
	if len(_bindata) > 0 {
		return Asset(name)
	}
	return assetFS.AssetFileSystem.Asset(name)
}

func (assetFS *BindataFS) Glob(pattern string) (matches []string, err error) {
	if len(_bindata) > 0 {
		for key, _ := range _bindata {
			if ok, err := filepath.Match(pattern, key); ok && err == nil {
				matches = append(matches, key)
			}
		}
		return matches, nil
	}

	return assetFS.AssetFileSystem.Glob(pattern)
}

func (assetFS *BindataFS) Compile() error {
	fmt.Println("Compiling QOR templates...")
	os.RemoveAll(filepath.Join(assetFS.Path, "templates"))
	assetFS.copyFiles(filepath.Join(assetFS.Path, "templates"))

	config := bindata.NewConfig()
	config.Input = []bindata.InputConfig{
		{
			Path:      filepath.Join(assetFS.Path, "templates"),
			Recursive: true,
		},
	}
	config.Package = "bindatafs"
	config.Tags = "bindatafs"
	config.Output = filepath.Join(assetFS.Path, "templates_bindatafs.go")
	config.Prefix = filepath.Join(assetFS.Path, "templates")
	config.NoMetadata = true

	defer os.Exit(0)
	return bindata.Translate(config)
}

func (assetFS *BindataFS) copyFiles(templatesPath string) {
	for i := len(assetFS.ViewPaths) - 1; i >= 0; i-- {
		viewPath := assetFS.ViewPaths[i]
		filepath.Walk(viewPath, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				var relativePath = strings.TrimPrefix(path, viewPath)

				if info.IsDir() {
					err = os.MkdirAll(filepath.Join(templatesPath, relativePath), os.ModePerm)
				} else if info.Mode().IsRegular() {
					if source, err := ioutil.ReadFile(path); err == nil {
						err = ioutil.WriteFile(filepath.Join(templatesPath, relativePath), source, os.ModePerm)
					}
				}
			}
			return err
		})
	}
}
