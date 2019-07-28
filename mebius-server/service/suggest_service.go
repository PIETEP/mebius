package service

import (
	"fmt"

	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/entity"
)

type SuggestService struct{}

func (ss SuggestService) SuggestCardNames(name string) ([]string, error) {
	db := db.GetDB()
	var cards []entity.Card

	if err := db.Where("name LIKE ?", "%"+name+"%").Find(&cards).Error; err != nil {
		fmt.Println("test", err)
		return nil, err
	}

	names := []string{}
	for _, c := range cards {
		names = append(names, c.Name)
	}

	return names, nil
}
