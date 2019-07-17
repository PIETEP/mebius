package service

import (
	"github.com/PIETEP/mebius/mebius-server/db"
	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/gin-gonic/gin"
)

type CardService struct{}

type Card entity.Card

func (cs CardService) GetAll() ([]Card, error) {
	db := db.GetDB()
	var c []Card

	if err := db.Find(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cs CardService) CreateModel(c *gin.Context) (*Card, error) {
	db := db.GetDB()
	var card Card

	if err := c.BindJSON(&card); err != nil {
		return nil, err
	}

	if err := db.Create(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (cs CardService) GetByID(id string) (*Card, error) {
	db := db.GetDB()
	var card Card

	if err := db.Where("id = ?", id).First(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}
