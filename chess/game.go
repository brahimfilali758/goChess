package chess

import (
	"github.com/google/uuid"
)


type Game struct {
	uuid string
	board *Board
	p1 *Player
	p2 *Player
}

// create a new game
func NewGame(p1 *Player, p2 *Player) *Game {
	fen_str := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	return &Game{
		uuid: uuid.New().String(),
		board: GenerateBoardFromFen(fen_str),
		p1: p1,
		p2: p2,
	}
}

func (g *Game) GetBorad() *Board {
	return g.board
}