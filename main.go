package main 

import (
	"fmt"
	"github.com/brahimfilali758/goChess"
)

func main() {
	p1 := goChess.NewPlayer(chess.White, "Player 1")
	p2 := goChess.NewPlayer(chess.Black, "Player 2")
	game := goChess.NewGame(p1, p2)
	game.board.PrintBoard()
}