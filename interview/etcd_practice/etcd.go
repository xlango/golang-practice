package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func conn() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.10.36:2379", "192.168.10.36:2380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}
	defer cli.Close()
	fmt.Println("conn success")

	//设置超时及设置值
	ctx, cancelFunc := context.WithTimeout(context.Background(), 50*time.Second)
	_, err = cli.Put(ctx, "name", "Jaye")
	cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
	}

	//取值
	ctx, cancelFunc = context.WithTimeout(context.Background(), 50*time.Second)
	res, err := cli.Get(ctx, "name")
	if err != nil {
		log.Println("cli.Get", err.Error())
	}

	for k, v := range res.Kvs {
		fmt.Println("查询结果", k, string(v.Key), string(v.Value))
	}
}

func watch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.10.36:2379", "192.168.10.36:2380"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}
	defer cli.Close()
	fmt.Println("conn success")

	//设置超时及设置值
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = cli.Put(ctx, "home", "sdfwsdf ")
	cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
	}

	//watch
	for {
		rch := cli.Watch(context.Background(), "home")
		for resp := range rch {
			for k, v := range resp.Events {
				fmt.Println(k, v.Type, string(v.Kv.Key), string(v.Kv.Value))
			}
		}
	}
}

func main() {
	conn()
	watch()
}
