package controller

import "github.com/gin-gonic/gin"

type SuggestController struct{}

type SuggestIndexResponse struct {
	Words []string `json:"words"`
}

func (sc SuggestController) Index(c *gin.Context) {
	r := SuggestIndexResponse{
		Words: []string{"ドラゴンナイト", "ドラゴンなんちゃら", "ドラゴン強い子"},
	}
	c.JSON(200, r)
}
