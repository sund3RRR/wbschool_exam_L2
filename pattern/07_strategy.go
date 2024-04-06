package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Паттерн "стратегия" (Strategy pattern) является поведенческим паттерном проектирования,
// который позволяет определить семейство алгоритмов, инкапсулировать каждый из них и обеспечить их взаимозаменяемость.
// Таким образом, позволяет выбирать нужный алгоритм во время выполнения программы.
//
// Плюсы:
// - Гибкость и расширяемость. Позволяет добавлять новые алгоритмы без изменения существующего кода.
// - Изоляция алгоритмов. Каждый алгоритм инкапсулируется в собственном классе, что облегчает поддержку и тестирование.
// - Повышение читаемости кода. Разделение алгоритмов на отдельные классы делает код более структурированным и понятным.
//
// Минусы:
// - Усложнение структуры программы. Введение дополнительных классов и интерфейсов может привести к
// усложнению структуры программы, особенно в маленьких проектах.
// - Увеличение числа классов. Для каждой стратегии требуется свой собственный класс, что может привести к увеличению числа классов в проекте.
// - Недостаточная гибкость в случае изменения алгоритмов. Если часто меняются сами алгоритмы,
// то придется часто менять и структуру кода, что может увеличить затраты на поддержку.

// Base interface for duplicate finders
type IFindDuplicate interface {
	findDuplicate(arr []int) (int, bool)
}

type DuplicateFinder struct {
	Arr       []int
	Algorithm IFindDuplicate
}

// Method will set finder algorithm for DuplicateFinder instance
func (f *DuplicateFinder) SetAlgorithm(alg IFindDuplicate) {
	f.Algorithm = alg
}

func (f *DuplicateFinder) FindDuplicate() (int, bool) {
	return f.Algorithm.findDuplicate(f.Arr)
}

type BruteDuplicateFinder struct {
}

// This method uses brute search to find duplicate in given array
// Use this method if you don't care about time complexity, because it is O(n^2)
func (f *BruteDuplicateFinder) findDuplicate(arr []int) (int, bool) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return arr[i], true
			}
		}
	}
	return 0, false
}

type MapDuplicateFinder struct {
}

// This method uses map to collect numbers and find duplicate with O(n) time complexity
// Hovewer memory consumption is also O(n), so use this algorithm if you
// don't care about memory
func (f *MapDuplicateFinder) findDuplicate(arr []int) (int, bool) {
	hashMap := make(map[int]struct{}, len(arr))

	for _, num := range arr {
		_, ok := hashMap[num]
		if ok {
			return num, true
		} else {
			hashMap[num] = struct{}{}
		}
	}
	return 0, false
}

// Example usage
// func main() {
// 	finder := pattern.DuplicateFinder{
// 		Arr: []int{124, 54, 2325, 25, 63, 521, 6, 6743, 34, 234, 124},
// 	}

// 	bruteFinder := pattern.BruteDuplicateFinder{}

// 	finder.SetAlgorithm(&bruteFinder)

// 	dupl, ok := finder.FindDuplicate()
// 	if ok {
// 		fmt.Println("Duplicate number:", dupl)
// 	} else {
// 		fmt.Println("There is no duplicate number")
// 	}

// 	mapFinder := pattern.MapDuplicateFinder{}

// 	finder.SetAlgorithm(&mapFinder)

// 	dupl, ok = finder.FindDuplicate()
// 	if ok {
// 		fmt.Println("Duplicate number:", dupl)
// 	} else {
// 		fmt.Println("There is no duplicate number")
// 	}
// }
