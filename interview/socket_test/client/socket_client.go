package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:7777")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	go func() {
		for {
			//result, err := ioutil.ReadAll(conn)
			b := make([]byte, 1024)
			n, _ := conn.Read(b)
			a := make([]byte, n)
			a = b[:n]
			fmt.Println(string(a))
		}
	}()

	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println("你输入的是：", input.Text())
		_, err = conn.Write([]byte(input.Text()))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
