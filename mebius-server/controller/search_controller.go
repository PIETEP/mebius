package controller

import "github.com/gin-gonic/gin"

type SearchController struct{}

type Deck struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DeckScore struct {
	Point uint `json:"point"`
	Deck  Deck `json:"deck"`
}

type SearchIndexResponse struct {
	Scores []DeckScore `json:"scores"`
}

// Index action: GET /users
func (sc SearchController) Index(c *gin.Context) {
	s0 := DeckScore{
		Point: 100,
		Deck: Deck{
			ID:   0,
			Name: "",
			URL:  "http://<domain>/deck/0",
		},
	}
	s1 := DeckScore{
		Point: 90,
		Deck: Deck{
			ID:   1,
			Name: "",
			URL:  "http://<domain>/deck/1",
		},
	}
	s2 := DeckScore{
		Point: 80,
		Deck: Deck{
			ID:   2,
			Name: "",
			URL:  "http://<domain>/deck/2",
		},
	}

	ss := []DeckScore{s0, s1, s2}

	r := SearchIndexResponse{
		Scores: ss,
	}

	c.JSON(200, r)
}
