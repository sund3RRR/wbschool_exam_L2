package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func shellEventLoop(shell *Shell) {
	scanner := bufio.NewScanner(shell.in)

	for {
		dir := shell.SerializePath(shell.currentDir)

		fmt.Fprintf(shell.out, "btrsh %s %% ", dir)
		scanner.Scan()
		command := scanner.Text()

		output := shell.HandleCommand(command)

		fmt.Fprintf(shell.out, "%s", output)
	}
}

func main() {
	shell, err := NewShell()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	shellEventLoop(shell)
}
