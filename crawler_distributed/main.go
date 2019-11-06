package main

import (
	"flag"
	"log"
	"net/rpc"
	"practice/crawler/engine"
	"practice/crawler/idoctor/parser"
	"practice/crawler/scheduler"
	"practice/crawler_distributed/config"
	itemsaver "practice/crawler_distributed/repository/client"
	"practice/crawler_distributed/rpcsupport"
	worker "practice/crawler_distributed/worker/client"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_host", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSave(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	h := strings.Split(*workerHosts, ",")

	pool := createClientPool(h)

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		//Url:        "http://www.zhenai.com/zhenghun",
		//Url:        "http://www.daxuecn.com/chaxun/",
		//Url:        "https://cd.julive.com/project/s",
		Url:    "https://ysk.99.com.cn/department/all/",
		Parser: engine.NewFuncParser(parser.ParserKeshiList, config.ParserKeshi),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf(
				"Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}

	}()

	return out
}
