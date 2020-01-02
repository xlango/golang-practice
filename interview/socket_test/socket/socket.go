package socket

import (
	"fmt"
	"net"
	"os"
)

type Socket struct {
	Conn net.Conn
}

func (s *Socket) NewSocketServer(listener *net.TCPListener) {
	//service := ":7777"
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//checkError(err)
	//listener, err := net.ListenTCP("tcp", tcpAddr)
	//checkError(err)

	conn, err := listener.Accept()
	if err != nil {
		checkError(err)
	}

	s.Conn = conn

}

func (s *Socket) NewSocketClient() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7777")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	s.Conn = conn
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
