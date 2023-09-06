package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: rcc <output> <input>")
		fmt.Println("where <output> is the name of the archive file to write and <input> is a directory to compile")
		fmt.Println("")
		fmt.Println("\texample: rcc archive.red ./assets")
		flag.PrintDefaults()
	} else {
		compile(args...)
	}
}

func compile(args ...string) {
	files := make([]string, 0)
	compiledFiles := make([]string, 0)
	filepath.Walk(args[1], func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(path)
			switch ext {
			case ".png", ".jpg", ".bmp", ".mp3", ".aac", ".wav", ".glsl":
				files = append(files, path)
			case ".gltf", ".glb":
				compiledFiles = append(compiledFiles, compileAsset(path))
			}
		}
		return nil
	})
	if len(files) > 0 {
		writeArchive(args[0], args[1], files)
	}
}

func writeArchive(filename, prefix string, files []string) {
	fp, e := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if e != nil {
		panic(e)
	}
	defer fp.Close()
	gw := gzip.NewWriter(fp)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, file := range files {
		archiveAdd(tw, prefix, file)
	}
	tw.Flush()
}

func archiveAdd(tw *tar.Writer, prefix, file string) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	info, err := fp.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	filename := strings.TrimPrefix(file, prefix)
	if !strings.HasPrefix(filename, "/") {
		filename = "/" + filename
	}
	header.Name = filename

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, fp)
	if err != nil {
		return err
	}

	return nil
}

func compileAsset(path string) string {
	doc, e := gltf.Open(path)
	if e != nil {
		panic(fmt.Errorf("error compiling asset: %v", e))
	}
	for _, m := range doc.Meshes {
		for _, p := range m.Primitives {
			pIdx := *p.Indices
			indexAccessor := doc.Accessors[int(pIdx)]
			meshIndices, err := modeler.ReadIndices(doc, indexAccessor, nil)
			positionAttr := p.Attributes[gltf.POSITION]
			normalAttr := p.Attributes[gltf.NORMAL]
			binormalAttr := p.Attributes[gltf.TANGENT]
			texcoordAttr := p.Attributes[gltf.TEXCOORD_0]
			position, err := modeler.ReadPosition(doc, doc.Accessors[positionAttr], doc.Accessors[positionAttr].Sparse.Values.BufferView)
		}
	}
}
