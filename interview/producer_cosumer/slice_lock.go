package main

import (
	"fmt"
	"sync"
	"time"
)

type Topic struct {
	S []interface{}
	sync.Mutex
}

func NewTopic() *Topic {
	return &Topic{
		S: make([]interface{}, 0),
	}
}

func (t *Topic) Producer(v interface{}) {
	t.Lock()
	defer t.Unlock()
	t.S = append(t.S, v)
}

func (t *Topic) Consumer() (v *interface{}) {
	t.Lock()
	defer t.Unlock()
	time.Sleep(1 * time.Second)
	v = &t.S[0]
	t.S = t.S[1:len(t.S)]
	return
}

func main() {
	t := NewTopic()
	go func() {
		for i := 0; i < 10; i++ {
			t.Producer(i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v := t.Consumer()
			fmt.Println(*v)
		}
	}()

	time.Sleep(11 * time.Second)
}
