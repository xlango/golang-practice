package client

import (
	"log"
	"practice/crawler/engine"
	"practice/crawler_distributed/rpcsupport"
)

func ItemSave(host string) (chan engine.Item, error) {

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver : got item #%d : %v", itemCount, item)
			itemCount++

			//Rpc
			result := ""
			err = client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item saver %v error : %v", item, err)
			}
		}
	}()

	return out, nil
}
