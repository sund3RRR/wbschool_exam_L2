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
	if len(str) == 0 {
		return str, nil
	}

	arr := []rune(str)

	if unicode.IsDigit(arr[0]) {
		return "", errors.New("Bad string")
	}

	var char rune
	var builder strings.Builder

	for _, c := range arr {
		if unicode.IsDigit(c) {
			count, _ := strconv.Atoi(string(c))

			builder.Grow(count)

			for i := 0; i < count; i++ {
				builder.WriteRune(char)
			}
			char = 0
		} else if char != 0 {
			builder.WriteRune(char)
			char = c
		} else {
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
