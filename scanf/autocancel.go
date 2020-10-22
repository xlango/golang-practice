package main

import (
	"context"
	"fmt"
	"time"
)

// 3 seconds for example
var deadline = time.Second * 10

func main() {
	c := make(chan string, 1)
	go scan(c)

	ctx, _ := context.WithTimeout(context.Background(), deadline)

	select {
	case <-ctx.Done():
		// didnt type for deadline seconds

	case <-c:
		// did it in time

	}
}

func scan(in chan string) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	in <- input
}
