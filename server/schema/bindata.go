// Code generated by go-bindata.
// sources:
// schema/query.graphql
// schema/schema.graphql
// schema/type/article.graphql
// schema/type/hello.graphql
// DO NOT EDIT!

package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcb\x31\x0a\x02\x31\x10\x85\xe1\x3e\xa7\x78\x8b\x8d\x36\x1e\x20\x9d\x60\x63\x29\x7a\x81\x45\x9f\x66\x60\x98\x84\xc9\x58\x04\xf1\xee\xc2\xa6\xfd\xf8\xff\x1d\xee\x85\xb8\x7e\xe8\x03\x31\x1a\xe1\x6c\xce\x4e\x8b\x8e\x55\x15\xf5\x85\x28\x04\x2d\x7c\xa0\x55\xb1\xe8\xc7\xb4\x85\xf3\xf9\x26\x00\x28\x54\xad\x19\xb7\x70\xb1\xf7\xb2\xd1\xea\x21\x0f\xe5\x5e\x9e\x19\x97\xf3\x72\xc8\x38\x4d\x49\xbf\x7f\x00\x00\x00\xff\xff\xb6\x51\x25\xa9\x74\x00\x00\x00")

func schemaQueryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaQueryGraphql,
		"schema/query.graphql",
	)
}

func schemaQueryGraphql() (*asset, error) {
	bytes, err := schemaQueryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/query.graphql", size: 116, mode: os.FileMode(436), modTime: time.Unix(1519241699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x80\x00\x00\x00\xff\xff\x54\xe0\x78\x3a\x1b\x00\x00\x00")

func schemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaGraphql,
		"schema/schema.graphql",
	)
}

func schemaSchemaGraphql() (*asset, error) {
	bytes, err := schemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.graphql", size: 27, mode: os.FileMode(436), modTime: time.Unix(1519165664, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypeArticleGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2c\x2a\xc9\x4c\xce\x49\x55\xa8\xe6\xe2\xe4\x54\x56\x08\xc9\x48\x55\xf0\x74\x51\xc8\x4f\x53\x28\xc9\x48\x55\x48\x84\xc8\x71\x71\x72\x66\xa6\x58\x29\x78\xba\x28\xc2\xd5\x94\x64\x96\xe4\xa4\x62\x2a\x03\x0b\x5b\x29\x04\x97\x14\x65\xe6\xa5\x2b\x72\xd5\x02\x02\x00\x00\xff\xff\x24\xb8\x65\xb3\x62\x00\x00\x00")

func schemaTypeArticleGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypeArticleGraphql,
		"schema/type/article.graphql",
	)
}

func schemaTypeArticleGraphql() (*asset, error) {
	bytes, err := schemaTypeArticleGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/type/article.graphql", size: 98, mode: os.FileMode(436), modTime: time.Unix(1519241953, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypeHelloGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x48\xcd\xc9\xc9\x57\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\x93\xf3\xf3\x4a\x52\xf3\x4a\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x15\xb9\x6a\x01\x01\x00\x00\xff\xff\xfe\x2a\x04\x6c\x2f\x00\x00\x00")

func schemaTypeHelloGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypeHelloGraphql,
		"schema/type/hello.graphql",
	)
}

func schemaTypeHelloGraphql() (*asset, error) {
	bytes, err := schemaTypeHelloGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/type/hello.graphql", size: 47, mode: os.FileMode(436), modTime: time.Unix(1519241009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"schema/query.graphql": schemaQueryGraphql,
	"schema/schema.graphql": schemaSchemaGraphql,
	"schema/type/article.graphql": schemaTypeArticleGraphql,
	"schema/type/hello.graphql": schemaTypeHelloGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"schema": &bintree{nil, map[string]*bintree{
		"query.graphql": &bintree{schemaQueryGraphql, map[string]*bintree{}},
		"schema.graphql": &bintree{schemaSchemaGraphql, map[string]*bintree{}},
		"type": &bintree{nil, map[string]*bintree{
			"article.graphql": &bintree{schemaTypeArticleGraphql, map[string]*bintree{}},
			"hello.graphql": &bintree{schemaTypeHelloGraphql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

