package main

import (
	"fmt"
	"regexp"
)

//枚举
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

const (
	text = `my email is qq132@qq.com@111.com 123@eewrer.cn` //([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)
)

func main() {
	//fmt.Println(b,kb,mb,gb,tb,pb)
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	//match:=re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}
