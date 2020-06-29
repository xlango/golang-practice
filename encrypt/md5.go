package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	str := "123456"
	//方法一
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str1)
	//方法二
	w := md5.New()
	io.WriteString(w, str)
	//将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))

	fmt.Println(md5str2)

}
