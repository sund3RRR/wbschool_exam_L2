package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CmdFlags struct {
	filepath string
	key      int
	numeric  bool
	reverse  bool
	unique   bool
}

func NewCmdFlags() *CmdFlags {
	// Create cmd flags for cli utility
	key := flag.Int("k", 1, "integer value for sorting key")
	numeric := flag.Bool("n", false, "sort numerically")
	reverse := flag.Bool("r", false, "reverse the result of comparisons")
	unique := flag.Bool("u", false, "print unique strings")

	flag.Parse()

	// Get file path
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Usage: sort -k <integer> -n -r <filename>")
		os.Exit(1)
	} else if *key < 1 {
		fmt.Println("Column key must be non-zero positive integer")
		os.Exit(1)
	}

	filepath := args[0]

	return &CmdFlags{
		filepath: filepath,
		key:      *key,
		numeric:  *numeric,
		reverse:  *reverse,
		unique:   *unique,
	}
}

func ParseFile(filename string) ([][]string, error) {
	// Read file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split content on lines
	lines := strings.Split(string(content), "\n")

	// Create data slice
	data := make([][]string, 0, len(lines))

	for _, line := range lines {
		// Split line on words
		words := strings.Fields(line)
		if len(words) > 0 {
			data = append(data, words)
		}
	}

	return data, nil
}

func getSortedString(data [][]string, params *SortParams) *string {
	result := make([]string, 0, len(data))
	lastStr := ""

	for _, strSlice := range data {
		str := strings.Join(strSlice, " ")

		if params.unique && str == lastStr {
			continue
		}

		result = append(result, str)
		lastStr = str
	}

	resultStr := strings.Join(result, "\n")
	return &resultStr
}

func main() {
	cmdFlags := NewCmdFlags()

	data, err := ParseFile(cmdFlags.filepath)
	if err != nil {
		log.Fatal(err)
	}

	params := SortParams{
		colon:       cmdFlags.key - 1,
		descending:  cmdFlags.reverse,
		numericSort: cmdFlags.numeric,
		unique:      cmdFlags.unique,
	}

	data = sortData(data, &params)

	fmt.Println(*getSortedString(data, &params))
}
