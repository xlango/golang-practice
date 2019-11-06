package client

import (
	"net/rpc"
	"practice/crawler/engine"
	"practice/crawler_distributed/config"
	"practice/crawler_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(
	//	fmt.Sprintf(":%d", config.WorkerServerPort0),
	//)
	//if err != nil {
	//	return nil, err
	//}

	return func(request engine.Request) (engine.ParserResult, error) {
		serializeRequest := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.WorkerRpc, serializeRequest, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
