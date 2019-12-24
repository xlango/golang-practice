package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client"
	"time"
)

func main() {
	//链接etcd
	cfg := client.Config{
		Endpoints: []string{"http://192.168.10.36:32799/", "http://192.168.10.36:32801/", "http://192.168.10.36:32803/"},
		Transport: client.DefaultTransport,
		//set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, _ := client.New(cfg)
	keysApi := client.NewKeysAPI(c)

	//读取etcd
	res, err := keysApi.Get(context.Background(), "job", nil)
	fmt.Println("res====>", res)
	fmt.Println("err====>", err)
	//添加etcd
	res1, err := keysApi.Set(context.Background(), "work", "This is my work", nil)
	fmt.Println("res1===>", res1)
	fmt.Println("err===>", err)
	//读取etcd
	values, err := keysApi.Get(context.Background(), "work", nil)
	fmt.Println("values====>", values)
	fmt.Println("err====>", err)
	//更新etcd
	up, err := keysApi.Update(context.Background(), "job", "my job")
	fmt.Println("up====>", up)
	fmt.Println("err====>", err)
	//删除
	del, err := keysApi.Delete(context.Background(), "work", nil)
	fmt.Println("del====>", del)
	fmt.Println("err====>", err)

}
