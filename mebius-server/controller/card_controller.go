package controller

import (
	"fmt"

	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/PIETEP/mebius/mebius-server/service"
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
	var s service.CardService

	cards, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(200, cards)
}

// TODO: Implement the contents
func (cc CardController) Create(c *gin.Context) {
	var s service.CardService

	card, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
		return
	}

	c.JSON(201, card)
}

// TODO: Implement the contents
func (cc CardController) Show(c *gin.Context) {
	var s service.CardService

	id := c.Params.ByName("id")
	card, err := s.GetByID(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(200, card)
}

// TODO: Implement the contents
func (cc CardController) Update(c *gin.Context) {
	var s service.CardService

	id := c.Params.ByName("id")
	card, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
		return
	}

	c.JSON(200, card)
}

// TODO: Implement the contents
func (cc CardController) Delete(c *gin.Context) {
	var s service.CardService

	id := c.Params.ByName("id")
	err := s.DeleteByID(id)

	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
		return
	}

	c.JSON(204, gin.H{"id #" + id: "deleted"})
}
