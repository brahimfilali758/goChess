package chess

import (
	"fmt"
	"slices"
)

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetMoveParams(move string, p Position) (string, Square, Square) {
	var pieceRepr string
	var start Square
	var end Square
	var rank string
	var file string
	if len(move) == 2 {
		if p.playerTurn == White {
			pieceRepr = "p"
		} else {
			pieceRepr = "P"
		}
		rank = string(move[1])
		file = string(move[0])
		
	} else if len(move) == 3 {
		pieceRepr = move[:1]
		rank = string(move[2])
		file = string(move[1])
		fmt.Println("move len 3 , piece ", pieceRepr, " rank ", rank, " file ", file)
	}
	end = *NewSquare(int(rank[0] - '0'), int(file[0] - 'a') + 1)
	for _ , piece := range p.board.Pieces {
		if piece.repr == pieceRepr {
			fmt.Println("comparing with ", piece.repr, "with moves ", piece.availableMoves)
			if slices.Contains(piece.availableMoves, end) {
				start = piece.pos
				break
			}
		}
	}
	
	fmt.Println("piece ", pieceRepr, " start ", start, " end ", end)
	return pieceRepr, start, end
}