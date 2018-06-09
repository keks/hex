package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: hex [-d] [file]")
	os.Exit(1)
}

func main() {
	var (
		dec bool
	)

	flag.BoolVar(&dec, "d", false, "decode instead of encode")
	flag.Parse()

	var (
		r io.Reader
		w io.Writer = os.Stdout
	)

	switch flag.NArg() {
	case 0:
		r = os.Stdin
	case 1:
		var err error
		r, err = os.Open(flag.Arg(0))
		check(err)
	default:
		usage()
	}

	if dec {
		r = hex.NewDecoder(r)
	} else {
		w = hex.NewEncoder(w)
	}

	_, err := io.Copy(w, r)
	check(err)
}
