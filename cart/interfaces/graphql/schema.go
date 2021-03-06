// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package graphql generated by go-bindata.// sources:
// schema.graphql
package graphql

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4d\x6f\xdc\x38\xd2\xbe\xfb\x57\xd0\x9e\x4b\x4f\x90\x77\x06\xef\x1e\x7d\x6b\x77\x27\x41\x63\xc6\x4e\x62\x3b\xd9\x43\x10\x18\xb4\x54\xdd\xcd\x0d\x45\xca\x24\x65\x5b\x58\xcc\x7f\x5f\xf0\x4b\x22\x29\xea\xc3\x09\x76\x80\xd9\xdd\x1c\x62\x4b\x2c\x16\x8b\xac\x87\xc5\xa7\x8a\xb2\x6a\x6b\x40\x1b\x5e\x55\x20\x0a\xb8\xdb\x42\xc1\x05\x56\x50\x6e\xb0\x50\xe8\x9f\x27\x08\x21\x54\x60\xa1\xce\x7b\x11\xdd\x72\x6a\x1a\x4a\x2f\xbc\x05\x4a\x1e\x41\x10\x90\xe7\xe8\x4b\x24\xb8\x4d\x44\xda\xd3\xaf\xa6\xeb\x01\x86\x4d\x17\xed\x86\x97\xb0\x2a\xdd\xa3\x7e\x38\x47\x37\x4a\x10\x76\x38\xfd\x39\x31\x60\xd0\xd9\x6b\x5d\x53\xfa\x01\xb7\x15\x30\x75\x0d\x0f\x0d\x11\x50\xee\x14\x54\x32\xe9\x7e\xf7\x41\x90\xc2\x35\x9d\x76\x93\xbc\x69\xaa\x0a\x8b\x36\x95\x75\xaf\x4f\x4f\xfe\x38\x39\x89\x57\x2b\x6c\x76\x8b\x55\x12\x59\xf0\x86\xa9\x74\xc4\x75\x5d\x53\x02\xe5\xd6\x37\xdb\x51\x65\x53\xa5\x0d\x41\x3f\x63\x64\x22\xf7\x8e\xec\xd5\x06\x8b\x72\x54\xee\x9d\xc0\xac\xbc\xe5\x0a\xd3\xbf\x13\x75\x9c\x15\x37\x92\x7e\xf0\xa8\xc7\xba\xd2\xaf\xb2\xfd\x8e\x58\x0e\xcd\xbe\xe0\x9c\x02\x66\xdd\xc4\x6e\xf1\x33\x0c\xd6\xdd\xbc\xf4\x12\xce\x51\x37\x40\xa1\x50\x84\x33\x2d\x71\x53\x53\xa2\x3e\x63\xda\x80\x1d\xff\xa2\xbd\x04\x75\xe4\xa5\x5c\x55\xf6\xe7\x39\xfa\xe2\x30\xf1\xf5\xe7\x81\x71\x59\x0f\x39\xcf\x90\xf2\x1c\xed\xb6\xd6\x3c\x60\x8a\xa8\x76\xb7\xed\xf0\x65\xde\xde\x13\x4a\x09\x3b\xac\xcb\x52\x80\x1c\x38\xd0\xbe\x35\x82\x75\x23\x8a\x23\x96\x20\x12\x99\x0f\x20\x24\x67\x6e\x6f\x8c\x6f\x89\x68\x27\xe0\xb2\x24\x7a\xf2\x98\x6e\xb1\xc2\xc3\x41\x83\x46\x6b\x65\x9d\xac\xda\x00\xda\x49\xbb\x9d\x1a\x50\xce\x0e\xf2\x96\xaf\x1b\x75\xd4\xb3\x2f\xf4\xe6\xf9\x64\xa6\x10\x39\x0e\xa7\xed\xe9\x22\x61\xeb\xf8\x0d\x6f\x6a\xce\xf4\x1e\x1d\x4c\xb0\x6f\x72\x53\x2c\x61\x8f\x1b\xaa\x36\x8d\x10\xc0\x8a\x36\xd6\xa7\x34\x00\x89\xdd\xa3\xb1\x9e\x5b\xdf\xe2\xd4\xe8\x5f\x37\x16\x93\x3b\xe6\x42\x50\x2d\x78\xd9\x14\x2a\x7d\x4d\x64\xb4\x0a\x50\x26\xb3\x3c\x74\x9b\x24\x85\xd0\x69\xb4\x31\x6e\xf1\x73\x7e\x1b\x78\xb1\x7b\x23\x76\x05\x23\x02\x78\xb0\x69\xbf\xe4\xa2\x82\x6f\x0f\x83\xe3\xf7\xc4\xc4\x38\x14\x6e\x83\x4e\xe1\xb6\x39\xf1\x02\x97\x98\xb0\x9b\x23\xa9\x6b\xc2\x0e\x6f\x2e\x31\xa1\xb1\x67\x88\x7c\x53\xd5\xaa\x4d\x96\xee\x88\xa5\x57\xfc\x96\x8b\x49\xeb\xba\x7e\xc3\x59\xe9\xc8\xbb\xdb\xae\x88\xf9\x31\x3b\xa3\x53\xaf\x60\x69\x47\x2d\xd5\x75\x32\x2e\xfa\xa8\xda\x55\x85\xc5\x37\x50\x1f\x28\x2e\x20\x32\xf5\x35\x7a\xc4\x82\x60\xa6\xd2\x09\xec\x98\xea\x47\x7e\xf3\xac\x40\x30\x4c\xaf\x61\x0f\x1a\xc7\xb0\x12\xb0\x9f\xb1\xc0\xf7\xfe\xcc\x9b\xe2\x08\xe2\x06\x3f\x12\x76\x18\xc4\xe2\xce\x52\x83\x7a\xc8\x04\x96\x3b\xfb\xd6\x29\x94\x4d\xe5\xdd\x36\x8a\xbc\x58\x46\x07\xf6\xd1\x13\xa6\xf3\xab\xef\xb0\xe1\x72\x10\xd0\x31\xa5\xbe\xf9\x96\x28\x9a\x01\x94\xdf\x0c\xef\x04\x97\x23\x63\x44\x22\x0b\x6c\x0a\xf6\xd7\x22\xe9\xf8\x34\x9b\xde\xb9\xd5\x15\x67\xda\x49\xd7\x40\x0d\x8f\x58\xd6\xe9\x85\x3d\xfa\x83\xb2\x0f\x8a\x99\x7d\xd1\x31\x16\x87\xac\x78\x1f\x7a\x08\x1b\xb6\x72\xd1\xde\xb6\x35\xac\xf4\x29\x97\xa2\x75\x3a\x7a\xf6\x21\x6f\x73\xc4\xe2\x00\x83\x45\xbc\x73\xef\x9d\x59\xbd\xe9\x41\xf4\x4a\x23\xc1\x35\x54\x98\x30\xc2\x0e\x39\x99\x3c\x5d\x0a\x98\x57\xc0\x2f\x1d\x49\x4b\xe6\xe0\x84\xbb\xfd\x64\x67\x22\x1d\x0e\x27\xfb\xdc\x04\x42\xae\x9f\xea\x16\x31\x5d\x2b\xd7\xa7\x5b\xe5\xd3\xaf\x93\xc6\x7b\x7b\x9c\xfd\x78\x02\x00\x49\x9c\x9a\x54\x1b\x9a\xbc\x40\xb5\x8f\xba\x3b\xb6\xe7\x11\x14\x26\x07\xe9\xe6\xb8\x60\x84\x62\x81\xd6\x5b\xfc\xbc\x40\x93\xee\x18\x83\x5a\x93\xf7\x73\xf4\x96\x72\xac\xc6\x35\x83\x87\x48\x96\x1f\x68\x89\xaf\xc1\xd1\x60\x37\x06\x7e\xbe\x0d\x06\x4b\xc3\xb2\xee\x33\x3a\x15\x13\x63\xdd\x88\x31\xb1\xe8\x4e\x82\x5d\xcf\x41\xcc\xa3\x7b\x9d\x3f\x6a\x0d\x8a\x08\x53\x20\xf6\xb8\x18\xb8\x23\xa1\x69\x6e\xdc\x03\x56\xf0\x84\x73\x1c\xc9\x90\xe2\x11\x47\x79\xe2\x3c\x04\x76\x32\xca\x9d\x11\x1b\xc7\x77\x56\xdc\x99\xf6\xd0\x60\x4a\xf6\x64\x78\x38\xe5\x7b\x7d\xf4\xe2\xce\x46\x13\x5d\x46\x82\xce\x28\x66\xa7\x35\x3b\xc3\x86\xe8\xb2\xb9\x42\x82\xb8\x61\x74\xcd\x0f\xba\xb5\x74\x75\xe0\x20\x52\xd5\x14\xf4\x2b\xf9\x17\x70\xe5\x20\x41\xf6\xf9\xa9\x7b\x9c\x64\x5a\x5d\x62\x9f\x8d\x96\xdb\xb0\x75\x7c\xfc\xec\xb0\x3a\x58\x8d\x0c\xad\x9b\xba\x25\xc8\x6e\xf8\x91\x33\x20\xd1\x17\x86\xd1\x79\x62\x32\x93\x0e\x2c\xcb\x06\xe6\x92\x81\x17\xd0\x93\xef\x61\x27\x2f\x26\x27\x2f\x24\x63\xdf\xc1\xc5\x8e\x58\x3a\xec\x4c\xd3\x81\xd0\xf9\x9e\x0e\x44\xa7\x8e\x7e\xf3\xc4\xc5\xb7\x3d\xe5\x4f\xf3\x7b\xbc\xc0\x42\x98\x00\x15\xbe\xf4\xd8\xfb\x9d\x17\x38\x93\x30\x6f\x93\x66\xd7\x47\x12\x01\xe5\x2d\xa9\xe0\x1c\xe9\xff\xbb\xfa\x52\x94\x91\xaf\xbe\x41\x1b\x52\xb0\x28\x51\x8e\x24\x7f\x83\x36\xa2\xcc\x5a\xe2\xa7\x44\x2c\x58\x0b\x79\x8e\x2a\x5c\x7f\x91\xf6\x1c\xf9\x87\xe4\xec\x97\x6b\xfc\x74\x09\x52\xe2\x03\x2c\xe8\x7c\x89\xeb\x5e\x2a\x36\x3b\x10\x4c\xcd\xbf\xc4\xf5\xc0\xf6\x40\x3c\x9d\xc3\xa4\x47\xfd\x72\xa2\xd1\x20\x8d\x67\xeb\x2c\x8d\x84\x8b\xa4\x26\x13\x31\xd0\x05\x04\x25\x43\xaa\x94\xce\x5f\x62\x53\x6a\x8d\xdc\xb1\x8d\xab\x26\xb7\x3d\x1e\xaf\xdf\x8d\xd7\xfd\x94\xaf\xcf\xf9\xf7\x3b\x56\xe8\xf0\x32\xc2\x9e\xa2\x86\x19\x1a\x93\x0e\x38\xc5\xa0\x12\x59\x07\xcb\xfb\x76\x83\xab\x1a\x93\x83\x49\x57\x56\x45\xf0\x10\xd0\xaa\x25\xd3\xbc\xb7\x9c\x6c\x4f\xa8\x02\x31\x45\xcb\x86\xdd\x97\xcc\xad\xcb\x1f\x42\x03\xe3\x78\x10\x64\x5d\x28\x6e\xa2\xf8\x1e\xa8\x65\x71\x69\x93\x73\xa9\x6f\x1c\x27\xb4\xd9\xde\x44\x06\x71\x38\x2d\x8b\x72\xa1\xde\x8b\x52\x47\x28\x47\x1f\x47\xe3\x62\x70\xc6\x06\x2e\xcc\x54\x19\x42\xba\x1a\xe1\xc7\xbc\xc9\xab\x0f\xb5\x86\x65\xd1\xb4\xc4\x91\x44\x5c\x53\x3f\xa9\x07\xf5\x13\xd3\xe8\x4a\x28\x97\x23\x35\x96\xd0\xca\x2b\x5c\x25\x0d\x92\x37\xa2\x80\xb4\xd4\xf8\xa0\xda\xa0\xa6\x37\x1f\x4f\x63\x09\x43\xb2\x06\x32\x0b\x43\x78\x97\x02\xa7\x83\xa6\xe2\xce\xbd\x76\x16\x84\x1d\x28\x18\x94\x4c\x15\x41\x7a\xa9\xd1\xea\x8d\xe0\x4f\x73\x6a\xbc\xc8\x5c\xed\xf1\x65\x81\xe9\x27\xa7\x3a\x2d\xde\x9b\xe7\xb1\x5d\x69\x63\xb3\xc3\xd3\x23\x56\xc1\xc6\xc8\x6f\x91\x3d\x11\x52\x31\x83\x82\x51\x19\x8a\xb3\x22\x31\x20\x49\x59\x52\xb8\x1a\x48\x45\x84\xdb\x46\xfb\x49\x7b\x24\xa6\x8d\x72\xdc\x60\x54\x46\x09\x80\xcc\xd4\x86\x32\x57\x62\xca\xe6\x1e\xa4\x6e\xdd\x7e\x27\x6c\x08\xd3\x82\x57\x35\x66\xed\x60\xb8\x28\xb8\x11\x35\x14\x48\x64\x6a\x2e\x55\x17\xfe\x46\xad\x36\xb9\xf8\xa4\x1e\x01\x07\x12\x04\xd2\xbc\x3d\x1a\x47\x62\xc6\x66\x2b\x33\x50\x14\x79\x0c\x28\xd4\x47\xce\xa6\xd0\x01\x95\x29\x57\x8f\xda\x9c\x05\xaa\xbd\x9e\xf1\xe5\x8a\xf9\x5b\x1e\x23\xae\x29\x90\xc2\x84\xa6\x92\x1f\xe2\x56\x1f\x40\x89\x54\x84\x1d\x36\x8d\x54\xbc\x02\x91\xb9\xd2\x79\x93\x11\xc9\x9b\x9b\x93\x4c\x82\xf6\xc4\x34\x3b\xcb\x7c\x06\x86\x15\xbc\xdf\x5f\x10\xa1\x8e\x49\x50\xc6\x52\xd6\x5c\xd8\x52\x87\x68\xf3\x8d\x57\x4d\x75\x9f\xf2\x6a\x86\x2d\x8e\x0d\x0c\x27\x17\x3e\x8e\xa2\xce\x20\x13\x6a\x0a\x33\xb7\xb5\x52\x82\xdc\x37\x0a\x02\xe6\x2a\x40\x82\x78\x84\xd2\x1c\x97\xb3\x25\xb4\xae\xda\x39\x9a\x44\x8c\xb1\xbe\x25\x05\xab\xec\x90\x7d\x45\x37\x3b\xe6\x14\x81\xf1\xd5\xd2\x51\x63\x3b\x06\x92\x8d\xfc\xbe\xe8\x3a\x9a\x7a\x5d\xf7\x12\x33\xd5\xd8\xcf\x98\x92\xd2\xf8\xf1\x1a\x64\x43\x3d\xa5\x3a\x62\xa9\xe5\x38\x7b\x23\x04\xef\xc3\x59\x42\xbe\x3b\x01\x97\x97\xfc\x06\x09\x7a\x88\x21\x42\x5a\xaf\x0c\xf7\x6a\x52\xd8\xd0\x64\xa4\xb7\xc3\x28\x1c\x2d\x50\x65\x64\x03\x76\xa4\x61\x92\x0f\x17\x63\x56\xfe\x71\x92\x1d\xe6\xa3\x6a\xaf\x41\x27\x5f\xc5\x60\x69\x88\xf4\x2d\x3d\x43\x8c\x17\xa6\xc2\xcf\x6b\x4a\xf9\x53\xd0\x8e\x7a\x1a\xd3\x79\x6f\x4b\xf6\x1d\xcb\x0a\x5a\xad\x6e\x2e\x82\x53\x6d\xa6\xae\xab\xb9\x96\xdb\x26\x7d\x16\xcd\xf5\xb3\xdf\xb5\xc9\x6a\xc4\xf7\x75\x73\xfa\xe3\x0c\xec\x2d\x17\x7e\x8f\x9d\xb9\x16\x1f\x4a\xd1\x5e\xb7\x95\x58\xe1\x33\x7b\xca\x73\x51\xd9\x00\x68\xff\xc5\x6a\x03\x7d\x56\x5b\xef\x56\xc4\xf7\x48\x36\x76\x0b\xf8\x4b\x79\x3f\xc8\x6b\x04\x55\xad\x5a\x44\xf6\xdd\xb0\x44\xa2\x47\xdd\xf7\xcc\xf1\x0f\xaf\x26\x53\x6b\xba\xd3\xc3\x05\xa0\xef\x6a\x4e\x67\x37\x47\xfe\x24\xb5\x56\x75\x04\x24\xe0\xa1\x01\xa9\xd0\x13\x96\x48\x36\x45\x01\x52\xee\x1b\x4a\x5b\x4d\x60\xf5\x03\xb8\xb1\xba\xc7\x9e\x07\x8e\x7c\x23\xe2\xae\xa1\xbb\x8b\x9e\x00\x50\xdf\x67\xf0\xe2\xa1\x33\x0a\xbc\xff\xde\x12\xa0\x25\x92\x35\x14\x64\x4f\x8a\xc0\x10\xbb\x5f\xa4\x73\xa3\x96\x32\x3b\x6d\x58\x81\x37\xca\xdf\x76\x02\x8e\xbc\x9c\xbd\x03\x06\x02\xd3\x31\x8d\x07\xdb\x3c\xa5\x73\x3a\x0a\xf4\x22\x7e\x2a\x6b\xf4\x0d\x5a\x8d\x1b\xed\x3e\x33\x16\xaa\xec\x76\xff\x05\xbd\xdf\x2b\x60\xa8\x91\x50\x6a\x48\x22\x25\x30\x93\xd4\x58\x75\xe6\x0a\x49\xf9\xe8\x75\xb6\xd6\x6b\x83\xbf\x69\xf4\x59\x95\x26\x67\x8c\x14\x2a\x8e\xe4\x91\x3f\xe9\x9f\xc0\x4a\xfd\x4e\xa0\xff\x43\x84\xa1\x02\x4b\x40\x8c\x87\xa3\x59\x76\xe0\xd6\xc0\x7d\x13\xf1\xbb\xcd\x42\xa7\x77\x60\xb2\xca\xff\x61\x73\x36\xc3\xee\x4a\x60\xca\xd6\xf4\xb5\xbd\xd8\xc6\x12\x03\xbd\x00\x85\x71\xe2\x98\x5f\xac\x61\x9c\xfa\x5f\x46\x32\x9b\x91\xf8\x3c\xe4\xff\xcf\xe7\x65\xfe\x36\x26\xf3\xdf\x9c\xb3\x98\x7c\x25\x38\x6e\x73\x32\x0b\x72\x16\xc2\xea\x46\x8d\x03\x7a\x67\x9a\x97\xa0\xfa\x4f\x04\xf5\x02\x4c\x2f\x80\xf4\x02\x44\x2f\x00\xf4\x02\x3c\x2f\x80\xf3\x02\x34\x2f\x00\xf3\x02\x2c\x2f\x80\xf2\x02\x24\x2f\x00\xf2\x02\x1c\x2f\x80\xf1\x8f\xa0\xd8\x5f\x0b\x38\x34\x87\x48\x3e\xfb\xc4\xc8\x43\x03\x1d\x2d\x35\xf9\x90\x3e\x5d\x88\x3d\x14\x5a\x73\xc0\xf9\xd6\xb3\x0c\x85\x8d\x8e\x92\xee\xe6\x71\x84\x96\x96\xb1\x25\x29\xe1\x4a\xb7\x5b\x47\x0f\x1b\x5a\x5a\x43\x74\x0a\xeb\x4e\xdd\x84\x9c\xa2\x7b\x08\xce\xdc\x23\x91\x89\xd5\x73\xf7\x19\x67\xef\x6b\x9b\x26\x23\x7f\x6d\x81\xec\xd7\xaf\xfe\xd0\x0e\x6f\xbc\x96\xf4\x48\xee\xc3\x92\x2e\xee\x92\xab\x5f\xf8\x12\x2b\x40\xbf\x22\x45\x2a\xf0\x6b\x95\x5e\x83\x8d\x5d\x99\x47\x6b\x1a\x26\x09\x7f\xae\x73\x5f\x98\x73\x44\xac\x7f\xca\xb1\xd2\xfa\xff\x47\xfd\xbb\xd8\xad\x9d\xe0\xc6\x7a\xf0\xdf\xe5\xce\xc9\xd4\xab\x4c\x16\xfb\x2f\x90\x7b\x4d\x85\x1e\xbf\xa6\x76\xc1\x16\xe3\x13\x33\xf4\xea\x95\x2f\xec\xbd\x7a\xb5\x1c\xab\x0b\x9c\x9d\x4a\x4e\x79\xdb\x84\x56\x78\x56\x9a\x73\x9b\x2d\xf8\xb1\xe9\xbf\xaf\x88\x66\x1c\x2c\x7c\xf4\xf7\x1c\xa7\x43\x51\xef\x0f\x3e\xf8\xa8\x27\xad\x0c\x39\x53\xa7\xea\x24\x48\x80\x6a\x04\xeb\x5c\xe9\xae\x7b\x34\x48\x44\x57\x33\xd1\x89\x82\x02\x51\x49\xbf\xd5\xb0\xad\x95\xa0\x87\x06\x9b\xaf\xf2\xdd\xe6\x02\x54\x98\xef\xc5\x95\xf9\x1e\x04\x61\x66\xa3\xef\x81\x3c\x02\x4b\x7c\x30\x65\xd3\x6a\xec\xd2\x2a\xfb\xd1\xef\xeb\xbc\x3f\x07\x5f\x91\xe5\xea\x43\x03\xff\x5c\x3a\xba\x93\xba\x68\x5d\x96\xb7\x5c\xab\x19\xda\xb6\xdb\x9e\xbe\xee\xaf\xbd\x16\x58\x33\xe5\xdf\x2d\x50\x50\x10\xde\xc9\xcf\x7f\x47\x3e\xaf\x6f\xa7\xa0\xea\x3e\xbf\x36\xf6\xfe\x90\xd2\x4f\xb5\x8e\x52\x5a\xe9\x47\xd5\x2e\xd0\x1b\x2c\xcf\xcc\x10\x67\xeb\xb2\x94\xbf\x5a\xfd\xd2\x60\xc7\xd7\xab\xfc\x85\x95\x03\x60\x88\xb3\x1c\xa0\xac\x8a\x38\xb0\xaf\x70\x7f\x94\xcc\x31\x88\x01\x7a\x86\x15\xb5\xdc\xce\xb4\xc3\x26\x95\xa3\x55\xfa\x51\xd9\xeb\x34\xa2\x0c\x46\xcb\xd6\x9e\x72\x03\xae\xeb\x9a\xb6\x7d\x55\xfb\xbd\xf0\x65\xea\x55\xb1\xc8\xb3\x19\x95\xd7\x50\xf1\x47\xe8\xf4\x1c\xdc\x2f\xcb\x90\x32\xaa\xaf\xb7\x71\x15\x5e\xf0\x2f\xd2\x17\xa3\x82\x33\xf8\xb5\x6a\xa8\x22\x35\x85\xee\x23\x4e\x8f\x0f\x90\xe3\x68\x48\x08\x0f\xc8\x55\x4f\x2b\xed\x8b\x41\x5d\x2b\xc7\x80\xcd\xdf\x30\x4d\xcb\x69\x78\x7c\xcd\x21\xfa\x87\x6d\x8f\xcf\x43\xb9\x92\xf1\xf3\xa8\x61\x71\xbf\x17\x4e\x61\xa3\x4f\x6a\x39\xbb\xe7\x8c\x58\x74\x7b\xf1\xaf\x00\x00\x00\xff\xff\xf0\xfc\x94\x34\xa8\x38\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"schema.graphql": schemaGraphql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
