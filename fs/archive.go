package fs

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

type RArchive struct {
	fp *os.File
	ar *tar.Reader
	gz *gzip.Reader
}

func (a *RArchive) GetFile(path string) ([]byte, error) {
	buf := make([]byte, 0)
	for {
		header, e := a.ar.Next()
		if e == io.EOF {
			break
		} else if e != nil {
			return buf, e
		}
		if header.Name == path {
			for {
				var b [1]byte
				_, e = a.ar.Read(b[:])
				if e == io.EOF {
					break
				} else if e != nil {
					return buf, e
				}
				buf = append(buf, b[0])
			}
			break
		}
	}
	return buf, nil
}
func (a *RArchive) Close() {
	a.gz.Close()
	a.fp.Sync()
	a.fp.Close()
}

func OpenArchive(path string) (Archive, error) {
	fp, e := os.OpenFile(path, os.O_RDONLY, 0755)
	if e != nil {
		return nil, e
	}
	gr, e := gzip.NewReader(fp)
	if e != nil {
		return nil, e
	}
	tr := tar.NewReader(gr)
	a := &RArchive{
		fp: fp,
		gz: gr,
		ar: tr,
	}
	return a, nil
}

type Archive interface {
	GetFile(path string) ([]byte, error)
	Close()
}
