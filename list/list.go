package main

import (
	"fmt"
	"reflect"
)

type List struct {
	A string
	B int
}

func Re(v interface{})  {
	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)
	for i:=0;i<valueOf.NumField() ; i++ {
		fieldof := reflect.TypeOf(valueOf.Field(i))
		fieldvalue := reflect.ValueOf(valueOf.Field(i))
		fmt.Println(fieldof,fieldvalue)
	}
	fmt.Println(typeOf,valueOf)

}

func main() {
	l := List{
		A: "1",
		B: 0,
	}
	Re(l)
}
