// Code generated by go-bindata.
// sources:
// data/exercises.template.scm
// data/tests.template.scm
// DO NOT EDIT!

package pset

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

var _dataExercisesTemplateScm = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xc8\xc9\x4f\x4c\x51\x50\xaa\xae\xd6\x0b\x49\x2d\x2e\x29\x76\xcb\xcc\x49\xad\xad\x55\xd2\xe4\x42\x88\xbb\xa4\x16\x20\x84\xb9\xaa\xab\x8b\x12\xf3\xd2\x53\x15\xf4\x02\x8a\xf2\x93\x72\x52\x73\x8b\x6b\x6b\xb9\xac\xad\xad\x15\x5c\x2b\x52\x8b\x92\x33\x8b\x53\x15\x80\x3a\x9c\x33\x12\x0b\x4a\x52\x8b\x6a\x6b\xf5\x80\x1c\xbf\xd2\xdc\x24\x10\x9b\x8b\x4b\xa3\x04\x68\x83\x6e\x2a\x54\xa1\x2e\xb2\x42\x5d\x24\x85\x60\x4b\x52\xf3\x52\x6a\x6b\x01\x01\x00\x00\xff\xff\xad\x6f\xfe\x40\x9c\x00\x00\x00")

func dataExercisesTemplateScmBytes() ([]byte, error) {
	return bindataRead(
		_dataExercisesTemplateScm,
		"data/exercises.template.scm",
	)
}

func dataExercisesTemplateScm() (*asset, error) {
	bytes, err := dataExercisesTemplateScmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/exercises.template.scm", size: 156, mode: os.FileMode(420), modTime: time.Unix(1461601371, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataTestsTemplateScm = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xc8\xc9\x4f\x4c\x51\x50\x2a\x49\x2d\x2e\xc9\xcc\x4b\xd7\x2b\x4e\xce\x55\xd2\xe4\x02\x82\xea\xea\xa2\xc4\xbc\xf4\x54\x05\xbd\x80\xa2\xfc\xa4\x9c\xd4\xdc\xe2\xda\x5a\x2e\x8d\x94\xd4\xb4\xcc\xbc\x54\x05\x0d\x90\x6a\xdd\xd4\x8a\xd4\xa2\xe4\xcc\xe2\x54\xdd\xea\x6a\x3d\xe7\x8c\xc4\x82\x92\xd4\xa2\xda\x5a\x10\xc7\xaf\x34\x37\x09\xc4\xd6\xd4\x04\x9a\x92\x9a\x97\x02\xd4\x09\x08\x00\x00\xff\xff\x9c\x97\xbb\x16\x66\x00\x00\x00")

func dataTestsTemplateScmBytes() ([]byte, error) {
	return bindataRead(
		_dataTestsTemplateScm,
		"data/tests.template.scm",
	)
}

func dataTestsTemplateScm() (*asset, error) {
	bytes, err := dataTestsTemplateScmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/tests.template.scm", size: 102, mode: os.FileMode(420), modTime: time.Unix(1461597510, 0)}
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
	"data/exercises.template.scm": dataExercisesTemplateScm,
	"data/tests.template.scm": dataTestsTemplateScm,
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
	"data": &bintree{nil, map[string]*bintree{
		"exercises.template.scm": &bintree{dataExercisesTemplateScm, map[string]*bintree{}},
		"tests.template.scm": &bintree{dataTestsTemplateScm, map[string]*bintree{}},
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

