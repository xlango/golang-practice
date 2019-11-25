package main

import "fmt"

/**
多个defer的执行顺序为“后进先出”；
defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出
*/

func main() {
	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
		fmt.Println("一")
	}()
	fmt.Println(":", t)
	return t
}
func DeferFunc2(i int) int { //未声明返回值，return时将自动声明一个返回值变量，return对其进行赋值，最后函数携带当前返回值退出
	t := i
	defer func() {
		t += 3
		fmt.Println("二")
	}()
	fmt.Println(":", t)
	return t
}
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
		fmt.Println("三")
	}()
	fmt.Println(":", t)
	return 2
}
