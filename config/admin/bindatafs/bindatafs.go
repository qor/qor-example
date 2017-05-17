package bindatafs

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jteeuwen/go-bindata"
	"github.com/qor/admin"
)

type AssetFSInterface interface {
	NameSpace(nameSpace string) AssetFSInterface
	RegisterPath(path string) error
	Asset(name string) ([]byte, error)
	Glob(pattern string) (matches []string, err error)
	FileServer(dir AssetFS) http.Handler
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
	AssetFS = &bindataFS{AssetFileSystem: &admin.AssetFileSystem{}, Path: "config/admin/bindatafs/"}
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
	copyFiles(filepath.Join(assetFS.Path, "templates"), assetFS.ViewPaths)
	for _, fs := range assetFS.nameSpacedFS {
		copyFiles(filepath.Join(assetFS.Path, "templates", fs.nameSpace), fs.ViewPaths)
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

	return bindata.Translate(config)
}

var cacheSince = time.Now().Format(http.TimeFormat)

type AssetFS struct {
	Prefix     string
	Dir        string
	AssetPaths []string
}

func (assetFS *bindataFS) FileServer(fs AssetFS) http.Handler {
	fileServer := assetFS.NameSpace("file_server")
	fileServer.RegisterPath(fs.Dir)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("If-Modified-Since") == cacheSince {
			w.WriteHeader(http.StatusNotModified)
			return
		}
		w.Header().Set("Last-Modified", cacheSince)

		requestPath := r.URL.Path
		if fs.Prefix != "" {
			requestPath = strings.TrimPrefix(requestPath, "/"+strings.TrimPrefix(fs.Prefix, "/"))
		}

		if content, err := fileServer.Asset(requestPath); err == nil {
			etag := fmt.Sprintf("%x", md5.Sum(content))
			if r.Header.Get("If-None-Match") == etag {
				w.WriteHeader(http.StatusNotModified)
				return
			}

			if ctype := mime.TypeByExtension(filepath.Ext(requestPath)); ctype != "" {
				w.Header().Set("Content-Type", ctype)
			}

			w.Header().Set("Cache-control", "private, must-revalidate, max-age=300")
			w.Header().Set("ETag", etag)
			w.Write(content)
			return
		}

		http.NotFound(w, r)
	})
}

func copyFiles(templatesPath string, viewPaths []string) {
	for i := len(viewPaths) - 1; i >= 0; i-- {
		viewPath := viewPaths[i]
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
