package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CmdFlags struct {
	filepath  string
	fields    []int
	delimiter string
	separated bool
}

func parseFields(text string) ([]int, error) {
	splitted := strings.Split(text, ",")

	fields := make([]int, 0, len(splitted))

	for _, str := range splitted {
		field, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	return fields, nil
}

func NewCmdFlags() *CmdFlags {
	// Create cmd flags for cli utility
	fields_str := flag.String("f", "", "list of integer fields")
	delimiter := flag.String("d", "\t", "delimiter")
	separated := flag.Bool("s", false, "print lines only with separator")

	flag.Parse()

	// Get file path
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: cut -<parameters> <filename>")
		os.Exit(1)
	}

	filepath := args[0]

	fields, err := parseFields(*fields_str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &CmdFlags{
		filepath:  filepath,
		fields:    fields,
		delimiter: *delimiter,
		separated: *separated,
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

	fmt.Print(Cut(text, cmdFlags))
}
