package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func (ua *UserAges) GetLock(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {

	var wg sync.WaitGroup
	userAges := UserAges{ages: make(map[string]int)}
	wg.Add(100)

	go func() {
		userAges.Add("a", 99)
		wg.Done()
	}()

	for j := 0; j < 99; j++ {
		go func() {
			get := userAges.Get("a")
			fmt.Println(get)
			wg.Done()
		}()
	}

	wg.Wait()
}
