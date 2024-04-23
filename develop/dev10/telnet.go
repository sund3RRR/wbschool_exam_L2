package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type TelnetClient struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTelnetClient(cmdFlags *CmdFlags) (*TelnetClient, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(cmdFlags.host, cmdFlags.port), cmdFlags.timeout)
	if err != nil {
		return nil, err
	}

	client := &TelnetClient{
		conn:    conn,
		timeout: cmdFlags.timeout,
	}

	return client, nil
}

// Start function runs telnet client
func (tc *TelnetClient) Start() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go tc.readFromSocket()
	go tc.readFromStdin()

	<-sigCh
	fmt.Println("Ctrl+D pressed. Closing connection.")
	tc.conn.Close()
}

func (tc *TelnetClient) readFromSocket() {
	buf := make([]byte, 1024)
	for {
		n, err := tc.conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed by remote server.")
			os.Exit(0)
		}
		os.Stdout.Write(buf[:n])
	}
}

func (tc *TelnetClient) readFromStdin() {
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Error reading from STDIN:", err)
			break
		}
		_, err = tc.conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing to socket:", err)
			break
		}
	}
}
