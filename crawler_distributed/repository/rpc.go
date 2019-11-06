package repository

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"practice/crawler/engine"
	"practice/crawler/repository"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := repository.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "OK"
	} else {
		log.Printf("Error item %v : %v", item, err)
	}
	return err
}
