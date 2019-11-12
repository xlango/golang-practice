package main

import (
	"fmt"
	"practice/interview/goroutine_pool/pool"
	"time"
)

func main() {
	task := pool.NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	p := pool.NewPool(4)

	go func() {
		for {
			p.EntryChan <- task
		}
	}()

	p.Run()
}
