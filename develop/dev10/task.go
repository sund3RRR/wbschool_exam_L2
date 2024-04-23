package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type CmdFlags struct {
	host    string
	port    string
	timeout time.Duration
}

func NewCmdFlags() *CmdFlags {
	var timeout time.Duration

	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := args[0]
	port := args[1]

	return &CmdFlags{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func main() {
	cmdFlags := NewCmdFlags()

	client, err := NewTelnetClient(cmdFlags)
	if err != nil {
		fmt.Printf("Error connecting to %s:%s: %v\n", cmdFlags.host, cmdFlags.port, err)
		os.Exit(1)
	}

	client.Start()
}
