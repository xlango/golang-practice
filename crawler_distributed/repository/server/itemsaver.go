package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"practice/crawler_distributed/config"
	"practice/crawler_distributed/repository"
	"practice/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.EsIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(config.EsUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	rpcsupport.ServeRpc(host,
		&repository.ItemSaverService{
			Client: client,
			Index:  index,
		})
	return nil
}
