package engine

import (
	"log"
	"practice/set"
)

var isDupSet *set.Set

func init() {
	isDupSet = set.New()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParserResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount:=0

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item : %v \n", item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func creatWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	//in :=make(chan  Request)
	go func() {
		for {
			//tell scheduler i`m ready
			ready.WorkerReady(in)
			request := <-in
			result, err := work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func isDuplicate(url string) bool {
	if isDupSet.IsExist(url) {
		return true
	}
	return false
}
