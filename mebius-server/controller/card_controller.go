package controller

import (
	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/gin-gonic/gin"
)

type CardController struct{}

// This is for mock.
var card0 entity.Card
var card1 entity.Card

type UserIndexResponse struct {
	Cards []entity.Card `json:"cards"`
}

func init() {
	// This is for mock.
	card0 = entity.Card{
		ID:       0,
		Name:     "ドラゴンナイト",
		Point:    4,
		Species0: "Dragon",
		Species1: "Human",
		Job:      "Knight",
	}

	// This is for mock.
	card1 = entity.Card{
		ID:       1,
		Name:     "デストロイヤー",
		Point:    5,
		Species0: "Goblin",
		Species1: "",
		Job:      "Mechanic",
	}
}

// TODO: Implement the contents
func (cc CardController) Index(c *gin.Context) {
	r := UserIndexResponse{}
	r.Cards = []entity.Card{card0, card1}

	c.JSON(200, r)
}

// TODO: Implement the contents
func (cc CardController) Create(c *gin.Context) {
	c.JSON(200, card0)
}

// TODO: Implement the contents
func (cc CardController) Show(c *gin.Context) {
	c.JSON(200, card0)
}

// TODO: Implement the contents
func (cc CardController) Update(c *gin.Context) {
	c.JSON(200, card0)
}

// TODO: Implement the contents
func (cc CardController) Delete(c *gin.Context) {
	c.JSON(200, card0)
}
