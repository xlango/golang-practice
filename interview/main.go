package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"practice/interview/algorithm"
)

type A struct {
	a int
}

type B struct {
	b int
}

type C interface {
	c() int
}

type D struct {
	d string
	c C
}

type E struct {
	f F
}

func (e *E) c() int {
	return e.f()
}
func NewE(f F) *E {
	return &E{f: f}
}

type F func() int

func F1() int {
	fmt.Println("11111111111")
	return 1
}

func main() {
	//d:=D{
	//	d:"d",
	//	c:NewE(F1),
	//}
	//fmt.Println(d)

	//sum10000()
	//locktest()
	//time.Sleep(time.Second)

	//algorithm.Algo1(100)

	//sort
	arr:=[]int{2,4,1,5,7,8,3}
	//algorithm.BubbleSort([]int{2,4,1,5,7,8,3})
	//algorithm.SelectSort([]int{2,4,1,5,7,8,3})
	//algorithm.InsertSort([]int{2,4,1,5,7,8,3})
	//algorithm.Quick3Sort(arr,0,6)
	algorithm.QuickSort(arr,0, len(arr)-1)
	fmt.Println(arr)
}


func sum10000() {
	var wg sync.WaitGroup
	sumChan := make(chan int, 10)
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func(s, e, index int,w *sync.WaitGroup) {
			defer wg.Done()
			sum := 0
			for p := s; p <= e; p++ {
				sum += p
			}
			fmt.Printf("i :%d=%d \n", index, sum)
			sumChan <- sum
		}(i*1000+1, (i+1)*1000, i,&wg)
	}

	wg.Wait()


	close(sumChan)
	sum := 0

	for s := range sumChan {
		sum += s
	}

	fmt.Println(sum)

}

func c(sumChan chan int) {
	select {
	case s := <-sumChan:
		fmt.Println(s)
	default:
		fmt.Println("??")
		close(sumChan)
	}
}

func locktest() {
	var (
		m = map[string]int{
			"a": 1,
		}

		mutex      sync.Mutex
		mutexCount int64

		mutexR      sync.RWMutex
		mutexRCount int64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mutex.Lock()
				_ = m["a"]
				mutex.Unlock()
				atomic.AddInt64(&mutexCount, 1)
			}

		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mutexR.Lock()
				_ = m["a"]
				mutexR.Unlock()
				atomic.AddInt64(&mutexRCount, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("%d\n%d", mutexCount, mutexRCount)
}
