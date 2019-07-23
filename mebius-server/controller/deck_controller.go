package controller

import (
	"fmt"

	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/PIETEP/mebius/mebius-server/service"
	"github.com/gin-gonic/gin"
)

type DeckController struct{}

type DeckIndexResponse struct {
	Decks []entity.Deck `json:"decks"`
}

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

func (dc DeckController) Update(c *gin.Context) {
	var s service.DeckService

	id := c.Params.ByName("id")
	deck, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
		return
	}

	c.JSON(200, deck)
}

func (dc DeckController) Delete(c *gin.Context) {
	var s service.DeckService

	id := c.Params.ByName("id")
	err := s.DeleteByID(id)

	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
		return
	}

	c.JSON(204, gin.H{"id #" + id: "deleted"})
}
