package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Паттерн "Команда" (Command) является поведенческим паттерном проектирования,
// который используется для инкапсуляции запроса как объекта, позволяя таким образом параметризовать
// клиентов с запросами, организовывать очереди или регистрировать запросы, а также поддерживать отмену операций.
//
// Плюсы:
// - Изолирует отправителя и получателя. Паттерн позволяет разделить код, который
// инициирует операцию, от кода, который её выполняет.
// - Поддерживает отмену и повторение операций. Легко добавлять операции отмены,
// так как каждая команда может содержать метод отмены.
// - Позволяет построить сложные системы. Комбинируя команды, можно создавать более сложные системы,
// такие как менеджеры транзакций или регистраторы команд.
//
// Минусы:
// - Может увеличить количество классов. Использование паттерна "Команда" может привести к созданию большого
// количества классов, что может усложнить структуру программы.
// - Дополнительные накладные расходы на память. Каждый объект команды может содержать свои собственные данные,
// что может привести к дополнительным затратам на память.

// Main command interface
type Command interface {
	execute()
}

// Sender implementation (in that case - console like terminal emulator)
type Console struct {
	CommandVar Command
}

func (c *Console) Run() {
	c.CommandVar.execute()
}

// Command implementation
type ListCommand struct {
	Files []string
}

func (c *ListCommand) execute() {
	for _, str := range c.Files {
		fmt.Println(str)
	}
}

// Another command implementation
type NeofetchCommand struct {
	Distro   string
	AsciiArt string
}

func (c *NeofetchCommand) execute() {
	fmt.Println(c.Distro)
	fmt.Print(c.AsciiArt + "\n")
}

// Example usage
//
// func main() {
// 	nc := pattern.NeofetchCommand{
// 		Distro: "Ubuntu",
// 		AsciiArt: `
// 		 _nnnn_
//         dGGGGMMb     ,"""""""""""""".
//        @p~qp~~qMb    |  I'm Ubuntu! |
//        M|@||@) M|   _;..............'
//        @,----.JM| -'
//       JS^\__/  qKL
//      dZP        qKRb
//     dZP          qKKb
//    fZP            SMMb
//    HZM            MMMM
//    FqM            MMMM
//  __| ".        |\dS"qML
//  |    '.       | '' \Zq
// _)      \.___.,|     .'
// \____   )MMMMMM|   .'
//      '-'       '--' hjm`,
// 	}

// 	console := pattern.Console{
// 		CommandVar: &nc,
// 	}
// 	console.Run()

// 	ls := pattern.ListCommand{
// 		Files: []string{
// 			"Documents",
// 			"Desktop",
// 			"Downloads",
// 			"Pictures",
// 			"Movies",
// 			".config",
// 			".local",
// 			".bashrc",
// 		},
// 	}

// 	console.CommandVar = &ls
// 	console.Run()
// }
