package main

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"practice/crawler_distributed/repository"
	"practice/crawler_distributed/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(":1234", "crawler_dist"))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.222:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	rpcsupport.ServeRpc(host,
		&repository.ItemSaverService{
			Client: client,
			Index:  index,
			//Index:"crawler_dist",
		})
	return nil
}
