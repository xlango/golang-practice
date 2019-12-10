package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"time"
)

func main() {

	hystrix.ConfigureCommand("get_baidu", hystrix.CommandConfig{
		500,
		100,
		50,
		3,
		1000,
	})
	for i := 0; i < 100; i++ {

		TestHystix()
		time.Sleep(1 * time.Second)

	}
	time.Sleep(2 * time.Second) // 调用Go方法就是起了一个goroutine，这里要sleep一下，不然看不到效果
}

func TestHystix() {

	// 根据自身业务需求封装到http client调用处
	hystrix.Go("get_baidu", func() error {

		// 调用关联服务
		res, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Println("get error")
			return err
		}

		fmt.Println("请求成功：", res.Status)
		return nil
	},
		// 失败重试，降级等具体操作
		func(err error) error {
			fmt.Println("get an error, handle it")
			return nil
		})
}
