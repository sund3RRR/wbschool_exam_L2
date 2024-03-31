package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type LinuxKernel struct {
	FHS map[string]string
}

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
