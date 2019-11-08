package main

import "practice/interview/design_patterns/proxy"

func main() {
	//f:=new(factory.Factory)
	//a := f.Create("A")
	//fmt.Println(a.GetName())
	//b := f.Create("B")
	//fmt.Println(b.GetName())

	//http.HandleFunc("/hello", decorator.AutoAuth(decorator.Hello))
	//err := http.ListenAndServe(":5666", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	proxy.F(&proxy.A{"11111111"})
}
