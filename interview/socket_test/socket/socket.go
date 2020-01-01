package socket

import (
	"fmt"
	"net"
	"os"
)

type Socket struct {
	net.TCPConn
}

func (s *Socket) NewSocket() {

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
