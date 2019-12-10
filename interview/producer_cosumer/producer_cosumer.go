package main

import (
	"fmt"
	"time"
)

//无缓冲：同步通信
//有缓存：异步通信

func producer(out chan<- int) {
	for i := 0; i < 20; i++ {
		out <- i
		fmt.Printf("生产：%d\n", i)
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("消费：%d\n", num)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	//ch:=make(chan int,5)
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
