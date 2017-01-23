// Code generated by go-bindata.
// sources:
// data/config_schema_v3.0.json
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

var _dataConfig_schema_v30Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x4f\x8f\xdb\xb8\x0e\xbf\xe7\x53\x18\x6e\x6f\xcd\xcc\x14\x78\xc5\x03\x5e\x6f\xef\xb8\xa7\xdd\xf3\x0e\x5c\x43\xb1\x99\x44\x1d\x59\x52\x29\x39\x9d\xb4\xc8\x77\x5f\xc8\xff\x22\x2b\x92\xe5\x24\xee\xb6\x87\x9e\x66\x62\x91\x14\xff\xe9\x47\x8a\xf6\xf7\x55\x92\xa4\x6f\x55\xb1\x87\x8a\xa4\x1f\x93\x74\xaf\xb5\xfc\xf8\xf4\xf4\x59\x09\xfe\xd0\x3e\x7d\x14\xb8\x7b\x2a\x91\x6c\xf5\xc3\xfb\x0f\x4f\xed\xb3\x37\xe9\xda\xf0\xd1\xd2\xb0\x14\x82\x6f\xe9\x2e\x6f\x57\xf2\xc3\x7f\x1e\xdf\x3f\x1a\xf6\x96\x44\x1f\x25\x18\x22\xb1\xf9\x0c\x85\x6e\x9f\x21\x7c\xa9\x29\x82\x61\x7e\x4e\x0f\x80\x8a\x0a\x9e\x66\xeb\x95\x59\x93\x28\x24\xa0\xa6\xa0\xd2\x8f\x89\x51\x2e\x49\x06\x92\xfe\x81\x25\x56\x69\xa4\x7c\x97\x36\x8f\x4f\x8d\x84\x24\x49\x15\xe0\x81\x16\x96\x84\x41\xd5\x37\x4f\x67\xf9\x4f\x03\xd9\xda\x95\x6a\x29\xdb\x3c\x97\x44\x6b\x40\xfe\xd7\xa5\x6e\xcd\xf2\xa7\x67\xf2\xf0\xed\xff\x0f\x7f\xbf\x7f\xf8\xdf\x63\xfe\x90\xbd\x7b\x3b\x5a\x36\xfe\x45\xd8\xb6\xdb\x97\xb0\xa5\x9c\x6a\x2a\xf8\xb0\x7f\x3a\x50\x9e\xba\xff\x4e\xc3\xc6\xa4\x2c\x1b\x62\xc2\x46\x7b\x6f\x09\x53\x30\xb6\x99\x83\xfe\x2a\xf0\x25\x66\xf3\x40\xf6\x93\x6c\xee\xf6\xf7\xd8\x3c\x36\xe7\x20\x58\x5d\x45\x23\xd8\x53\xfd\x24\x63\xda\xed\xef\x8b\xdf\xaa\x37\x7a\x92\xb6\xa5\xb0\xf6\x6e\x14\x1c\x65\xbb\xcf\x55\xbe\x6c\x0b\xfb\x6a\x70\x56\xc0\x4b\x25\x48\x26\x8e\xe6\x59\xc0\x1f\x2d\x41\x05\x5c\xa7\x83\x0b\x92\x24\xdd\xd4\x94\x95\xae\x47\x05\x87\x3f\x8d\x88\x67\xeb\x61\x92\x7c\x77\x0f\xb6\x25\xa7\x59\x1f\xfd\x0a\x07\x7c\x58\x0f\xd8\x32\xac\x17\x82\x6b\x78\xd5\x8d\x51\xd3\x5b\xb7\x2e\x10\xc5\x0b\xe0\x96\x32\x98\xcb\x41\x70\xa7\x26\x5c\xc6\xa8\xd2\xb9\xc0\xbc\xa4\x85\x4e\x4f\x0e\xfb\x85\xbc\x78\x3e\x0d\xac\xd6\xaf\x6c\xe5\x11\x98\x16\x44\xe6\xa4\x2c\x47\x76\x10\x44\x72\x4c\xd7\x49\x4a\x35\x54\xca\x6f\x62\x92\xd6\x9c\x7e\xa9\xe1\x8f\x8e\x44\x63\x0d\xae\xdc\x12\x85\x5c\x5e\xf0\x0e\x45\x2d\x73\x49\xd0\x24\xd8\xb4\xfb\xd3\x42\x54\x15\xe1\x4b\x65\xdd\x35\x76\xcc\xf0\xbc\xe0\x9a\x50\x0e\x98\x73\x52\xc5\x12\xc9\x9c\x3a\xe0\xa5\xca\xdb\xfa\x37\x99\x46\xdb\xbc\xe5\x57\x8e\x80\xa1\x18\x2e\x1a\x8f\x92\x4f\x25\x76\x2b\xc6\xa4\xb6\xd1\x2d\x75\x18\x73\x05\x04\x8b\xfd\x8d\xfc\xa2\x22\x94\xcf\xf1\x1d\x70\x8d\x47\x29\x68\x9b\x2f\xbf\x5c\x22\x00\x3f\xe4\x03\x96\x5c\xed\x06\xe0\x07\x8a\x82\x57\xfd\x69\x98\x03\x30\x03\xc8\x1b\xfe\x57\x29\x14\xb8\x8e\x71\x0c\xb4\x97\x06\x53\x47\x3e\xe9\x39\x9e\x7b\xc3\xd7\x49\xca\xeb\x6a\x03\x68\x5a\xba\x11\xe5\x56\x60\x45\x8c\xb2\xfd\xde\xd6\xf2\xc8\xd3\x9e\xcc\xb3\x1d\x68\xdb\x60\xca\x3a\x61\x39\xa3\xfc\x65\xf9\x14\x87\x57\x8d\x24\xdf\x0b\xa5\xe7\x63\xb8\xc5\xbe\x07\xc2\xf4\xbe\xd8\x43\xf1\x32\xc1\x6e\x53\x8d\xb8\x85\xd2\x73\x92\x9c\x56\x64\x17\x27\x92\x45\x8c\x84\x91\x0d\xb0\x9b\xec\x5c\xd4\xf9\x96\x58\xb1\xdb\x19\xd2\x50\xc6\x5d\x74\x2e\xdd\x72\xac\xe6\x97\x48\x0f\x80\x73\x0b\xb8\x90\xe7\x86\xcb\x5d\x8c\x37\x20\x49\xbc\xfb\x1c\x91\x7e\x7a\x6c\x9b\xcf\x89\x53\xd5\xfc\xc7\x58\x9a\xb9\xed\x42\xe2\xd4\x7d\xdf\x13\xc7\xc2\x79\x0d\xc5\x28\x2a\x15\x29\x4c\xdf\x80\xa0\x02\x71\x3d\x93\x76\xcd\x7e\x5e\x89\x32\x94\xa0\x17\xc4\xae\x6f\x82\x48\x7d\x75\x21\x4c\x6e\xea\x1f\x67\x85\x2e\x7a\x81\x88\x58\x13\x52\x6f\xae\x9a\x67\x75\xe3\x29\xd6\xd0\x11\x46\x89\x82\xf8\x61\x0f\x3a\x72\x24\x8d\xca\xc3\x87\x99\x39\xe1\xe3\xfd\xef\x24\x6f\x80\x35\x28\x73\x7e\x8f\x1c\x11\x75\x56\xa5\x39\x6e\x3e\x45\xb2\xc8\x69\xfb\xc1\x2d\xbc\xa4\x65\x18\x2b\x1a\x84\xb0\x0f\x98\x14\xa8\x2f\x4e\xd7\xbf\x53\xee\xdb\xad\xef\xae\xf6\x12\xe9\x81\x32\xd8\xc1\xf8\xd6\xb2\x11\x82\x01\xe1\x23\xe8\x41\x20\x65\x2e\x38\x3b\xce\xa0\x54\x9a\x60\xf4\x42\xa1\xa0\xa8\x91\xea\x63\x2e\xa4\x5e\xbc\xcf\x50\xfb\x2a\x57\xf4\x1b\x8c\xa3\x79\xc6\xfb\x4e\x50\x36\xe2\x39\xaa\x42\xdf\x56\xaf\x95\x2e\x29\xcf\x85\x04\x1e\xf5\x8e\xd2\x42\xe6\x3b\x24\x05\xe4\x12\x90\x8a\xd2\x67\xe0\xda\x8e\x75\x59\x23\x31\xfb\x5f\x8a\x51\x74\xc7\x09\x8b\x39\x5a\x57\x72\x7b\xe3\xc5\x42\xeb\x78\xb8\x6b\x46\x2b\x1a\x3e\x07\x1e\x80\x9d\x51\x03\x5a\xfc\xf7\xc3\xfe\x04\xe4\x9f\x35\xa5\x5c\xc3\x0e\xd0\x87\x94\x13\x5d\xc7\x74\xd3\x31\xa3\xdb\xd8\x13\x1c\x07\x74\x42\x8f\x86\x41\x89\xad\xf6\x33\xf8\x7a\x11\xaf\x5e\xa3\xe1\x6f\x23\x6f\xdd\x29\x92\x79\xe9\xaf\x82\x73\x57\x8d\x2c\x88\xa8\x27\x2f\xa2\xd6\x2a\xda\x18\x36\x34\x5c\x4d\x35\x35\x03\xa9\x35\xc5\x5c\x14\x2f\x4c\xa3\x64\x0e\x41\x49\xfd\xda\xae\x1c\xcb\xae\x98\x23\x3b\x77\x96\x5e\x80\x6f\xa2\x68\x93\x46\x27\xb0\xd3\xd3\xcd\x8e\x28\x38\x79\xa4\x8a\x6c\x9c\x99\x9b\xef\x70\x9b\x6c\xc4\x43\x1c\x63\x10\x34\x52\x27\x2e\x1d\xda\x8e\xf0\x04\xd4\xaf\x39\x38\xd0\xb4\x02\x51\xfb\x6b\xd6\xca\xce\xef\x8e\x29\xb5\x26\xb3\x91\xa0\x5a\x94\x6e\x4c\x9f\x87\xa0\xf6\xfd\x45\x34\x70\x73\x0e\x09\x82\x64\xb4\x20\x2a\x06\x44\x77\x5c\x50\x6b\x59\x12\x0d\x79\xfb\xa2\xea\x2a\xe8\x9f\xc0\x7c\x49\x90\x30\x06\x8c\xaa\x6a\x0e\x86\xa6\x25\x30\x72\xbc\xa9\x7c\x36\xec\x5b\x42\x59\x8d\x90\x93\x42\x77\xef\xc2\x22\x39\x97\x56\x82\x53\x2d\xbc\x08\x31\x6f\xcb\x8a\xbc\xe6\xfd\xb6\x0d\x89\xf7\xc0\x04\xdb\xba\xb9\x77\x4b\x2b\x13\x94\xa8\xb1\xb8\x70\xf6\xcd\x21\x3a\xd7\xfa\x40\xc6\xf4\x3b\x5e\x98\x8e\xa0\x0c\x92\x0c\x57\xff\x28\x7f\xb4\xb4\x74\x7d\x66\x2e\x05\xa3\xc5\x71\x29\x0b\x0b\xc1\x5b\x27\xcf\x49\x88\x3b\x33\xd0\xa4\x83\x69\x85\x2a\xa9\xa3\x87\xb5\x61\xf8\x4a\x79\x29\xbe\x5e\xb1\xe1\x72\xa9\x24\x19\x29\xc0\xc1\xbb\x7b\x1d\xad\x34\x12\xca\xf5\xd5\xe5\xfc\x5e\xb3\xee\xa8\xe6\x43\x7e\x46\x50\x7f\xa0\x8b\xbf\x49\x0d\x20\x7d\x21\xeb\xe8\x3c\xa8\x82\x4a\xa0\x37\x01\x17\x78\xf3\x1d\x33\xb1\x27\x5b\xa0\xaa\xcd\x1a\x20\x76\x54\xe6\xbe\xb8\xf8\x6d\x23\x3e\x24\xcc\xe2\x80\x44\x25\xa9\x96\x3a\x1d\xb3\x47\xaa\xa9\xb7\x06\x27\xd3\xa3\x88\x24\x3c\x8e\x88\x69\x1d\xd7\xbd\xa3\x50\xf5\x86\xc3\x64\x47\x65\xf9\xd3\xf7\x9e\x77\xfe\x35\xe5\x14\xbe\x94\xdc\x07\x7a\xfd\xdb\x90\x40\x54\x9f\x87\x9e\x79\x3d\xf8\x2a\x9b\x1d\xe2\xe0\xab\x88\xe5\xf4\xbf\xb2\xc1\xbb\x03\x33\xba\x2f\x37\x22\x90\xd1\x51\xfd\x46\x8c\xdf\xf9\x75\x65\x7e\x39\x43\x2a\x2b\xcf\x2e\xef\x8f\x53\x29\x31\x7b\x3a\xdf\x71\x64\x63\x35\x5c\x32\xcf\x07\x74\x63\xb4\x9d\x1a\x4a\xf4\x24\x81\x69\xad\xb3\x69\xe7\xc4\x69\xcb\x17\xcc\xf0\xc7\x77\x13\x35\x65\xea\x2d\xda\x0f\x02\xe3\x05\x06\x3e\xfe\x98\x3a\x8d\x68\xef\xdd\xcb\xaf\xc0\x02\xa0\x66\xf1\x5f\x7c\x13\x66\xec\xe4\xc7\x8b\xf9\xc6\xf7\xf1\xd0\xae\xfd\x9e\x2b\x1b\xf9\xc7\x21\x69\xdf\x49\x5b\x90\x92\xd9\xbd\x79\x28\x8c\xde\x2f\xc5\xdc\x91\x61\xff\xc5\x56\xe6\x87\xab\x95\xfd\xb7\xf9\xba\x6e\x75\x5a\xfd\x13\x00\x00\xff\xff\x46\xf7\x7b\x23\xe5\x2a\x00\x00")

func dataConfig_schema_v30JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v30Json,
		"data/config_schema_v3.0.json",
	)
}

func dataConfig_schema_v30Json() (*asset, error) {
	bytes, err := dataConfig_schema_v30JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.0.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/config_schema_v3.0.json": dataConfig_schema_v30Json,
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
		"config_schema_v3.0.json": &bintree{dataConfig_schema_v30Json, map[string]*bintree{}},
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

