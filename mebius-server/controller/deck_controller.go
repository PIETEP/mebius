package controller

import (
	"github.com/PIETEP/mebius/mebius-server/entity"
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
	r := DeckIndexResponse{
		Decks: []entity.Deck{deck0, deck1},
	}

	c.JSON(200, r)
}

// TODO: Implement the contents
func (dc DeckController) Create(c *gin.Context) {
	c.JSON(200, deck0)
}

// TODO: Implement the contents
func (dc DeckController) Show(c *gin.Context) {
	c.JSON(200, deck0)
}

// TODO: Implement the contents
func (dc DeckController) Update(c *gin.Context) {
	c.JSON(200, deck0)
}

// TODO: Implement the contents
func (dc DeckController) Delete(c *gin.Context) {
	c.JSON(200, deck0)
}
