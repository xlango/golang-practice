package main

import "C"

type P interface {
	Show(string) string
}

type S struct {
}

func (S) Show(th string) (t string) {
	if th == "b" {
		t = "b"
	} else {
		t = "a"
	}
	return
}

func main() {
	//var p  P=S{}
	//t:="b"
	//fmt.Println(p.Show(t))

}

type Integer int

func (a *Integer) Add(b *Integer) Integer {

	return *a + *b

}
