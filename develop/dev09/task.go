package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"os"
)

type CmdFlags struct {
	url   string
	dir   string
	depth int
}

func NewCmdFlags() *CmdFlags {
	// Create cmd flags for cli utility
	dir := flag.String("o", "downloaded", "outputs download directory")
	depth := flag.Int("d", 1, "depth")

	flag.Parse()

	// Get file path
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: wget -<parameters> <filename>")
		os.Exit(1)
	}

	url := args[0]

	return &CmdFlags{
		url:   url,
		dir:   *dir,
		depth: *depth,
	}
}

func main() {
	cmdFlags := NewCmdFlags()

	wget := NewWget(cmdFlags.dir)

	err := wget.DownloadPage(cmdFlags.url, cmdFlags.depth)
	if err != nil {
		fmt.Printf("Error downloading page: %v\n", err)
		os.Exit(1)
	}
}
