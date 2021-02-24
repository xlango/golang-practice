package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"
)

type engine struct {
	time string
	name string
}

//// 一个[]byte的对象池，每个对象为一个[]byte
//var pool = sync.Pool{
//	New: func() interface{} {
//		b := engine{}
//		return &b
//	},
//}
//
//func main() {
//
//	// 使用对象池
//	for i := 0; i < 10; i++{
//		obj := pool.Get().(*engine)
//		a := fmt.Sprintf("%+v",*obj)
//		fmt.Println(a)
//		_ = obj
//		pool.Put(obj)
//	}
//}

type reusableobj struct{}

type objpool struct {
	bufchan chan *engine
}

func newobjpool(numofobj int) *objpool {
	objpool := objpool{}
	objpool.bufchan = make(chan *engine, numofobj)
	for i := 0; i < numofobj; i++ {
		objpool.bufchan <- &engine{
			time: "",
			name: fmt.Sprintf("n%d",i),
		}
	}
	return &objpool
}

func (p *objpool) getobj(timeout time.Duration) (*engine, error) {
	select {
	case ret := <-p.bufchan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *objpool) releaseobj(obj *engine) error {
	select {
	case p.bufchan <- obj:
		return nil
	// 放不进去了就会返回overflow
	default:
		return errors.New("overflow")
	}
}

func main() {
	pool := newobjpool(10)
	// if err := pool.releaseobj(&reusableobj{}); err != nil {
	//  fmt.Println(err)
	// }
	for i := 0; i < 20; i++ {
		if v, err := pool.getobj(time.Second); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
			if err := pool.releaseobj(v); err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("done.")
}

func WithTimeout(dur time.Duration, fn func() error) error {
	var ch = make(chan error, 1)
	go time.AfterFunc(dur, func() {
		ch <- errors.New("函数执行超时: " + pkg.Caller.FromFunc(reflect.ValueOf(fn)))
	})
	go func() {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case error:
					ch <- err.(error)
				default:
					log.Fatalln(err)
				}
			}
		}()
		ch <- fn()
	}()

	return <-ch
}