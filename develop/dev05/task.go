package main

import (
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CmdFlags struct {
	filepath   string
	pattern    string
	after      int
	before     int
	count      int
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	color      string
}

func NewCmdFlags() *CmdFlags {
	// Create cmd flags for cli utility
	after := flag.Int("A", 0, "integer value for number of strings after")
	before := flag.Int("B", 0, "integer value for number of strings before")
	context := flag.Int("C", 0, "integer value for number of strings after+before")
	count := flag.Int("c", -1, "integer value for number of strings")
	ignoreCase := flag.Bool("i", false, "boolean value for case ignore")
	invert := flag.Bool("v", false, "boolean value to invert result")
	fixed := flag.Bool("F", false, "boolean value to exact match")
	lineNum := flag.Bool("n", false, "boolean value for line num printing")

	flag.Parse()

	// Get file path
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("Usage: grep -<parameters> <pattern> <filename>")
		os.Exit(1)
	} else if *after < 0 || *before < 0 || *context < 0 {
		fmt.Println("You can't set number of strings < 0")
		os.Exit(1)
	}

	if *context != 0 {
		*after = *context
		*before = *context
	}

	pattern := args[0]
	filepath := args[1]

	return &CmdFlags{
		filepath:   filepath,
		pattern:    pattern,
		after:      *after,
		before:     *before,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		color:      Red,
	}
}

func openFile(filepath string) (string, error) {
	text, err := os.ReadFile(filepath)
	return string(text), err
}

func main() {
	cmdFlags := NewCmdFlags()

	text, err := openFile(cmdFlags.filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(Grep(text, cmdFlags))
}
