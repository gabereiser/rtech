package fs

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
)

//go:embed assets/*
var embeded embed.FS

type FileSystem interface {
	OpenFile(path string) (io.Reader, error)
	WriteFile(path string, b []byte) error
	AddArchive(path string) error
	Close()
}

type RFileSystem struct {
	archives map[string]Archive
}

func NewFilesystem() FileSystem {
	return &RFileSystem{
		archives: make(map[string]Archive),
	}
}
func (fs *RFileSystem) OpenFile(path string) (io.Reader, error) {
	buf, e := embeded.ReadFile(path)
	if e == nil && len(buf) > 0 {
		return bytes.NewReader(buf), e
	}
	// not in the embeded, let's search the archives...
	for _, a := range fs.archives {
		fp, e := a.GetFile(path)
		if e == nil && len(fp) > 0 {
			return bytes.NewReader(fp), e
		}
	}
	// not in the archives, or embeded, let's find it on disk...
	fp, e := os.OpenFile(path, os.O_RDONLY, 0755)
	if e != nil {
		return nil, e
	}
	return bufio.NewReader(fp), nil
}

func (fs *RFileSystem) WriteFile(path string, b []byte) error {
	return os.WriteFile(path, b, 0755)
}

func (fs *RFileSystem) AddArchive(path string) error {
	if _, ok := fs.archives[path]; ok {
		return fmt.Errorf("archive already loaded")
	}
	a, e := OpenArchive(path)
	if e != nil {
		return e
	}
	fs.archives[path] = a
	return nil
}

func (fs *RFileSystem) Close() {
	for _, a := range fs.archives {
		a.Close()
	}
}
