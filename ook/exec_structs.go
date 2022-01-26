package ook

import (
	"bufio"
	"net"
)

type Exec struct {
	Command string
	Args    []string
	Stdout  *bufio.Reader
	Stderr  *bufio.Reader
	Time    string
	Conn    net.Conn
	Stream  ExecStream
}

type ExecStream struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Hash string `json:"hash"`
}
