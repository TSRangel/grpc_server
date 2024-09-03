package entities

import "github.com/google/uuid"

type Game struct {
	ID        string
	Name      string
	Platform string
	Price     float64
}

func NewGame(name, platform string, price float64) *Game {
	return &Game{
		ID: uuid.New().String(),
		Name: name,
		Platform: platform,
		Price: price,
	}
}
