// Code generated by "esc -private -pkg=driver -o=bindata.go memcopy.hsaco"; DO NOT EDIT.

package driver

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/memcopy.hsaco": {
		name:    "memcopy.hsaco",
		local:   "memcopy.hsaco",
		size:    9520,
		modtime: 1600888528,
		compressed: `
H4sIAAAAAAAC/+xaTW8bRRh+d71xLKcShVMpBxaKFKjqzdoOrpsLcRJiqjqpU0PagKpovDt2ttkv7a4j
G4m0CAnl0APwAzhz4C9Awh/gwLkHLr33AkeMdndmv+rNh2iVSuwjJY/n/Zr3nR3v7ozn4cetVZZhFoEg
B38C43646LepQnzX56uerA4FWIQLUIQ8AHARuyQfMXEuEDlD/NIgz8SZ5uP6TUXaSX7MxDnq5+YKPJEn
2IQ4Uz/2jH60vjtPHZk7hV80PxcbTx05D2cHR8eThTDxCD8rxJmL+BVI/421Fc+cXpu3vPngyzmYDmqj
ssbaSrP9mW/7JgDMEDnS5L6kl5Amu387NiqV+r1hXaySuGYBoEhsS6VScRNbtmLoCzzFF3z5Gi/y94u3
sKVj1V4o8nyJX0caDm14npcMc+RbFN1mZ6R1DTViNRsaLO7Ks55RC+n9AeqHgW6bWF9u8csxbZCRl0mF
v+9pG1bfS8XFhHTkbUUv0sanIxPHLGZ7qoGcq7OBRUf5Mu5fD1QNVenrCxNVm0gd4FuKLlN1UzW6SF0a
9HrYilu5KVCr1WoljC7LVsdEEt4YIHUhCBHqJYlqfKzgHhqoTmBw0142dNsJDRxrgI8bF2Pg/A8GZkLp
6+llK7qTXvF8esXz6RUvjTxRerE3o8WeVMwLugyfKLKMdX8kb/d6NnbuHZNgbf7l9791zv1/fg79rxv6
cfOifsrvQJbVuWW1NlAdpWkpcmekSw2r/58zXDZk3LYMM3ikuU9KZPU7uK9h3fGTr4tE2bSMgUlUq8oQ
y76eqtuWsoccnG4QD07Kp4neRXu4Zxm0U54PvgbrA63TbN+xg6Eql0PNZkxDXdbQcFVFzl3D2vWT9mJW
PqwVBUEopr0/ue88l3P5594jmcjfZfo69chv//b7lV9YImOSAcN3D8iQIUOGVxFMcP8q+Cs75nj7RfgZ
vvPWevFbaTvyeSaxsuc4Lj8ej8evYv0ssEdutg/h8SGM4Vs3ywIUfvgG2KOL3ojkjy75a+SKZwNQYeHg
qwf8o1/fgIPDHBSPvEcCy+5DjtsH+P6PH1nWi8ew+X0ArpLL5esA7ScsAJtj8/ssQCWXY+umF8984j9U
vj7MZmOGDBkyZMiQIUOGDBkyZHhZoL81PyvStbuPS4SnCN8jCrrqv0D4r3/Ghss80dOV/+LM5P4kFel9
fs//fZUvlwVREPn352xLmsNDB1s6UudUdU8rmZbxAEvOnO/QvVGvzt/oliu9Hrou1mQR12tVVO2Vy9dr
3YooV1BNmq/Wux8AtBR9F1sLfKu1cpr4qiqfJfpx+yjuaB68HpdP01G+GJdf8TYfpsNzAZFxn7QPDYJu
OBgEeaTbIw2Evj4QdpC9A+S/K3csEBw8dLwW0hQJBMnQNKw7INgjzUFdEOwd27H8Tz7D0pK4XYlsV7+n
GhJSo/vX2ytb6421m8svar9pOnI8Ie2cQrB3BM+P50zEjc5bymJk3jKR8xh0Pr8GAH+Pxwb1p/M24ERa
iXS8cw5TET2d55QvJfy5BL9Nzk+wie8V5QsT51WI2ehZFkg//5IWoER8c1SQci5lKlE/7aZGQoqJbkzi
f5jSPeWPotc+AvEdn4cQ3oe4CdevGc09gp+I/9YJ47eR4s+Q803XTvD/NwAA///1H0lWMCUAAA==
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}