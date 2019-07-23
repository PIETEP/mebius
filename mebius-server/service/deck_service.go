package service

import (
	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/entity"
)

type DeckService struct{}

func (ds DeckService) GetAll() ([]entity.Deck, error) {
	db := db.GetDB()
	var d []entity.Deck

	if err := db.Find(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}
