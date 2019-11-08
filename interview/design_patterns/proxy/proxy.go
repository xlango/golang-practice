package proxy

import "fmt"

//golang实现代理模式很简单，因为函数是golang的一等公民，能作为参数传递
type A struct {
	Name string
}

type Func func()

func (a *A) GetName() string {
	return a.Name
}

func F(a *A) {
	Proxy(func() {
		fmt.Println(a.GetName())
	})

}

func Proxy(f Func) {
	fmt.Println("start transaction")
	f()
	fmt.Println("end transaction")
}
