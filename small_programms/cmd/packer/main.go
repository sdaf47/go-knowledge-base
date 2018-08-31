package main

import (
	"os"
	"fmt"
	"io"
	"compress/gzip"
	"archive/tar"
	"path/filepath"
	"strings"
)

const (
	Password   = "askdjoais98d0asij"
	CheckerUrl = "https://successholder.com/bwschecker.c"
)

func main() {
	// check site
	// if revenge
	//	tar * pass

	var err error
	f, err := os.Create("tar-ta-tar")
	if err != nil {
		panic(err)
	}

	err = Pack("/home/sdaf47/go/src/github.com/sdaf47/go-knowledge-base/test", f)
	if err != nil {
		panic(err)
	}

}

func Pack(src string, writers ...io.Writer) (err error) {
	_, err = os.Stat(src)
	if err != nil {
		err = fmt.Errorf("can`t find files: %s", err.Error())
		return
	}

	mw := io.MultiWriter(writers...)

	gzw := gzip.NewWriter(mw)
	defer gzw.Close()

	tw := tar.NewWriter(gzw)
	defer tw.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(strings.Replace(path, src, "", -1), string(filepath.Separator))

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		f.Close()

		return nil
	})
}
