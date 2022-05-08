package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		enableDecompress bool
		err              error
	)
	flag.BoolVar(&enableDecompress, "d", enableDecompress, "decompress")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing input argument")
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	var inputReader io.Reader
	if inputFile == "-" {
		inputReader = os.Stdin
	} else {
		inputReader, err = os.Open(inputFile)
		if err != nil {
			errExit(err)
		}
	}

	if enableDecompress {
		err = decompress(inputReader, os.Stdout)
	} else {
		err = compress(inputReader, os.Stdout)
	}

	if err != nil {
		errExit(err)
	}

}

func errExit(args ...any) {
	fmt.Fprintln(os.Stderr, args...)
}

func compress(reader io.Reader, writer io.Writer) error {
	gzipWriter := gzip.NewWriter(writer)

	_, err := io.Copy(gzipWriter, reader)
	if err != nil {
		return err
	}
	return gzipWriter.Close()
}

func decompress(reader io.Reader, writer io.Writer) error {
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, gzipReader)
	return gzipReader.Close()
}
