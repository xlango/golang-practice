package main

import (
	"practice/crawler/engine"
	"practice/crawler/idoctor/parser"
	"practice/crawler/scheduler"
	"practice/crawler_distributed/repository/client"
)

func main() {
	itemChan, err := client.ItemSave(":1234")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		//Url:        "http://www.zhenai.com/zhenghun",
		//Url:        "http://www.daxuecn.com/chaxun/",
		//Url:        "https://cd.julive.com/project/s",
		Url:        "https://ysk.99.com.cn/department/all/",
		ParserFunc: parser.ParserKeshiList,
	})
}
