package entity

type Card struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Point    uint   `json:"point"`
	Species0 string `json:"species0"`
	Species1 string `json:"species1"`
	Job      string `json:"job"`
}
