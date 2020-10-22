/*server.go */

package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var unixAddr *net.UnixAddr

	unixAddr, _ = net.ResolveUnixAddr("unix", "/tmp/unix_sock")

	unixListener, _ := net.ListenUnix("unix", unixAddr)

	defer unixListener.Close()

	for {
		unixConn, err := unixListener.AcceptUnix()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + unixConn.RemoteAddr().String())
		go unixPipe(unixConn)
	}

}

func unixPipe(conn *net.UnixConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
