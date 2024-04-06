package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Паттерн "Состояние" (State) - это поведенческий паттерн проектирования, который позволяет объекту изменять
// свое поведение в зависимости от внутреннего состояния. Он позволяет объекту изменять свое поведение при изменении
// его внутреннего состояния. Это достигается путем выделения различных состояний в отдельные классы и делегирования
// поведения соответствующему классу состояния.
//
// Плюсы:
//
// - Уменьшение сложности кода за счет разделения состояний на отдельные классы.
// - Упрощение добавления новых состояний и изменения поведения объекта без изменения существующего кода.
// - Увеличение уровня повторного использования кода.
//
// Минусы:
//
// - Увеличение количества классов в системе.
// - Усложнение отслеживания состояний и переходов между ними.
// - Неудобство при работе с объектами, которые имеют большое количество возможных состояний и переходов между ними.

// Base state interface
type State interface {
	handle(ship *SpaceShip)
}

// Spacship has two states and field with current state
type SpaceShip struct {
	shootState  *ShootState
	reloadState *ReloadState
	state       State
}

func NewSpaceShip(ammo int) *SpaceShip {
	ship := SpaceShip{
		shootState: &ShootState{
			ammo: ammo,
		},
		reloadState: &ReloadState{
			ammo: ammo,
		},
	}
	ship.SetState(ship.shootState)

	return &ship
}
func (ship *SpaceShip) SetState(state State) {
	ship.state = state
}

func (ship *SpaceShip) Do() {
	ship.state.handle(ship)
}

type ShootState struct {
	ammo int
}

// Make a single shot and decrease ammo by 1
// Reload if ammo == 0
func (s *ShootState) handle(ship *SpaceShip) {
	fmt.Println("Shoot")
	s.ammo -= 1
	if s.ammo == 0 {
		ship.SetState(ship.reloadState)
	}
}

type ReloadState struct {
	ammo int
}

// Reload gun with ammo
func (s *ReloadState) handle(ship *SpaceShip) {
	fmt.Println("Reloading")
	ship.shootState.ammo = s.ammo
	ship.SetState(ship.shootState)
}

// Example usage
//
// func main() {
// 	ship := pattern.NewSpaceShip(10)

// 	for i := 0; i < 70; i++ {
// 		ship.Do()
// 	}
// }
