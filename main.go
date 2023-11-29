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
	fmt.Println("Board printed")
	i := 0
	var p string
	var startrank int 
	var startfile int
	var endrank int
	var endfile int
	m := make(map[string]string)
	m["p"] = "♟"
	m["P"]= "♙"
	m["r"] = "♜"
	m["R"] = "♖"
	m["n"] = "♞"
	m["N"] = "♘"
	m["b"] = "♝"
	m["B"] = "♗"
	m["q"] = "♛"
	m["Q"] = "♕"
	m["k"] = "♚"
	m["K"] = "♔"


	for {
		fmt.Println("Enter piece")
		fmt.Scanln(&p)
		fmt.Println("Enter startrank startfile endrank endfile")
		fmt.Scanf("%d %d %d %d", &startrank, &startfile, &endrank, &endfile)
		pieceToMove := board.GetPieceByRepr(m[p], *chess.NewSquare(int(startrank), int(startfile)))
		if pieceToMove != nil {
			fmt.Println("Piece Found !!")
			move := chess.NewMove(pieceToMove, *chess.NewSquare(int(startrank), int(startfile)), *chess.NewSquare(int(endrank), int(endfile)))
			board.UpdateBoard(*move)
			board.PrintBoard()
			i++
			// clear the screen
			fmt.Print("\033[H\033[2J")
			board.PrintBoard()
		} else {
			fmt.Println("Piece not found !!")
		}
	}
}

