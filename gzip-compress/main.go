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
		inputReader      io.Reader
		outputWriter     io.Writer
	)

	// parse command line
	flag.BoolVar(&enableDecompress, "d", enableDecompress, "decompress")
	flag.Parse()
	if flag.NArg() < 2 {
		errExit("missing input argument")
	}

	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)

	// open input
	if inputFile == "-" {
		inputReader = os.Stdin
	} else {
		inputReader, err = os.Open(inputFile)
		if err != nil {
			errExit(err)
		}
	}

	// open output
	if outputFile == "-" {
		outputWriter = os.Stdout
	} else {
		outputWriter, err = os.Create(outputFile)
		if err != nil {
			errExit(err)
		}
	}

	if enableDecompress {
		err = decompress(inputReader, outputWriter)
	} else {
		err = compress(inputReader, outputWriter)
	}

	if err != nil {
		errExit(err)
	}

}

func errExit(err any) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
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
