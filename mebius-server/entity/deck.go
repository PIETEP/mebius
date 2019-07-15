package entity

type Deck struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Cards   []Card
	Creator string
}
