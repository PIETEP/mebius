package service

import (
	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/gin-gonic/gin"
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

func (ds DeckService) CreateModel(c *gin.Context) (*entity.Deck, error) {
	db := db.GetDB()
	var deck entity.Deck

	if err := c.BindJSON(&deck); err != nil {
		return nil, err
	}

	if err := db.Create(&deck).Error; err != nil {
		return nil, err
	}

	return &deck, nil
}

func (ds DeckService) GetByID(id string) (*entity.Deck, error) {
	db := db.GetDB()
	var deck entity.Deck

	if err := db.Where("id = ?", id).First(&deck).Error; err != nil {
		return nil, err
	}

	return &deck, nil
}

func (ds DeckService) UpdateByID(id string, c *gin.Context) (*entity.Deck, error) {
	db := db.GetDB()
	var deck entity.Deck

	if err := db.Where("id = ?", id).First(&deck).Error; err != nil {
		return nil, err
	}

	if err := c.BindJSON(&deck); err != nil {
		return nil, err
	}

	db.Save(&deck)

	return &deck, nil
}

func (ds DeckService) DeleteByID(id string) error {
	db := db.GetDB()
	var deck entity.Deck

	if err := db.Where("id = ?", id).First(&deck).Delete(&deck).Error; err != nil {
		return err
	}

	return nil
}
