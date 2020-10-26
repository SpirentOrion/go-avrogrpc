// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// codegen.template
package main

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _codegenTemplate = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x6f\xdb\xba\x15\x7f\x96\x3e\xc5\x99\x57\x14\x76\xe7\xca\xc3\x5e\x06\x78\x08\x86\xde\xb4\xc5\xf2\xd0\xb4\x68\xda\xbd\x14\x45\x2e\x4d\x1d\xcb\x44\x24\x52\x25\x29\x27\x81\xa0\xef\x3e\x90\x94\x64\xd9\x96\x12\x29\x71\x72\x93\xdd\x3c\xb4\x91\x28\xf2\xfc\x3f\x3f\x1e\x1d\xca\xb3\x19\x1c\x8b\x10\x21\x42\x8e\x92\x68\x0c\x61\x71\x6d\x6e\xde\x46\x32\xa5\x6f\x53\x29\xb4\xa0\x22\x0e\xe0\xfd\x67\x38\xfd\xfc\x0d\x3e\xbc\x3f\xf9\x16\xf8\x7e\x9e\xbf\xc2\xab\xf4\x0c\xe5\x9a\x51\x84\xf9\x11\x50\x92\x60\x7c\x4c\x14\x42\xf0\xa5\x5a\x53\x5d\xc0\xdb\xa2\x30\x2b\x18\xd7\xbb\x2b\x7e\x23\xf4\xa2\x6b\x85\x9f\x12\x7a\x41\x22\x84\x3c\x0f\xbe\xb8\x4b\x33\xca\x92\x54\x48\x0d\x63\xdf\xa3\x82\x6b\xbc\xd2\x30\x2a\x2f\x46\xbe\x17\x09\xb2\x96\x02\x46\x11\xd3\xab\x6c\x11\x50\x91\xcc\x62\xc6\x2f\x30\x64\x7c\xe6\x9e\xcd\xd6\xff\x30\xf3\x64\x4a\x61\x14\x09\x11\xc5\x18\x44\x22\x26\x3c\x0a\x84\x8c\x66\x66\x7c\x64\x28\x87\xa8\xba\x9e\xcf\xec\xd3\x91\xef\x29\x4d\x74\xd6\x3d\xcd\x3d\x1e\xf9\x13\xdf\x9f\xcd\xe0\xdb\x8a\x29\x60\x0a\x08\x50\x91\xa4\x2c\xc6\xb7\x9a\x25\x08\x44\x29\x94\x9a\x09\x0e\x5a\x00\x72\x95\x49\x04\xbd\x22\x1a\xb4\x99\xbf\x71\xca\x92\xc5\x68\xc8\x30\x65\xd7\x13\xcd\x16\x31\xc2\x25\xd3\x2b\xd0\x2b\x04\xab\x4f\x65\x2f\xa6\xcd\xb4\x05\x32\x1e\x55\xcc\x42\x20\x11\x61\x5c\xe9\xc0\xa7\x82\x2b\x0d\xe7\x70\x64\x17\x05\x67\x59\x6a\xec\x59\x1a\xf8\x44\xfd\x17\xa5\x62\x82\xff\xd3\x0a\xbd\xe5\xe6\xa2\x38\x8e\x19\x72\x4b\xdc\xf0\xa4\xee\xee\xdd\x97\x13\x58\x0a\x69\x87\x8c\xa7\x76\x7d\x59\x14\xa0\x1c\x81\xc0\x9f\xcd\x0c\xd5\x8f\x42\x82\xc2\x84\x70\xcd\xa8\x02\x22\x45\xc6\x43\xa0\xfa\x0a\x32\x85\x40\xcc\x75\x2c\x14\xe3\xd1\x0c\x79\x68\x74\x50\x5a\x22\x49\xcc\xd5\xd7\x2f\xc7\x6a\x0a\x69\x8c\x26\xd0\x24\x2e\x51\x1a\xb3\xad\xb4\x4e\xd5\x7c\x36\x4b\x2f\xa2\x20\x12\x41\x88\xeb\x59\x87\x4b\xfe\xad\xc9\xe2\x28\x14\xf4\xaf\x4e\x93\x63\xc1\x79\x70\x8a\x97\x67\x96\x41\xe0\xeb\xeb\x14\xbb\x74\xe6\x1a\xe5\x92\x50\x84\xdc\xcf\x73\x49\x78\x84\xf0\x8a\x93\x04\xa7\xf0\x2a\x41\xa5\x8c\xdd\xe7\x47\x8d\x50\xfe\xe4\x06\x95\x0d\x65\xcf\xd9\xb2\x9c\x18\xbc\x17\xd4\x66\x04\x5b\x02\xfe\xaa\xd7\x07\x4e\x0c\x18\x49\xfc\x95\xa1\xd2\x23\xb7\x34\xcf\x37\xb9\x65\x39\x16\xc5\xd8\x18\xab\x0c\xfa\xe0\xd8\xfd\x9d\x82\x48\xb5\x82\x20\x08\xac\x5b\x8f\x49\x1c\x7f\x4e\x4d\x60\x4d\x60\xbc\xa3\xd2\x79\x0b\x49\xa7\xe6\x14\x50\x4a\x21\x27\x7e\x9e\x63\xac\x10\x3a\x05\x54\xa9\xe0\x0a\x87\x4a\xc8\x38\x24\x24\xfd\xa1\xb4\x64\x3c\xfa\x59\x9b\x34\x2f\x1e\x53\xf8\x05\x0b\x99\x7c\x02\xb6\x7d\x04\xd3\x6d\x4d\xdb\x70\xe7\x61\x89\xc8\xd5\x55\xe1\xd7\xa1\xbf\xc1\xe8\x3a\xf4\x95\x96\x19\xd5\x90\xfb\x1e\xa5\x0e\x33\x36\xd9\x73\x52\x71\x30\x34\x96\x19\xa7\x70\x8a\x97\xad\x19\x34\xbe\x61\xf1\xa4\x23\xe9\x72\xdf\x93\xa8\x33\xc9\xe1\x75\xab\x68\x39\xa5\x56\xf8\x3c\x7f\xe5\x30\xe2\x84\x87\x78\x65\xf2\xf0\xef\x56\xbf\x01\x69\xda\x27\x1f\x8b\xc2\xa9\x38\xa6\xf0\xa6\x55\x1e\xa3\xc7\xe3\x07\x94\x31\x93\x53\xdf\x8e\xd8\xbd\x35\xa0\x74\x03\x6c\x46\x82\x29\xbc\x3e\xdf\xa5\x5c\xc2\xf2\x7b\x54\xb4\xd4\x55\xfd\xd8\x36\x65\x51\xfc\x9c\xc2\x68\x96\xe7\xaf\xda\x70\xdd\x8c\x3b\x81\x46\x4e\xa1\x20\x08\x26\xbe\x67\xac\x28\x25\xfc\xe5\x08\x38\x8b\x8d\x70\x95\x13\x39\x8b\xad\x84\xbe\x57\xf8\x9e\x75\xd3\xae\x5b\x3b\x75\xcd\x9d\x50\x45\x1d\x10\x57\x53\x43\xae\x19\xb9\xfd\x8c\xb6\x05\xe4\xde\x19\xf2\xb0\x99\x25\x13\x67\x52\xdf\x3b\x8e\x85\xc2\x77\x3c\xfc\x8a\x74\x3d\xee\xc8\x24\xaf\x11\xce\xce\x7c\x5d\x89\xd4\x2d\xcd\x26\xb7\x5a\x89\xb9\x70\xbb\xda\x0b\xb7\x4e\x82\x13\xb0\x2a\x25\xb0\xaf\x94\x61\xb2\xc8\x96\x75\x8c\xec\x85\x43\x0b\xd1\xf3\x32\xf4\x4d\x85\x48\x83\xdf\x18\x27\xf2\xfa\xa3\x14\xc9\x29\xd1\x6c\x8d\x63\xeb\xd0\xe4\x26\x97\x57\xde\xae\xbc\xb6\xa5\x61\x60\x64\xfd\xa4\xa2\xf1\x22\x5b\x4e\xee\xa8\x6e\x0f\x47\x19\x79\x4a\x01\xe7\x47\xbb\x22\xd8\xf5\xd6\x66\x93\x7f\xf5\x09\xdb\x35\x91\xb0\xc8\x96\xf0\xe3\xe7\xe2\x5a\xe3\x0d\x84\x8d\x40\x46\xb7\xd7\x46\xb9\x5e\xa4\x93\x29\x9c\x0f\xf5\x8e\xdb\x87\x9d\x7b\x9c\x53\x8c\x7b\x9c\xa3\x9c\x59\x7b\xa6\x63\x39\x96\x54\x69\xb5\x0f\xa9\x8c\x53\xd8\xc6\x86\xbe\x55\xc2\x21\x71\xf3\x71\x4a\x88\xe7\x09\xa9\x0f\x91\xdd\x8c\x3f\x38\xa2\x77\xa6\x50\x13\x1e\x7a\x65\xd0\xc1\xb2\xfc\xa0\x9b\xcc\xf3\xdb\x43\x6e\x01\xd3\x17\x0c\xec\x8d\x81\xee\x65\xe3\xa5\x70\xfc\xd3\x14\x8e\xcf\x2f\xd9\xff\x64\x05\xe3\x0b\xb6\x0d\xc5\xb6\x3f\xbc\x7c\xeb\x70\xd5\x1f\x5c\xee\x94\xb3\x4a\xcc\x3c\xe1\x6b\x71\x81\x0e\x30\xfb\xe1\x9e\x15\xff\xb5\xfd\xbf\x82\xc0\x5e\x61\x24\x32\xfd\x04\x02\xc9\x4a\x51\x85\x12\xf2\xb0\xd8\xfc\x6d\x69\x1c\x9b\x0b\x94\x55\xe3\x58\xb9\xbb\x61\x8d\x63\x78\x17\xc7\xc0\x92\x34\xc6\x04\xb9\x26\x26\x34\x14\x24\x99\xd2\x80\xc9\x02\x43\xf8\xce\xeb\x87\x18\xb6\xb2\x37\x44\x0c\xbf\xa5\x90\x97\x44\x86\x75\x1f\x9d\xc5\x4c\x5f\xb7\xee\x06\x95\xd8\x4f\xaa\xf7\xab\xe4\xba\xcf\xae\xe5\x64\xaf\xf6\xa5\x7b\xf7\x73\x6f\x48\xdc\x07\x10\xe8\xb6\x1e\xed\x7d\x58\x1e\xa8\xef\x3a\xa0\xbd\xea\x99\x30\xfd\x60\xa2\xf4\xf6\x20\x1d\xdb\x7d\x6d\x36\xeb\x11\xcf\x2e\xfa\x17\xe8\x12\x20\xc4\xd0\x9e\x85\x90\x35\xee\x45\x78\x8c\xbb\x99\x53\x9e\x75\xf4\x60\x52\x17\x1c\xb6\xcf\xfa\x50\x1d\xd5\xdb\x05\x69\xdf\x64\x86\x06\x41\xa3\xa5\xec\x8e\xe7\x82\x0f\x66\x78\x39\xb6\x67\x7a\xc1\x96\x18\x53\x18\x25\xa8\x57\x22\x6c\xe3\x0c\x5c\x68\x68\x4c\x1e\x4d\xfc\xe1\xfd\x90\xbb\x6a\xdd\x95\x89\xcf\xc9\x1a\x3b\x6f\x46\xff\x4f\x01\x70\x6f\xa5\xf6\x50\x68\x08\x04\x35\x34\xb4\x1b\xf7\xc3\xa8\xb9\xbd\xe3\xf7\x57\x77\x10\x14\x42\x5e\x61\xa1\x22\x4b\xec\x00\x41\x72\xbd\x8b\x81\x22\xd5\xa6\x3e\x01\xb1\x6c\xdf\xeb\xcb\xaa\x83\xa9\xad\xfa\xe2\xbb\x42\xb3\xc2\x8e\x6f\x76\x7c\xa6\xac\x05\x24\x52\x91\x24\xc8\x43\x63\x2e\xa2\x80\x58\x5e\xce\x6c\xca\xf0\x6c\x17\xee\x92\xc5\xb1\xa1\x2d\x51\x65\xb1\x79\x85\x2c\xcf\xe3\x2d\x08\x3b\x7f\x6d\xa0\xb8\x5b\xc7\xad\x57\xcf\xc1\x9b\x89\x75\xce\x57\x8c\x98\xd2\x28\xdb\xe7\xa9\xf2\x5b\x00\x37\xec\xe6\x4a\x22\x5b\xf7\xf6\x3a\x70\x7d\x4f\x05\x15\xd9\xf2\xe9\xf8\xc6\x5e\x81\x25\x37\x79\xd0\x8d\xa4\x57\x35\xfc\x1f\xc2\xc3\xd8\x68\x2d\xd7\xb0\x5d\xc9\x38\xaa\xb5\x2d\x50\x3a\x3e\x2d\xd8\x21\xd7\xc1\x2e\xf6\x94\x86\x09\xda\x32\xba\x47\xbf\xc2\x2d\xaf\xfa\x15\x93\x21\x1d\x8a\xb6\x30\x31\xaf\xca\xef\x78\x68\x3b\x9f\x77\xec\x54\x34\x4d\x30\xa0\x53\xb1\x5b\x38\xb4\x13\x1b\xf0\xea\x5e\x45\xdc\x96\x4a\x07\xec\x58\x34\xdf\x8d\x0e\xd1\xb2\x68\xaa\x7a\x9f\x96\x45\xa5\xf7\x3d\x5a\x16\x5b\xa2\x3c\x4e\xcb\xa2\xf1\x8e\x7d\xe0\x8e\xc5\xa0\xe2\xea\xc1\x91\xa0\xdb\xec\x6a\x80\xb9\xeb\xb3\x9b\x58\x2c\x1e\xd7\xd8\x75\x1f\x83\x1b\x7e\x86\x7f\xd0\x51\x58\x4e\xee\x0c\x7b\x8c\x4f\xfb\xb4\x6a\x1f\x00\xfa\x5a\x21\xef\xc9\xa2\xda\x73\x45\xb3\x21\x25\xfe\xcb\xd6\x7c\x80\xc3\x83\x97\xe0\x7d\xd9\x8a\x9f\xca\x56\x7c\x98\xc4\x6e\xed\xba\x85\x48\xc1\xd0\x6e\xc9\x92\xa9\x5b\x4e\x31\xd5\x42\x3a\x3c\xf8\x6e\x14\x72\xe6\x3d\xd9\x3c\xbb\x8b\x97\x42\xa4\x03\x5c\xf2\xf8\x7b\x76\xe7\x0b\xbd\xd5\x9b\x93\x78\x0a\x23\x97\x7e\x19\x4f\x88\x54\x2b\x12\x33\x1e\xb5\xbe\xd7\x97\x72\xcc\x61\x04\x7f\x33\xdc\x1c\xbd\xf1\x64\xd2\xbf\x28\x60\xcb\x2d\x5f\x1c\x6d\xc4\xb5\xa7\x03\x55\x35\x34\x10\x8f\xed\x11\x8a\x3d\x87\x31\x0c\x76\x0d\xb1\xef\x08\x23\xae\x99\x59\xf3\x3c\x20\xd8\x88\x4c\xef\x87\xc2\x70\x67\xdc\xea\x0a\x27\x49\x9b\x2f\xac\x76\xbb\x67\x2e\xce\x43\x4b\x61\x4f\xc2\xf7\x53\x60\x29\x8c\x98\xee\x6e\xee\x01\x80\xf1\xc1\xd4\xf7\xbc\x8f\x59\x1c\x7f\xb2\x3d\x8b\x79\xbf\x43\x2a\xcb\x69\xe5\x52\xd7\x30\xb3\x39\xd9\x9a\xb1\x12\x7f\x41\xaf\xd6\xd4\xfd\x63\x43\xe2\xaf\xce\xa0\x2c\xa3\xe6\x25\x68\xda\x83\xa6\x1c\x6b\x64\x6d\x95\x6e\xe6\xdf\x52\x4c\xa1\x74\x76\x6b\x97\xcf\xe0\xe6\xd8\xf7\x6e\x6a\xf1\xd4\xbf\x2a\xd9\x0c\x55\xb1\xc8\x28\x9e\x92\xc4\xc8\xdb\x7e\xd8\x67\xa2\xcd\x2b\xb7\x89\x6f\xd7\x29\xce\x61\xfc\xa6\x3d\x34\x8c\x99\x27\x66\xb6\x8b\x65\x35\x87\x1f\x3f\x2d\x57\x77\xef\x98\x0e\x3f\xa8\xeb\x6c\x32\x95\x07\x51\x9e\x67\x1d\x59\x72\xad\x75\x69\x24\x8b\x57\x2b\x30\x37\x59\x37\x64\x63\xb4\xab\x8b\x69\xfb\xf9\x91\x79\xe0\x79\xe5\x97\x36\x73\xa8\xd4\x75\x03\xf7\x52\x97\xe3\xad\xea\xba\x61\xa7\x2e\xc0\xcd\x1a\x0f\x56\xfa\x86\xc6\x5e\x75\x24\xb9\x75\x18\xe8\x55\xb8\x56\x1b\x43\xcb\x0c\xb7\xcd\x76\xdb\xc9\xea\x2e\xc1\xe6\x67\x1d\x2d\x04\x6f\xf5\xcb\x27\xd4\x24\x24\x9a\xb8\xc8\xae\xee\x2a\xf8\xbc\xcb\x89\xf1\xe0\x42\x02\xde\xb8\x1f\xb7\x05\xf6\xae\x37\x81\x06\x96\xed\x50\xd8\x28\x39\x29\xeb\x64\xc6\x99\x1e\xd7\xe5\x93\x41\xb7\xea\x54\x75\xb0\x7a\x6c\x39\xbc\x54\xaa\x50\xba\x94\xf2\x14\x2f\xed\xf0\xf8\xf7\xc6\xd1\xfa\x57\x37\xff\x8c\xae\x30\x21\x45\xf1\xfb\x3e\x0a\xa7\x84\x33\x3a\x46\x29\x27\xd5\x27\x1c\xc3\x2d\xd5\x4f\x12\xb7\xa0\xb7\x28\xcd\xdf\xe2\xfc\x2f\x00\x00\xff\xff\x19\x11\x9e\x31\x92\x39\x00\x00"

func codegenTemplateBytes() ([]byte, error) {
	return bindataRead(
		_codegenTemplate,
		"codegen.template",
	)
}

func codegenTemplate() (*asset, error) {
	bytes, err := codegenTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codegen.template", size: 14738, mode: os.FileMode(420), modTime: time.Unix(1603809484, 0)}
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
	"codegen.template": codegenTemplate,
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
	"codegen.template": &bintree{codegenTemplate, map[string]*bintree{}},
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