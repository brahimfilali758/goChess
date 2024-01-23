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
	position := chess.NewPosition(board)
	board.PrintBoard()
	fmt.Println("Board printed")
	var p string
	var startrank int 
	var startfile int
	var endrank int
	var endfile int
	var move string


	for {

		fmt.Println(position.GetPlayerTurn() ,"To move, Type move :")
		fmt.Scanf("%s\n", &move)
		fmt.Println("Scanf done with move ", move)
		p_ , start , end := chess.GetMoveParams(move, *position)
		p , startrank, startfile, endrank, endfile = "", 0, 0, 0, 0
		p = p_
		startfile = start.GetFile()
		startrank = start.GetRank()
		endfile = end.GetFile()
		endrank = end.GetRank()
		// fmt.Println("Enter piece startrank startfile endrank endfile")
		// fmt.Scanf("%s %d %d %d %d",&p, &startrank, &startfile, &endrank, &endfile)
		if p != "" && startrank != 0 && startfile != 0 && endrank != 0 && endfile != 0 {
			pieceToMove := board.GetPieceByRepr(p, *chess.NewSquare(int(startrank), int(startfile)))
			fmt.Println(pieceToMove)
			if pieceToMove != nil {
				fmt.Println("Piece Found !!")
				move := chess.NewMove(pieceToMove, *chess.NewSquare(int(startrank), int(startfile)), *chess.NewSquare(int(endrank), int(endfile)))
				position.UpdatePosition(*move)
				board.PrintBoard()
				// clear the screen
				fmt.Print("\033[H\033[2J")
				board.PrintBoard()
			} else {
				fmt.Println("Piece not found !!")
			}
		} 
		
	}
}

