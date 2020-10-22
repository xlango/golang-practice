package main

import (
	"golang.org/x/exp/errors/fmt"
)

//priority channel test
func main() {
	c1 := make(chan string, 10)
	c2 := make(chan string, 10)
	c3 := make(chan int)

	go func() {
		for i := 0; i < 2000; i++ {
			c1 <- fmt.Sprintf("a%d", i)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			c2 <- fmt.Sprintf("b%d", i)
		}
	}()
	go func() {
		c3 <- 1111
	}()

	for {
		select {
		case a := <-c1:
			fmt.Println(a)
			//time.Sleep(1 * time.Second)
		default:
			select {
			case b := <-c2:
				fmt.Println(b)
			default:
				select {
				case c := <-c3:
					fmt.Println(c, 11111111111)
					break
				default:
				}
			}
		}
	}

}
