package chess

import (
	"strconv"
	"strings"
)



func GenerateBoardFromFen(fen_str string) *Board{
	board := NewBoard()
	rankStrs := strings.Split(fen_str, "/")
	for i, rankStr := range rankStrs {
		j := 0
		for _, c := range rankStr {
			if c == 'p' {
				board.Pieces = append(board.Pieces, NewPiece(pawn, *NewSquare(8 - i, j+1), Black))
			} else if c == 'r' {
				board.Pieces = append(board.Pieces, NewPiece(rook, *NewSquare(8 - i, j+1), Black))
			} else if c == 'n' {
				board.Pieces = append(board.Pieces, NewPiece(knight, *NewSquare(8 - i, j+1), Black))
			} else if c == 'b' {
				board.Pieces = append(board.Pieces, NewPiece(bishop, *NewSquare(8 - i, j+1), Black))
			} else if c == 'q' {
				board.Pieces = append(board.Pieces, NewPiece(queen, *NewSquare(8 - i, j+1), Black))
			} else if c == 'k' {
				board.Pieces = append(board.Pieces, NewPiece(king, *NewSquare(8 - i, j+1), Black))
			} else if c == 'P' {
				board.Pieces = append(board.Pieces, NewPiece(pawn, *NewSquare(8 - i, j+1), White))
			} else if c == 'R' {
				board.Pieces = append(board.Pieces, NewPiece(rook, *NewSquare(8 - i, j+1), White))
			} else if c == 'N' {
				board.Pieces = append(board.Pieces, NewPiece(knight, *NewSquare(8 - i, j+1), White))
			} else if c == 'B' {
				board.Pieces = append(board.Pieces, NewPiece(bishop, *NewSquare(8 - i, j+1), White))
			} else if c == 'Q' {
				board.Pieces = append(board.Pieces, NewPiece(queen, *NewSquare(8 - i, j+1), White))
			} else if c == 'K' {
				board.Pieces = append(board.Pieces, NewPiece(king, *NewSquare(8 - i, j+1), White))
			} else if  i, err := strconv.Atoi(string(c)); err == nil {
				for k := 0; k < i-1; k++ {
					j += 1
				}
			}
			j += 1
		}
	}
	return board

}	