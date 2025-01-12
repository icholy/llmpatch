package main

import (
	"flag"

	"github.com/icholy/llmpatch"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "file to edit")
}
