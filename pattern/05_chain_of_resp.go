package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Паттерн "цепочка вызовов" (Chain of Responsibility) является поведенческим паттерном проектирования,
// который позволяет передавать запросы последовательно через цепочку обработчиков.
// Каждый обработчик решает, может ли он обработать запрос или должен передать его следующему обработчику в цепочке.
//
// Плюсы:
//
// - Разделение обязанностей. Каждый обработчик отвечает только за свою часть логики обработки запроса,
// что делает код более модульным и упрощает его поддержку.
// - Гибкость. Можно легко добавлять или изменять обработчики в цепочке без изменения клиентского кода.
// - Уменьшение зависимостей. Клиентский код знает только о первом объекте в цепочке,
// что снижает связанность между компонентами системы.
//
// Минусы:
//
// - Нет гарантии обработки. Если запрос не может быть обработан ни одним из обработчиков в цепочке,
// это может привести к ошибке или нежелательному поведению.
// - Перегрузка цепочки. Если цепочка становится слишком длинной или слишком сложной,
// это может усложнить отладку и понимание логики обработки.
// - Потенциальное снижение производительности. Каждый запрос должен пройти через всю цепочку,
// что может быть неэффективно в случае больших цепочек или при частых запросах.

// Main handler interface
type Handler interface {
	HandleRequest(pkg string)
	SetNext(handler Handler)
}

// PackageManager is the base class for other package managers
type PackageManager struct {
	next Handler
}

// Set next handler
func (pm *PackageManager) SetNext(handler Handler) {
	pm.next = handler
}

type Apt struct {
	PackageManager
	packages map[string]struct{}
}

// Apt constructor with his own packages
func NewApt() *Apt {
	packages := make(map[string]struct{})
	packages["wget"] = struct{}{}
	packages["htop"] = struct{}{}
	packages["firefox"] = struct{}{}

	return &Apt{
		packages: packages,
	}
}

// Handle request will check package in hashMap and print it. If there is no such package in
// this package manager, it will run HandleRequest with pkg for next handler (if next != nil).
// Otherwise, at the end of chain, just return and print that package is absent.
func (apt *Apt) HandleRequest(pkg string) {
	_, ok := apt.packages[pkg]
	if ok {
		fmt.Printf("Package <%s> is available to install in apt\n", pkg)
		return
	} else if apt.next != nil {
		apt.next.HandleRequest(pkg)
	} else {
		fmt.Printf("There is no <%s> package to install\n", pkg)
		return
	}
}

type Pip struct {
	PackageManager
	packages map[string]struct{}
}

// Pip constructor with his own packages
func NewPip() *Pip {
	packages := make(map[string]struct{})
	packages["pytorch"] = struct{}{}
	packages["numpy"] = struct{}{}
	packages["pillow"] = struct{}{}

	return &Pip{
		packages: packages,
	}
}

// Handle request will check package in hashMap and print it. If there is no such package in
// this package manager, it will run HandleRequest with pkg for next handler (if next != nil).
// Otherwise, at the end of chain, just return and print that package is absent.
func (pip *Pip) HandleRequest(pkg string) {
	_, ok := pip.packages[pkg]
	if ok {
		fmt.Printf("Package <%s> is available to install in pip\n", pkg)
		return
	} else if pip.next != nil {
		pip.next.HandleRequest(pkg)
	} else {
		fmt.Printf("There is no <%s> package to install\n", pkg)
		return
	}
}

type GoGet struct {
	PackageManager
	packages map[string]struct{}
}

// GoGet constructor with his own packages
func NewGoGet() *GoGet {
	packages := make(map[string]struct{})
	packages["gocache"] = struct{}{}
	packages["zap"] = struct{}{}

	return &GoGet{
		packages: packages,
	}
}

// Handle request will check package in hashMap and print it. If there is no such package in
// this package manager, it will run HandleRequest with pkg for next handler (if next != nil).
// Otherwise, at the end of chain, just return and print that package is absent.
func (goget *GoGet) HandleRequest(pkg string) {
	_, ok := goget.packages[pkg]
	if ok {
		fmt.Printf("Package <%s> is available to install in goget\n", pkg)
		return
	} else if goget.next != nil {
		goget.next.HandleRequest(pkg)
	} else {
		fmt.Printf("There is no <%s> package to install\n", pkg)
		return
	}
}

// Example usage
//
// func main() {
// 	apt := pattern.NewApt()
// 	apt.SetNext(nil)

// 	pip := pattern.NewPip()
// 	pip.SetNext(apt)

// 	goget := pattern.NewGoGet()
// 	goget.SetNext(pip)

// 	goget.HandleRequest("pytorch")
// 	goget.HandleRequest("zap")
// 	goget.HandleRequest("firefox")
// }
