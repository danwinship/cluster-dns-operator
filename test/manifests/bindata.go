// Code generated by go-bindata.
// sources:
// test/assets/app-dns/deployment.yaml
// test/assets/app-dns/namespace.yaml
// test/assets/app-dns/service.yaml
// test/assets/cluster-dns-cr.yaml
// DO NOT EDIT!

package manifests

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

var _testAssetsAppDnsDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\xb1\x4e\x34\x31\x0c\x84\xfb\x3c\xc5\xbc\xc0\xea\x74\x6d\xea\xbf\xfc\x6b\x7a\x93\x1d\xd8\x08\x27\xb1\x62\x83\xc4\xdb\xa3\x70\xc7\x8a\x13\x53\x25\x1e\x7f\xa3\xf1\x5b\xed\x7b\xc6\x3f\x9a\x8e\xcf\xc6\x1e\x49\xac\x3e\x71\x7a\x1d\x3d\x43\xcc\xfc\xf2\x71\x4d\x8d\x21\xbb\x84\xe4\x04\x74\x69\xfc\x76\xee\x6f\x37\x29\xcc\x28\xfa\xee\xc1\xb9\xed\xdd\xb7\xa0\x47\x72\x63\x59\xfb\x93\xa6\xb5\x88\x67\x5c\x13\xe0\x54\x96\x18\x73\x39\x40\x93\x28\xc7\x7f\x79\xa6\xfa\x6d\x80\x15\x9c\x71\x86\x00\xc1\x66\x2a\xc1\x3b\xf0\xab\xc8\x92\x3e\xb0\x7f\x69\xe0\xa7\xc6\x52\x19\x3d\xa4\x76\xce\x93\xd8\x1e\xce\xb9\xa9\x36\x79\x65\xc6\x30\x76\x3f\xea\x4b\x5c\x0e\xaa\x8e\xed\xfc\xa7\xaf\x00\x00\x00\xff\xff\x67\xd5\x2b\xab\x33\x01\x00\x00")

func testAssetsAppDnsDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppDnsDeploymentYaml,
		"test/assets/app-dns/deployment.yaml",
	)
}

func testAssetsAppDnsDeploymentYaml() (*asset, error) {
	bytes, err := testAssetsAppDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-dns/deployment.yaml", size: 307, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x12\x84\x30\x08\x05\xd0\x9e\x53\x70\x81\x14\xdb\x72\x88\x2d\xed\xff\x24\xbf\xc8\x28\x98\x09\xe8\xf9\x7d\xe7\x8c\x61\xfa\x87\x33\x17\x3a\x05\x6b\x1e\xdc\x39\xef\x30\x7d\x7f\xe2\x2c\x0c\x14\x4c\x54\x03\x4e\xd3\x7e\x3d\x59\xdc\x6d\x44\xb6\x62\x96\x7c\x01\x00\x00\xff\xff\xa7\x68\xf5\x65\x42\x00\x00\x00")

func testAssetsAppDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppDnsNamespaceYaml,
		"test/assets/app-dns/namespace.yaml",
	)
}

func testAssetsAppDnsNamespaceYaml() (*asset, error) {
	bytes, err := testAssetsAppDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-dns/namespace.yaml", size: 66, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8c\x41\xaa\xc3\x30\x0c\x44\xf7\x3e\x85\x2e\x10\xf8\x7f\x57\xb4\xed\x05\x02\x2d\xdd\x0b\x47\xb4\xa6\x8e\x2d\xa4\x69\xce\x5f\x1c\x4c\xb5\x1a\xcd\x7b\xcc\xbb\xb4\x8d\xe9\xa6\x7e\x94\xac\x49\xac\x3c\xd4\xa3\xf4\xc6\x74\xfc\xa7\x5d\x21\x9b\x40\x38\x11\x35\xd9\x95\x49\xcc\x66\x0e\x93\xac\x4c\xb9\x7e\x02\xea\xcb\xd6\x62\x81\x06\x52\x98\xe6\xe1\x87\x56\xcd\xe8\x3e\xf2\x38\x31\x63\xfa\x59\x44\xd6\x1d\x31\xe0\x32\xa7\x5f\x80\x9d\xee\x20\x4c\x97\xbf\xf3\x81\xf8\x53\xb1\xce\x6a\x96\xe6\x1d\x3d\xf7\xca\x74\xbf\xae\xe9\x1b\x00\x00\xff\xff\x20\xab\xa9\x82\xc3\x00\x00\x00")

func testAssetsAppDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppDnsServiceYaml,
		"test/assets/app-dns/service.yaml",
	)
}

func testAssetsAppDnsServiceYaml() (*asset, error) {
	bytes, err := testAssetsAppDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-dns/service.yaml", size: 195, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsClusterDnsCrYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8b\x3d\xaa\xc3\x30\x0c\xc7\x77\x9f\xc2\x17\x78\x7e\x64\xf5\xda\xce\x5d\x0a\xdd\xff\xd8\x0a\x11\x95\x25\x13\x29\x3d\x7f\x29\x78\xfd\x7d\x60\xf2\x8b\x4e\x67\xd3\x9a\xbb\x7a\xb1\x49\xea\x07\xef\x51\xd8\xfe\x3f\x1b\x64\x1e\xd8\xd2\x9b\xb5\xd7\x7c\x93\xcb\x83\xce\xfb\xe3\x99\x06\x05\x3a\x02\x35\xe5\xac\x18\x54\x73\x90\xc7\x5f\xa7\x1d\x97\x44\xf2\x49\xed\xa7\xda\x3a\x6c\x80\x75\x35\x8b\x15\xb1\x06\x49\xdf\x00\x00\x00\xff\xff\x4f\x5f\xeb\x4b\x80\x00\x00\x00")

func testAssetsClusterDnsCrYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsClusterDnsCrYaml,
		"test/assets/cluster-dns-cr.yaml",
	)
}

func testAssetsClusterDnsCrYaml() (*asset, error) {
	bytes, err := testAssetsClusterDnsCrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/cluster-dns-cr.yaml", size: 128, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"test/assets/app-dns/deployment.yaml": testAssetsAppDnsDeploymentYaml,
	"test/assets/app-dns/namespace.yaml": testAssetsAppDnsNamespaceYaml,
	"test/assets/app-dns/service.yaml": testAssetsAppDnsServiceYaml,
	"test/assets/cluster-dns-cr.yaml": testAssetsClusterDnsCrYaml,
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
	"test": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"app-dns": &bintree{nil, map[string]*bintree{
				"deployment.yaml": &bintree{testAssetsAppDnsDeploymentYaml, map[string]*bintree{}},
				"namespace.yaml": &bintree{testAssetsAppDnsNamespaceYaml, map[string]*bintree{}},
				"service.yaml": &bintree{testAssetsAppDnsServiceYaml, map[string]*bintree{}},
			}},
			"cluster-dns-cr.yaml": &bintree{testAssetsClusterDnsCrYaml, map[string]*bintree{}},
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

