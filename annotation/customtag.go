package main

import (
	"fmt"
	"reflect"
)

type InParam struct {
	StudentName string `json:"name" cc:"str,min=5,max=15"`
	Score       int    `json:"score" cc:"num,min=0,max=100"`
}

func main() {
	t := reflect.TypeOf(&InParam{"18", 25})
	field := t.Elem().Field(0)

	jsonName := field.Tag.Get("json")
	cc := field.Tag.Get("cc")

	fmt.Printf("%s(%s): %s", field.Name, jsonName, cc)
}
