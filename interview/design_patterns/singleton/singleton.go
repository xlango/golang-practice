package singleton

import (
	"sync"
	"sync/atomic"
	//"sync/once"
)

type singleton struct {
}

var (
	instance    = &singleton{}
	l           sync.Mutex
	initialized uint32
)

func NewSingleton() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func NewSingleton1() *singleton {
	l.Lock() // lock
	defer l.Unlock()
	if instance == nil { // check
		instance = &singleton{}
	}
	return instance
}

func NewSingleton2() *singleton {
	if instance == nil { // check
		l.Lock() // lock
		defer l.Unlock()
		if instance == nil { // check
			instance = &singleton{}
		}
	}
	return instance
}

func NewSingleton3() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	l.Lock()
	defer l.Unlock()
	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

//func NewSingleton4() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}
