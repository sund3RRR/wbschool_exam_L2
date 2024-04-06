package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Шаблон проектирования "Фабричный метод" относится к классу порождающих паттернов,
// которые используются для создания объектов. Он определяет интерфейс для создания объекта,
// но оставляет выбор конкретного класса создаваемого объекта на подклассы.
// Это позволяет классам-потомкам изменять тип создаваемых объектов.
//
// Плюсы:
//
// - Гибкость. Шаблон позволяет изменять тип создаваемого объекта, не изменяя основной код.
// - Расширяемость. Легко добавить новые варианты объектов, просто добавив новый подкласс.
// - Поддержка принципа открытости/закрытости. Существующий код остается незатронутым, пока только добавляются новые классы.
//
// Минусы:
//
// - Разрастание классов. Если каждый тип объекта требует своего подкласса фабрики, это может привести к увеличению числа классов в системе.
// - Сложность. Паттерн может усложнить структуру программы из-за большого количества классов.
// - Необходимость переопределения. В подклассах необходимо переопределить метод создания объекта,
// что может быть неудобным или привести к дублированию кода.

const (
	Fedora = iota
	Ubuntu
)

// Base factory interface
type Factory interface {
	CreateDistribution(kernel string) DistributionProduct
}

type DistributionProduct interface {
	Run()
}

//
// Fedora factory and image
//

type FedoraFactory struct {
}

func (ff *FedoraFactory) CreateDistribution(kernel string) DistributionProduct {
	return &FedoraImage{
		kernel: kernel,
	}
}

type FedoraImage struct {
	kernel string
}

func (f *FedoraImage) Run() {
	fmt.Printf("Starting Fedora Linux on kernel %s\n", f.kernel)
}

//
// Fedora factory and image
//

//
// Ubuntu factory and image
//

type UbuntuFactory struct {
}

func (uf *UbuntuFactory) CreateDistribution(kernel string) DistributionProduct {
	return &UbuntuImage{
		kernel: kernel,
	}
}

type UbuntuImage struct {
	kernel string
}

func (u *UbuntuImage) Run() {
	fmt.Printf("Starting Ubuntu Linux on kernel %s\n", u.kernel)
}

//
// Ubuntu factory and image
//

// Function return new factory based on distribution enum
func GetDistributionFactory(distribution int) Factory {
	switch distribution {
	case Fedora:
		return &FedoraFactory{}
	case Ubuntu:
		return &UbuntuFactory{}
	default:
		return nil
	}
}

// Example usage
//
//
// func main() {
// 	fedoraFactory := pattern.GetDistributionFactory(pattern.Fedora)
// 	fedora := fedoraFactory.CreateDistribution("6.8.0")
// 	fedora.Run()

// 	ubuntuFactory := pattern.GetDistributionFactory(pattern.Ubuntu)
// 	ubuntu := ubuntuFactory.CreateDistribution("5.15.187")
// 	ubuntu.Run()
// }
