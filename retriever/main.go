package main

import (
	"fmt"
	"xiaolin/retriever/mock"
	"xiaolin/retriever/real"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	//var r Retriver
	//r=mock.Retriver{"this is a fake imooc.com"}
	//r=&real.Retriver{
	//	UserAgent:"Mozilla/5.0",
	//	TimeOut:time.Minute,
	//}

	//断言
	//assertion(r)

	//fmt.Println(download(r))\

	//闭包
	//a:=closure()
	//for i:=0;i<10;i++ {
	//	fmt.Println(a(i))
	//}

	//斐波那契数列
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func assertion(r Retriver) {
	switch v := r.(type) {
	case mock.Retriver:
		fmt.Println(v.Content)
	case *real.Retriver:
		fmt.Println(v.UserAgent, v.TimeOut)
	}
}

func closure() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
