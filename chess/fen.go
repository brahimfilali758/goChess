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
				board.AddPiece(NewPiece(pawn, *NewSquare(8 - i, j+1), Black))
			} else if c == 'r' {
				board.AddPiece(NewPiece(rook, *NewSquare(8 - i, j+1), Black))
			} else if c == 'n' {
				board.AddPiece(NewPiece(knight, *NewSquare(8 - i, j+1), Black))
			} else if c == 'b' {
				board.AddPiece(NewPiece(bishop, *NewSquare(8 - i, j+1), Black))
			} else if c == 'q' {
				board.AddPiece(NewPiece(queen, *NewSquare(8 - i, j+1), Black))
			} else if c == 'k' {
				_king := NewPiece(king, *NewSquare(8 - i, j+1), Black)
				board.AddPiece(_king)
				board.blackKing = _king
			} else if c == 'P' {
				board.AddPiece(NewPiece(pawn, *NewSquare(8 - i, j+1), White))
			} else if c == 'R' {
				board.AddPiece(NewPiece(rook, *NewSquare(8 - i, j+1), White))
			} else if c == 'N' {
				board.AddPiece(NewPiece(knight, *NewSquare(8 - i, j+1), White))
			} else if c == 'B' {
				board.AddPiece(NewPiece(bishop, *NewSquare(8 - i, j+1), White))
			} else if c == 'Q' {
				board.AddPiece(NewPiece(queen, *NewSquare(8 - i, j+1), White))
			} else if c == 'K' {
				_king := NewPiece(king, *NewSquare(8 - i, j+1), White)
				board.AddPiece(_king)
				board.whiteKing = _king
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