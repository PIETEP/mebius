package controller

import (
	"github.com/PIETEP/mebius/mebius-server/service"
	"github.com/gin-gonic/gin"
)

type SuggestController struct{}

type SuggestIndexResponse struct {
	Words []string `json:"words"`
}

func (sc SuggestController) Index(c *gin.Context) {
	var s service.SuggestService
	name := c.Query("q")

	names, err := s.SuggestCardNames(name)
	if err != nil {
		c.AbortWithStatus(400)
	}

	r := SuggestIndexResponse{
		Words: names,
	}
	c.JSON(200, r)
}
