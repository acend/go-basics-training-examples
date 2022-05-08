package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	//err := convertServer()

	err := tgzToZIP(os.Stdin, os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func tgzToZIP(reader io.Reader, writer io.Writer) error {

	zipWriter := zip.NewWriter(writer)

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzipReader)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		zipFileWriter, err := zipWriter.Create(hdr.Name)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFileWriter, tarReader)
		if err != nil {
			return err
		}
	}

	return zipWriter.Close()
}
