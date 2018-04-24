// Code generated by fileb0x at "2018-04-23 13:04:28.876193434 +0300 MSK m=+0.003757517" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(9d2389f883ebcdcd9e28c5441aa18b5d.f7d60b70267ef579d9aad25918a6b087)

package static

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path"

	"context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileSwaggerJSON is "swagger.json"
var FileSwaggerJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x5b\xdd\x6f\xdb\x36\x10\x7f\xf7\x5f\x41\x70\x7b\x4c\xa2\x61\x18\xf6\xd0\xb7\xa0\x19\x3a\x63\xed\x1a\x24\x4d\x31\x60\x28\x0c\x5a\x3a\x2b\x6c\x25\x52\x25\xa9\x7c\x2c\xc8\xff\x3e\x90\xd4\x07\x29\x8b\x72\xac\x58\x8d\xd7\xb9\x4f\xb1\xc4\xbb\xfb\xf1\xee\x78\x5f\x62\x1f\x66\x08\xe1\x98\x33\x59\xe6\x20\xf1\x2b\xf4\xf7\x0c\x21\x84\x30\x29\x8a\x8c\xc6\x44\x51\xce\xa2\xcf\x92\x33\x3c\x43\xe8\xd3\x91\x5e\x5b\x08\x9e\x94\xf1\xd3\xd6\xca\x5b\x92\xa6\x20\xf0\x2b\x84\x7f\x3e\xf9\x09\x9b\x67\x94\xad\x38\x7e\x85\x1e\x2c\x6d\x02\x32\x16\xb4\xd0\xb4\x7a\xd5\x39\x88\x9c\x4a\x49\x39\x93\x48\x82\xb8\xa1\x31\x20\xa9\xb8\x00\x89\x4a\x09\x02\x09\x90\xbc\x14\x31\x48\x44\x58\x82\x48\x96\xf1\x5b\x89\x14\x47\x39\x61\x24\x05\x44\xe2\x18\xa4\x79\xd0\x2c\x34\x42\x11\xc2\x8a\xaa\x0c\xba\x22\x4e\xcf\xe7\x1a\xec\xa3\xdd\x18\x51\xd7\xb2\x45\x16\x59\x66\xcd\x03\x84\x70\x0a\xca\xf9\xa9\xf7\x57\xe6\x39\x11\xf7\x9a\xed\x05\xa8\x52\xb0\x0a\xa6\x25\x05\x1f\xc9\x49\x05\xc5\x50\xf2\x02\x84\xd1\xd8\x3c\xd1\xd4\x6f\x40\x5d\xd4\xeb\x4e\x2b\x62\x77\x79\x41\x04\xc9\x41\x81\x68\xd5\x6e\xff\x3d\x38\x7f\x23\x84\x7f\x14\xb0\xd2\xfc\x7e\x88\x5a\x8a\xe8\x4a\x82\x98\x9f\xfd\x0e\x24\x01\x81\x9d\xf5\x8f\x47\x23\x18\x5d\xf0\x0c\x76\xc0\xea\xb2\x5c\x4a\x45\x55\xa9\x20\xb1\xf0\x3c\x6e\xcd\xdf\x9f\x66\xee\x93\x4a\x0a\x8e\x48\x92\x53\x16\xd5\x5a\x76\x4d\x54\x94\x61\x13\x9d\x4a\x49\x53\x56\x3b\x49\x06\x37\x90\xa1\x15\x17\xda\x8d\x3a\xde\x75\x82\xae\x24\x24\xe6\xe5\x92\x66\x19\x65\x29\x2a\x4a\x51\x70\x39\x6c\xc5\xcb\x83\x15\x9f\xc8\x8d\x91\xdc\x1c\xc6\x25\x4f\xee\xf1\x91\xff\x8e\xb2\xd0\x1b\x19\x5f\x43\x4e\x3c\xfb\xae\xa1\x4b\x60\x45\x19\xd5\x36\x91\x51\x9f\x3d\x2e\xe0\x6b\x09\x52\x61\x8f\xc3\xe3\x36\xce\xa7\xb1\xcb\x82\xc4\x20\xa3\x87\x8c\x2c\x21\x7b\xec\x09\x15\x43\x7e\xf8\x46\x10\xa6\x50\xc3\x06\x15\x4d\x44\xd2\xf1\x42\x7b\xe2\x06\x27\xfb\xb3\x26\xb5\x9b\x3a\xb8\xd8\x5e\xb8\x98\x35\xc6\x66\xff\x1a\xc0\xdc\xc9\x86\x8d\x9d\x91\xf1\xb3\x2e\xd6\x7a\x87\xbd\x2f\xed\x16\x75\x4e\xdb\x10\x59\x6b\x3a\x9c\x40\x06\x0a\x82\x5e\x7b\x66\x5e\x8f\x76\x5b\x4b\x7e\xf0\xdc\xbd\xf1\x5c\x6b\x90\xff\xb8\xf3\xfa\x91\xf9\x86\x67\xba\x8a\x7d\x56\x58\xb6\x3c\xb6\x8f\xc9\x1f\x0d\xdd\xc1\xad\xbf\xeb\x80\x6c\x8d\xbc\x47\xd1\x78\x94\xb7\x5a\xda\x83\xc3\x1e\xe2\xf0\x8e\xe3\x70\xd3\x4a\x3b\xfb\x6b\x1b\x6a\xbb\xb5\xb7\xba\xed\x72\x43\xb2\xba\x2f\x0c\x00\xa9\x04\x65\x69\x83\x00\x03\x2b\x73\xcf\x03\x31\xe3\x0c\x5c\x4f\x15\x40\x92\xee\xef\xea\xd8\x38\x4f\x6f\x05\xf5\x1f\xf0\x5b\xd6\x7a\xd4\x27\x2f\x83\x58\x88\xef\x97\x9f\x21\x56\x7d\x18\xb9\x7d\xd3\x60\x2c\x84\x3e\x5a\x8a\x7a\x3d\x28\x42\x78\x2d\xf1\x0c\xd8\xdf\x55\xcb\xac\xc7\xbe\x98\x26\x5d\x46\x01\x95\x99\x77\x2b\x2e\x72\xa2\xc1\xe3\xb2\xa4\x49\x3f\x47\x6b\xeb\x61\xa6\xb3\xae\xb5\xfd\x5c\x3b\xe4\xb5\x8e\xe2\x3a\xbe\x68\xa9\xf4\x29\xf4\x28\x10\x95\x88\x20\x51\xfd\xb0\x4a\x36\xcd\xb7\xb1\xa6\xee\xbe\xd7\x07\x3b\xe6\xbd\x8e\x74\xad\x35\xb6\xb5\x92\xa6\xae\xdc\x7f\x84\x7a\x21\x27\xd4\x3f\x30\xf8\xee\x38\xe5\xc7\xf5\x81\xd2\xbb\xd4\x67\xaf\x47\x93\x0d\x32\x8f\x20\xa0\x1c\xec\xaf\x2e\x48\xfc\x85\xa4\x86\x20\xa5\xea\x24\xe6\x4c\x11\xca\x40\x94\xf9\x09\x03\x15\xc5\xd7\x51\x9b\x0c\x64\x54\x7c\x49\xa3\x9c\x27\xb5\x6b\xd5\xd6\x6b\x62\xc2\x80\xad\xda\xb8\x61\x5f\x2c\x41\xb6\xc5\xff\x78\xa5\xc7\x45\x19\xd2\x37\x65\x0a\x52\xc7\xa0\x1d\x85\x53\xa6\x7e\xfd\x65\x40\xe1\xaf\xcf\xaf\xfa\xdd\x3d\x16\x40\x14\x2c\x14\x1d\x6b\xea\x84\x28\x38\x36\xe4\x03\xd2\x8d\x90\x0f\xd4\x33\xb8\x03\xc2\x06\xa6\x89\x41\x58\x17\xda\x04\x22\x18\x4f\x96\x9c\x67\x40\xd8\x46\x01\xc9\x04\x81\x2a\x2c\xd3\x4d\xd8\xdb\x46\xb1\x30\xd7\xb7\x86\xba\x97\x71\x4e\xee\x16\x70\xa7\x74\x70\xc8\x16\xd5\x14\x5a\x4e\xe4\xb4\xef\xc8\xdd\x6f\x77\xea\xb2\x16\x12\xc4\xa3\xc5\x7c\x23\x3c\x73\xf6\x04\x3c\x4a\x90\xd5\x8a\xc6\xd3\xa1\xf8\x50\x09\xe8\x45\x60\x52\xf8\x42\xc7\xef\xc5\x44\x4e\xf7\x5e\x4b\xe8\x96\x8b\x2e\x04\x41\xf2\x89\x36\x7f\x71\xfa\xae\x5f\xa4\x22\x82\xae\x56\x53\xed\xf8\x83\xe1\x1e\xda\xae\xd6\x75\xd0\xe7\x88\x10\xc4\xaf\x83\x31\x55\x90\xcb\xb5\x22\x38\x50\x02\xb5\x9f\x64\x82\xa5\xb9\x8f\xd5\xf9\x86\xd3\x0f\xb7\x9a\x04\x4c\x05\xd8\xf6\x4e\x4f\x04\xfb\xb1\xc2\xb2\xa9\x0e\xd8\x49\x66\x77\x54\x19\x4e\xed\xed\x22\x24\xa0\x10\x20\x81\x29\x89\x28\xb3\xbe\xa2\x1f\x93\x25\x2f\x95\xfd\x28\xe2\xf7\x96\x75\xf9\x35\xbe\x00\xd8\x69\x6d\x6c\x99\x2d\xcc\xb7\x9c\x45\x7c\x4d\x58\x3a\x75\x8a\x75\x20\xbd\x36\xf2\xc2\xd9\x76\x2f\xea\x8e\x2f\x94\x25\xa3\x93\x64\x5d\xdf\xff\xa1\x99\xf4\x27\x61\x9e\x52\x36\x9a\xbf\x0e\xb0\x6f\x0d\x87\x5e\xe6\x0c\x6e\x17\xae\x85\x77\xe4\x34\xda\xa3\x17\xdf\xb8\x58\xa9\xcf\xcd\x54\x72\x6b\x4b\x0d\xc5\xef\x2d\x65\xaf\x59\xca\xe3\x3d\x61\x0c\xf3\xbb\x4a\xff\xfb\xee\xb6\xd1\x86\xf5\x74\x3a\x3b\x4d\x05\xde\xc8\xa0\x77\x48\xb2\x9e\x97\x5e\x10\x4b\x6f\x33\x1f\x9c\x99\x86\x53\xc8\x25\xa8\xa7\xb7\xf1\x12\xd4\x74\x5d\xfc\x4e\xf3\xc9\x9e\x8d\x04\xfa\xb4\x3c\xc5\x3c\x60\xe8\x43\xf9\x8e\x7d\xc0\xbb\x1f\xf3\x92\x56\xdf\x42\xf3\x30\xa9\xee\x15\x17\x96\x49\x50\xcf\x76\x85\x33\x87\xa9\xca\x5b\x73\x35\x4a\xd3\x8e\x56\xe3\x37\xce\x82\xb4\xd8\x41\x49\x1e\x1a\x54\x0e\x16\xe1\xf3\xf3\x40\xa7\xb0\xed\x69\xf7\xb8\x76\x0e\xb3\x97\xef\xcd\x9d\xb8\xa9\xda\xf4\x8b\x9a\x7d\xaf\x70\x49\xff\x81\x89\x04\x5f\x6a\xd6\xa1\xd0\x99\x4c\x24\xf4\x4a\x86\x46\x4e\xff\xdb\x3e\xef\x63\xb7\x90\x08\x7c\xaa\xec\xc6\x8c\xe7\x44\x5c\x45\x6f\x60\xfc\xe8\xf0\xd4\xd2\xf7\x77\x4a\xa4\x20\x31\x55\xf7\x53\x4d\x86\x6b\xf6\x87\xf1\xf0\x0b\x8e\x87\xd3\xac\x94\x0a\xc4\xe2\x59\x11\xf7\x8d\x65\x12\x0e\xbc\xdf\xcb\x10\xba\x69\x5c\xa6\xea\x18\x9b\xef\x3b\xa1\xad\xed\xc3\x90\xf5\x45\xb3\xa8\x2d\xae\xa6\xda\x7c\x55\xd5\x85\xb6\x7e\x18\xf6\x0e\x0f\x7b\xa7\xc8\xad\xce\x7f\x05\x70\xee\xc0\x54\xd7\x17\xd6\xaf\x91\x3c\xe1\x12\x43\xc0\x20\xdd\x64\xad\x19\xa2\xf9\x99\xe9\x8c\x4b\x86\x74\xaa\xe5\x4c\xda\xcb\xe9\x49\x4e\xdb\x90\xdb\x5c\xd7\xd0\x26\x3a\x76\x39\xda\xab\x1a\x5f\x4b\x10\xf7\x7e\xa5\xe0\x5d\xc9\x79\x06\xe4\x5a\xf2\x5f\xc7\x9a\xe3\xf1\xfc\xac\x2b\xfb\xda\x8a\x68\x9e\xea\x26\x90\x0a\x93\x4e\x94\x28\x61\x0d\x94\x73\xbd\xc7\x81\xb5\x7e\xed\xc3\x1b\x0f\xe8\x42\xc4\x68\xa4\xbe\xbf\x71\xb4\x69\x3f\x1d\xdc\x5a\xea\x28\xe4\xda\x39\x66\x8f\xff\x06\x00\x00\xff\xff\x65\x2c\xc8\x49\x2f\x33\x00\x00")

func init() {
	if CTX.Err() != nil {
		panic(CTX.Err())
	}

	var err error

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileSwaggerJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
