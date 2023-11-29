package main 

import (
	"fmt"
	"github.com/brahimfilali758/goChess/chess"
)

func main() {
	fmt.Println("Game Starts !")
	p1 := chess.NewPlayer(chess.White, "Player 1")
	p2 := chess.NewPlayer(chess.Black, "Player 2")
	game := chess.NewGame(p1, p2)
	board := game.GetBorad()
	board.PrintBoard()
}

