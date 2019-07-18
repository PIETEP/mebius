package controller

import (
	"fmt"

	"github.com/PIETEP/mebius/mebius-server/entity"
	"github.com/PIETEP/mebius/mebius-server/service"
	"github.com/gin-gonic/gin"
)

type CardController struct{}

type UserIndexResponse struct {
	Cards []entity.Card `json:"cards"`
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

	r := UserIndexResponse{
		Cards: cards,
	}

	c.JSON(200, r)
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
