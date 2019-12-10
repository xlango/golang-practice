package main

import "fmt"

type A struct {
	A1 int
}

type B struct {
	B1 int
	A
}

type C func(int, int) int

func FC(c C) C {
	return func(i int, i2 int) int {
		return i + i2
	}
}

type People interface {
	Show()
}

type Student struct {
}

func (s *Student) Show() {
}

func live() People {
	var s *Student
	fmt.Println(s)
	return s
}

func live1() *Student {
	var s *Student
	fmt.Println(s)
	return s
}

func main() {

	p := live()

	if p == nil {
		fmt.Println("AAAAAAAAAAA")
	} else {
		fmt.Println("BBBBBBBB")
	}
}
