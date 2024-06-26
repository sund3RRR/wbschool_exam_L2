package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Основная цель паттерна "Строитель" состоит в создании объекта пошагово.
// Этот паттерн позволяет создавать различные конфигурации объекта,
// избегая перегрузки конструктора с большим числом параметров и
// обеспечивая более гибкую и интуитивно понятную конфигурацию.
//
// Плюсы:
// - Гибкость и расширяемость. Позволяет создавать различные варианты объектов
// без изменения его структуры.
// - Избегание перегрузки конструктора. Избегает проблемы с большим числом параметров конструктора,
// что улучшает его читаемость и поддерживаемость.
//
// Минусы:
// - Усложнение кода/ Для каждого типа объекта может потребоваться создание отдельного строителя,
// что может усложнить структуру кода.
// - Дублирование кода. Некоторая часть кода может дублироваться между различными строителями,
// если они создают объекты с похожей структурой.

// Main object that needs to build
type LinuxDistribution struct {
	Name               string
	Kernel             string
	WindowManager      string
	DesktopEnvironment string
	Packages           []string
	Version            string
}

// Builde interface
type ILinuxBuilder interface {
	AddName(name string) ILinuxBuilder
	AddKernel(kernel string) ILinuxBuilder
	AddWM(wm string) ILinuxBuilder
	AddDE(de string) ILinuxBuilder
	AddPkgs(pkgs []string) ILinuxBuilder
	AddVersion(version string) ILinuxBuilder
	Build() *LinuxDistribution
}

// Builder
type LinuxBuilder struct {
	name               string
	kernel             string
	windowManager      string
	desktopEnvironment string
	packages           []string
	version            string
}

// Builder's constructor
func NewLinuxBuilder() ILinuxBuilder {
	return &LinuxBuilder{}
}

func (b *LinuxBuilder) AddName(name string) ILinuxBuilder {
	b.name = name
	return b
}
func (b *LinuxBuilder) AddKernel(kernel string) ILinuxBuilder {
	b.kernel = kernel
	return b
}

func (b *LinuxBuilder) AddWM(wm string) ILinuxBuilder {
	b.windowManager = wm
	return b
}

func (b *LinuxBuilder) AddDE(de string) ILinuxBuilder {
	b.desktopEnvironment = de
	return b
}

func (b *LinuxBuilder) AddPkgs(pkgs []string) ILinuxBuilder {
	copy(b.packages, pkgs)
	return b
}

func (b *LinuxBuilder) AddVersion(version string) ILinuxBuilder {
	b.version = version
	return b
}

func (b *LinuxBuilder) Build() *LinuxDistribution {
	return &LinuxDistribution{
		Name:               b.name,
		Kernel:             b.kernel,
		WindowManager:      b.windowManager,
		DesktopEnvironment: b.desktopEnvironment,
		Packages:           b.packages,
	}
}

// Director for builder
type LinuxDirector struct {
	b ILinuxBuilder
}

// Director's constructor
func NewLinuxDirector(b ILinuxBuilder) *LinuxDirector {
	return &LinuxDirector{
		b: b,
	}
}

// One of the director's methods
func (d *LinuxDirector) BuildUbuntu() *LinuxDistribution {
	return d.b.AddKernel("5.15.160").
		AddName("ubuntu").
		AddWM("Mutter").
		AddDE("GNOME").
		AddPkgs([]string{
			"nano", "grep", "awk", "snap",
		}).
		AddVersion("22.04").
		Build()
}

// One of the director's methods
func (d *LinuxDirector) BuildArch() *LinuxDistribution {
	return d.b.AddKernel("6.8").
		AddName("arch").
		AddWM("Hyprland").
		AddDE("").
		AddPkgs([]string{
			"nano", "grep", "awk", "pacman",
		}).
		AddVersion("unstable").
		Build()
}

// Example usage
//
// func main() {
// 	linuxBuilder := pattern.NewLinuxBuilder()
// 	fedora := linuxBuilder.
// 		AddName("fedora").
// 		AddKernel("6.1.72").
// 		AddWM("Mutter").
// 		AddDE("GNOME").
// 		AddPkgs([]string{
// 			"nano", "grep", "awk", "top", "ss", "flatpak",
// 		}).
// 		AddVersion("38").
// 		Build()

// 	fmt.Println(fedora.Name)

// 	director := pattern.NewLinuxDirector(linuxBuilder)

// 	ubuntu := director.BuildUbuntu()
// 	fmt.Println(ubuntu.Name)

// 	arch := director.BuildArch()
// 	fmt.Println(arch.Name)
// }
