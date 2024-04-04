package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Паттерн "Посетитель" (Visitor) представляет собой поведенческий паттерн, который позволяет добавлять
// новые операции к объектам без изменения самих объектов. Он достигается путем разделения алгоритма от структуры объекта,
// что позволяет добавлять новые операции, не изменяя классы объектов, к которым эти операции применяются.
//
// Плюсы:
// - Разделение алгоритма и структуры объекта. Паттерн позволяет добавлять новые операции без изменения
// существующих классов объектов, что способствует соблюдению принципа открытости/закрытости.
// - Увеличение расширяемости. Новые операции могут быть легко добавлены путем создания новых
// посетителей без изменения существующих объектов.
// - Упрощение структуры. Паттерн позволяет разделить алгоритмы от объектов, что делает код более чистым и понятным.
//
// Минусы:
// - Сложность добавления новых классов объектов. При добавлении нового типа объекта в структуру,
// требуется изменение всех существующих посетителей, чтобы поддержать новый класс.
// - Нарушение инкапсуляции. Посетитель требует открытия интерфейса объектов для выполнения операций,
// что может нарушить инкапсуляцию и повысить зависимость между классами.
// - Усложнение кода. Паттерн может усложнить код за счет создания дополнительных классов и интерфейсов.

type LinuxKernel struct {
	FHS map[string]string
}

// Method, which will accept visitors
func (k *LinuxKernel) Accept(v KernelVisitor) {
	v.visit(k)
}

type KernelVisitor interface {
	visit(k *LinuxKernel)
}

type OpenFileVisitor struct {
	FilePath string
}

func (v *OpenFileVisitor) visit(k *LinuxKernel) {
	fmt.Printf("Open file %s: %s\n", v.FilePath, k.FHS[v.FilePath])
}

type CloseFileVisitor struct {
	FilePath string
}

func (v *CloseFileVisitor) visit(k *LinuxKernel) {
	fmt.Printf("Close file %s: %s\n", v.FilePath, k.FHS[v.FilePath])
}

type ExecFileVisitor struct {
	FilePath string
}

func (v *ExecFileVisitor) visit(k *LinuxKernel) {
	fmt.Printf("Exec file %s: %s\n", v.FilePath, k.FHS[v.FilePath])
}

// Example usage
//
// func main() {
// 	fhs := make(map[string]string)
// 	fhs["/home/user/text.txt"] = "Hello, world!"
// 	fhs["/home/user/projects/pyproject.toml"] = "Your ad could be here"

// 	kernel := pattern.LinuxKernel{
// 		FHS: fhs,
// 	}

// 	ov := pattern.OpenFileVisitor{FilePath: "/home/user/text.txt"}
// 	kernel.Accept(&ov)

// 	cv := pattern.CloseFileVisitor{FilePath: "/home/user/projects/pyproject.toml"}
// 	kernel.Accept(&cv)

// 	ev := pattern.ExecFileVisitor{FilePath: "/home/user/text.txt"}
// 	kernel.Accept(&ev)
// }
