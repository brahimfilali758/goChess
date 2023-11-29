package chess

import (
	"strings"
	"strconv"
)



func GenerateBoardFromFen(fen_str string) *Board{
	board := NewBoard()
	rankStrs := strings.Split(fen_str, "/")
	for i, rankStr := range rankStrs {
		j := 0
		for _, c := range rankStr {
			if c == 'p' {
				board.Pieces = append(board.Pieces, NewPawn(*NewSquare(i, j), Black).Piece)
			} else if c == 'r' {
				board.Pieces = append(board.Pieces, NewRook(*NewSquare(i, j), Black).Piece)
			} else if c == 'n' {
				board.Pieces = append(board.Pieces, NewKnight(*NewSquare(i, j), Black).Piece)
			} else if c == 'b' {
				board.Pieces = append(board.Pieces, NewBishop(*NewSquare(i, j), Black).Piece)
			} else if c == 'q' {
				board.Pieces = append(board.Pieces, NewQueen(*NewSquare(i, j), Black).Piece)
			} else if c == 'k' {
				board.Pieces = append(board.Pieces, NewKing(*NewSquare(i, j), Black).Piece)
			} else if c == 'P' {
				board.Pieces = append(board.Pieces, NewPawn(*NewSquare(i, j), White).Piece)
			} else if c == 'R' {
				board.Pieces = append(board.Pieces, NewRook(*NewSquare(i, j), White).Piece)
			} else if c == 'N' {
				board.Pieces = append(board.Pieces, NewKnight(*NewSquare(i, j), White).Piece)
			} else if c == 'B' {
				board.Pieces = append(board.Pieces, NewBishop(*NewSquare(i, j), White).Piece)
			} else if c == 'Q' {
				board.Pieces = append(board.Pieces, NewQueen(*NewSquare(i, j), White).Piece)
			} else if c == 'K' {
				board.Pieces = append(board.Pieces, NewKing(*NewSquare(i, j), White).Piece)
			} else if  i, err := strconv.Atoi(string(c)); err == nil {
				for k := 0; k < i; k++ {
					j += 1
				}
			}
			j += 1
		}
	}
	return board

}	