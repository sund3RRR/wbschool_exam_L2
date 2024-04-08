package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(str string) (string, error) {
	// Check if string is empty, then return string
	if len(str) == 0 {
		return str, nil
	}

	// Convert string to array of runes because some unicode symbols
	// can be presented as multiple bytes
	arr := []rune(str)

	// Check if string starts with number, so return error
	if unicode.IsDigit(arr[0]) {
		return "", errors.New("Bad string")
	}

	var char rune
	var builder strings.Builder

	for _, c := range arr {
		if unicode.IsDigit(c) {
			// If char is digit, then grow buffer and write last char n count
			count, _ := strconv.Atoi(string(c))

			builder.Grow(count)

			for i := 0; i < count; i++ {
				builder.WriteRune(char)
			}
			char = 0
		} else if char != 0 {
			// If char is letter and last char is presented, then it means that it is a single char,
			// so write it to buffer
			builder.WriteRune(char)
			char = c
		} else {
			// If last char is not presented, then just set last char
			char = c
		}
	}

	if char != 0 {
		builder.WriteRune(char)
	}

	result := builder.String()

	return result, nil
}
func main() {
	str, err := UnpackString("a4bc2d5e")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(str)
}
