package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Основная идея фасада заключается в создании унифицированного интерфейса для набора интерфейсов в подсистеме.
// Фасад предоставляет более простой интерфейс для взаимодействия с подсистемой, скрывая её сложность и детали реализации.
// Это особенно полезно в системах с большим количеством компонентов или сложными взаимосвязями между ними.
//
// Плюсы:
// - Упрощение использования. Позволяет клиентскому коду взаимодействовать с системой через простой интерфейс,
// скрывая сложность и детали реализации.
// - Снижение зависимостей. Клиентский код зависит только от фасада, что уменьшает связанность системы и делает её
// более гибкой к изменениям внутренней реализации.
// - Улучшение безопасности. Фасад может контролировать доступ к компонентам подсистемы, что способствует более
// безопасному использованию системы.
//
// Минусы:
// - Ограниченность функциональности. Фасад может скрывать некоторые возможности подсистемы, что может ограничить
// функциональность, доступную клиентскому коду.
// - Усложнение поддержки. Если изменения в подсистеме требуют изменения в фасаде, это может повлечь за собой
// дополнительную работу по поддержке кода.
// - Добавление слоя абстракции. Использование фасада добавляет еще один уровень абстракции в систему, что может
// усложнить её понимание и отладку.

//
// OsFacade
//

type ProgramFacade struct {
	processService  *ProcessService
	threadService   *ThreadService
	resourceService *ResourceService
}

// Facade method that just shutdown process.
// We use methods of various interfaces, providing the client with a
// simplified interface, hiding the internal logic
func (of *ProgramFacade) ShutdownProcess() {
	pid := of.processService.pid
	of.processService.stopProcess()
	of.threadService.shutdownThread(pid)
	of.resourceService.freeResources(pid)
}

// Program facade constructor
func NewProgramFacade(pid int, threadCount int, memory int) *ProgramFacade {
	return &ProgramFacade{
		processService:  &ProcessService{pid: pid},
		threadService:   &ThreadService{count: threadCount},
		resourceService: &ResourceService{memory: memory},
	}
}

//
// OsFacade
//

//
// ProcessService
//

type ProcessService struct {
	pid int
}

func (ps *ProcessService) stopProcess() {
	fmt.Printf("Process PID=%d stopped\n", ps.pid)
}

//
// ProcessService
//

//
// ThreadService
//

type ThreadService struct {
	count int
}

func (ts *ThreadService) shutdownThread(pid int) {
	fmt.Printf("%d threads for process PID=%d shutted down\n", ts.count, pid)
}

//
// ThreadService
//

//
// ResourceService
//

type ResourceService struct {
	memory int
}

func (rs *ResourceService) freeResources(pid int) {
	fmt.Printf("%d MB of memory for process PID=%d freed\n", rs.memory, pid)
}

//
// ResourceService
//
