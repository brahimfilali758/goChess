package chess

import (
	"fmt"

)


type Player struct {
	Color Color
	username string
	rating int
}

func NewPlayer(color Color, username string) *Player {
	return &Player{
		Color: color,
		username: username,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("%s (%s)", p.username, p.Color)
}
