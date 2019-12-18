package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	defer conn.Close()

	//监听超时强踢退出
	isTimeout := make(chan bool)

	go func() {
		for {
			buf := make([]byte, 128)
			n, err := conn.Read(buf)

			if err != nil {
				isTimeout <- true
				fmt.Println("conn.Read err:", err)
				return
			}

			if n > 0 {
				fmt.Println(string(buf[0:n]))
			}

		}
	}()

	var msg = ""

	for {
		select {
		case <-isTimeout:
			return
		default:
			fmt.Scan(&msg)
			conn.Write([]byte(msg))
		}

	}

}
