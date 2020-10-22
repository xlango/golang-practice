package main


import (
"fmt"
"net"
"strconv"
"strings"
"time"
)

func smain() {
	lis, err := net.Listen("tcp", ":1789")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()
	for {
		con, _ := lis.Accept()
		go handler(con)
	}
}

func handler(con net.Conn) {
	defer con.Close()
	buf := make([]byte, 20)
	n, _ := con.Read(buf)
	fmt.Println(string(buf[:n]))
	con.Write(buf)
	con.Write(buf)
}

func main() {
	buf, err := Telnet([]string{"w_Hello World", "r_50", "r_30"}, "10.34.4.16:8000", 20)
	fmt.Println(err)
	fmt.Println(string(buf))
}

func PortIsOpen(ip string, timeout int) bool {
	con, err := net.DialTimeout("tcp", ip, time.Duration(timeout))
	if err != nil {
		return false
	}
	con.Close()
	return true
}

func Telnet(action []string, ip string, timeout int) (buf []byte, err error) {
	con, err := net.DialTimeout("tcp", ip, time.Duration(timeout)*time.Second)
	if err != nil {
		return
	}
	defer con.Close()
	con.SetReadDeadline(time.Now().Add(time.Second * 5))
	for _, v := range action {
		l := strings.SplitN(v, "_", 2)
		if len(l) < 2 {
			return
		}
		switch l[0] {
		case "r":
			var n int
			n, err = strconv.Atoi(l[1])
			if err != nil {
				return
			}
			p := make([]byte, n)
			n, err = con.Read(p)
			if err != nil {
				return
			}
			buf = append(buf, p[:n]...)
			fmt.Println(buf)
		case "w":
			_, err = con.Write([]byte(l[1]))
		}
	}
	return
}
