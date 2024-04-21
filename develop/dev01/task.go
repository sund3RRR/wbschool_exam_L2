package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

type NtpServer struct {
	host string
}

func NewNtpServer(host string) *NtpServer {
	return &NtpServer{
		host: host,
	}
}

func (s *NtpServer) GetTime() (time.Time, error) {
	return ntp.Time(s.host)
}

func main() {
	server := NewNtpServer("0.beevik-ntp.pool.ntp.org")

	time, err := server.GetTime()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(time.Format("2006-01-02 15:04:05 -07:00:00"))
}
