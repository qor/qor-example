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

type AssetFSInterface interface {
	NameSpace(nameSpace string) AssetFSInterface
	RegisterPath(path string) error
	Asset(name string) ([]byte, error)
	Glob(pattern string) (matches []string, err error)
	Compile() error
}

var AssetFS AssetFSInterface

type bindataFS struct {
	Path            string
	ViewPaths       []string
	AssetFileSystem admin.AssetFSInterface
	nameSpacedFS    []*nameSpacedBindataFS
}

type nameSpacedBindataFS struct {
	*bindataFS
	nameSpace       string
	ViewPaths       []string
	AssetFileSystem admin.AssetFSInterface
}

func init() {
	AssetFS = &bindataFS{AssetFileSystem: &admin.AssetFileSystem{}, Path: "config/admin/bindatafs"}
}

func (assetFS *bindataFS) NameSpace(nameSpace string) AssetFSInterface {
	nameSpacedFS := &nameSpacedBindataFS{bindataFS: assetFS, nameSpace: nameSpace, AssetFileSystem: &admin.AssetFileSystem{}}
	assetFS.nameSpacedFS = append(assetFS.nameSpacedFS, nameSpacedFS)
	return nameSpacedFS
}

func (assetFS *bindataFS) RegisterPath(path string) error {
	assetFS.ViewPaths = append(assetFS.ViewPaths, path)
	return assetFS.AssetFileSystem.RegisterPath(path)
}

func (assetFS *bindataFS) Asset(name string) ([]byte, error) {
	name = strings.TrimPrefix(name, "/")
	if len(_bindata) > 0 {
		return Asset(name)
	}
	return assetFS.AssetFileSystem.Asset(name)
}

func (assetFS *bindataFS) Glob(pattern string) (matches []string, err error) {
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

func (assetFS *bindataFS) Compile() error {
	fmt.Println("Compiling QOR templates...")
	os.RemoveAll(filepath.Join(assetFS.Path, "templates"))
	assetFS.copyFiles(filepath.Join(assetFS.Path, "templates"))
	for _, fs := range assetFS.nameSpacedFS {
		fs.copyFiles(filepath.Join(assetFS.Path, "templates", fs.nameSpace))
	}

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

func (assetFS *bindataFS) copyFiles(templatesPath string) {
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

func (assetFS *nameSpacedBindataFS) RegisterPath(path string) error {
	assetFS.ViewPaths = append(assetFS.ViewPaths, path)
	return assetFS.AssetFileSystem.RegisterPath(path)
}

func (assetFS *nameSpacedBindataFS) Asset(name string) ([]byte, error) {
	name = strings.TrimPrefix(name, "/")
	if len(_bindata) > 0 {
		return Asset(filepath.Join(assetFS.nameSpace, name))
	}
	return assetFS.AssetFileSystem.Asset(name)
}

func (assetFS *nameSpacedBindataFS) Glob(pattern string) (matches []string, err error) {
	if len(_bindata) > 0 {
		nameSpacedPattern := filepath.Join(assetFS.nameSpace, pattern)
		for key, _ := range _bindata {
			if ok, err := filepath.Match(nameSpacedPattern, key); ok && err == nil {
				matches = append(matches, key)
			}
		}
		return matches, nil
	}

	return assetFS.AssetFileSystem.Glob(pattern)
}
