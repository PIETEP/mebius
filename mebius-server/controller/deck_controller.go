package controller

import (
	"fmt"

	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/PIETEP/mebius/mebius-server/service"
	"github.com/gin-gonic/gin"
)

type DeckController struct{}

// This is for mock.
var deck0 entity.Deck
var deck1 entity.Deck

type DeckIndexResponse struct {
	Decks []entity.Deck `json:"decks"`
}

func init() {
	// This is for mock.
	card0 := entity.Card{
		ID:       0,
		Name:     "ドラゴンナイト",
		Point:    4,
		Species0: "Dragon",
		Species1: "Human",
		Job:      "Knight",
	}

	// This is for mock.
	card1 := entity.Card{
		ID:       1,
		Name:     "デストロイヤー",
		Point:    5,
		Species0: "Goblin",
		Species1: "",
		Job:      "Mechanic",
	}

	// This is for mock.
	deck0 = entity.Deck{
		ID:      0,
		Name:    "Dragon",
		Cards:   []entity.Card{card0, card1},
		Creator: "Yuta",
	}

	// This is for mock.
	deck1 = entity.Deck{
		ID:      1,
		Name:    "Goblin",
		Cards:   []entity.Card{card0, card1},
		Creator: "Yuya",
	}
}

// TODO: Implement the contents
func (dc DeckController) Index(c *gin.Context) {
	var s service.DeckService

	decks, err := s.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	r := DeckIndexResponse{
		Decks: decks,
	}

	c.JSON(200, r)
}

// TODO: Implement the contents
func (dc DeckController) Create(c *gin.Context) {
	var s service.DeckService

	deck, err := s.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
		return
	}

	c.JSON(200, deck)
}

// TODO: Implement the contents
func (dc DeckController) Show(c *gin.Context) {
	var s service.DeckService

	id := c.Params.ByName("id")
	deck, err := s.GetByID(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}

	c.JSON(200, deck)
}

// TODO: Implement the contents
func (dc DeckController) Update(c *gin.Context) {
	c.JSON(200, deck0)
}

// TODO: Implement the contents
func (dc DeckController) Delete(c *gin.Context) {
	c.JSON(200, deck0)
}
