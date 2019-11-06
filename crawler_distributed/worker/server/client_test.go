package main

import (
	"fmt"
	"practice/crawler_distributed/config"
	"practice/crawler_distributed/rpcsupport"
	"practice/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "https://ysk.99.com.cn/ys/202588.html",
		Parser: worker.SerializeParser{
			Name: config.ParserDoctorProfile,
			Args: worker.DoctorParser{
				UserName: "冯昌生 ",
				Id:       "202588",
				Url:      "https://ysk.99.com.cn/ys/202588.html",
			},
		},
	}

	var result worker.ParseResult
	err = client.Call(config.WorkerRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
