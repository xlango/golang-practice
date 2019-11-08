package main

import (
	"log"
	"net/http"
	"practice/interview/design_patterns/decorator"
)

func main() {
	//f:=new(factory.Factory)
	//a := f.Create("A")
	//fmt.Println(a.GetName())
	//b := f.Create("B")
	//fmt.Println(b.GetName())

	http.HandleFunc("/hello", decorator.AutoAuth(decorator.Hello))
	err := http.ListenAndServe(":5666", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
