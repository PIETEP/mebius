package service

import (
	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/gin-gonic/gin"
)

type CardService struct{}

func (cs CardService) GetAll() ([]entity.Card, error) {
	db := db.GetDB()
	var c []entity.Card

	if err := db.Find(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cs CardService) CreateModel(c *gin.Context) (*entity.Card, error) {
	db := db.GetDB()
	var card entity.Card

	if err := c.BindJSON(&card); err != nil {
		return nil, err
	}

	if err := db.Create(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (cs CardService) GetByID(id string) (*entity.Card, error) {
	db := db.GetDB()
	var card entity.Card

	if err := db.Where("id = ?", id).First(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (cs CardService) UpdateByID(id string, c *gin.Context) (*entity.Card, error) {
	db := db.GetDB()
	var card entity.Card

	if err := db.Where("id = ?", id).First(&card).Error; err != nil {
		return nil, err
	}

	if err := c.BindJSON(&card); err != nil {
		return nil, err
	}

	db.Save(&card)

	return &card, nil
}

func (cs CardService) DeleteByID(id string) error {
	db := db.GetDB()
	var card entity.Card

	if err := db.Where("id = ?", id).First(&card).Delete(&card).Error; err != nil {
		return err
	}

	return nil
}
