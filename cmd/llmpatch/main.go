package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/icholy/llmpatch"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "file to edit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -f <filename> <llm command> [args...]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()
	if filename == "" || flag.NArg() == 0 {
		flag.Usage()
	}
	// create the prompt
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	code := string(data)
	prompt := fmt.Sprintf("%s\n---\nCode:\n\n%s\n", llmpatch.Prompt, code)
	// run it through the llm
	var output strings.Builder
	cmd := exec.Command(flag.Arg(0), flag.Args()[1:]...)
	cmd.Stdin = strings.NewReader(prompt)
	cmd.Stderr = os.Stderr
	cmd.Stdout = io.MultiWriter(&output, os.Stdout)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// apply the changes to the code
	edits := llmpatch.Extract(output.String())
	code = llmpatch.Apply(code, edits)
	// write the file
	info, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(filename, []byte(code), info.Mode()); err != nil {
		log.Fatal(err)
	}
}
