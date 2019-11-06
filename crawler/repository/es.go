package repository

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"practice/crawler/engine"
)

func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("Must supply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v \n",resp)
	return nil
}

func ItemSave(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.222:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Printf("[error]elastic client create failed \n")
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver : got item #%d : %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item saver %v error : %v", item, err)
			}
		}
	}()

	return out, nil
}
