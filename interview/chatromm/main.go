package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listener.Close()

	go Manager()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}

		//启动go程处理客户端数据请求
		go HandlerConn(conn)
	}
}

func HandlerConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()

	clnt := Client{make(chan string), addr, addr}

	onlineMap[addr] = clnt

	//创建给当前用户发送消息的go程
	go WriteToClient(clnt, conn)

	//返送用户上线消息
	//message <- fmt.Sprintf("[ %v ] %v : login", addr, clnt.Name)
	message <- MakeMsg(clnt, "login")

	//创建一个channel，用于判断用户退出
	isExit := make(chan bool)

	//创建一个channel，用于判断用户是否活跃
	isActive := make(chan bool)

	//go程处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isExit <- true
				fmt.Printf("检测到客户端 : %s 退出\n", clnt.Name)
				return
			}

			if err != nil {
				isExit <- true
				fmt.Println("conn.Read err:", err)
				return
			}

			msg := string(buf[:n])

			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				for _, user := range onlineMap {
					userInfo := fmt.Sprintf("%v : %v \n", user.Addr, user.Name)
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				clnt.Name = msg[7:]
				onlineMap[clnt.Addr] = clnt
				conn.Write([]byte("rename success! \n"))
			} else {
				//将读到的用户消息，写入到message中
				message <- MakeMsg(clnt, msg)
			}

			//用户发送信息则表示活跃
			isActive <- true
		}
	}()

	//客户端退出
	for {
		select {
		case <-isExit:
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		case <-isActive:
			//不动作，至少让下面计时器重新计时
		case <-time.After(time.Second * 10):
			delete(onlineMap, clnt.Addr)
			message <- MakeMsg(clnt, "logout")
			return
		}
	}

}

func MakeMsg(client Client, msg string) string {
	return fmt.Sprintf("[ %v ] %v : %v", client.Addr, client.Name, msg)
}

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func Manager() {
	onlineMap = make(map[string]Client)

	for {
		msg := <-message

		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}

}

func WriteToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}
