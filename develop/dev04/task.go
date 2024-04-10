package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type AnagramSet struct {
	firstWord string
	set       []string
}

func getAnagramMap(arr []string) map[string]AnagramSet {
	hashMap := make(map[string]AnagramSet)

	for _, word := range arr {
		word = strings.ToLower(word)
		// Get a key for hash map. Key is the word with sorted letters
		keyWord := SortString(word)

		anagramSet, ok := hashMap[keyWord]
		if !ok {
			// If there is no such set, then create a new set
			hashMap[keyWord] = AnagramSet{
				firstWord: word,
				set:       []string{word},
			}
			continue
		}

		// Get an insert index for new word
		idx := sort.SearchStrings(anagramSet.set, word)

		// If the word already exists in set, then skip this iteration
		if idx < len(anagramSet.set) && anagramSet.set[idx] == word {
			continue
		}

		// Otherwise expand the set
		anagramSet.set = append(anagramSet.set, "")

		// If word needs to be inserted in the mid, then cut the set and shift the values on 1 position to the right
		if idx < len(anagramSet.set) {
			copy(anagramSet.set[idx+1:], anagramSet.set[idx:])
		}

		// Insert the word at index
		anagramSet.set[idx] = word
		// Update the set value in the map
		hashMap[keyWord] = anagramSet
	}

	return hashMap
}

func formatAnagramMap(hashMap map[string]AnagramSet) map[string][]string {
	resultMap := make(map[string][]string)

	for _, set := range hashMap {
		if len(set.set) != 1 {
			resultMap[set.firstWord] = set.set
		}
	}

	return resultMap
}

// GetAnagramSet will create a set of anagrams implemented using hash map
//
// Single anagrams will be skipped and not added to the result
func GetAnagramSet(arr []string) map[string][]string {
	hashMap := getAnagramMap(arr)

	return formatAnagramMap(hashMap)
}

func main() {
	anagrams := []string{"123", "231", "231", "kEk", "eKk", "single", "mmm", "mmm", "mmm"}

	anagramSet := GetAnagramSet(anagrams)

	fmt.Println(anagramSet)
}
