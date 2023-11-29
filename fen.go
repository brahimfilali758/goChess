package chess

import (
	"strings"
	"strconv"
)



func GenerateBoardFromFen(fen_str: string) *Board{
	board = NewBoard()
	rankStrs := strings.Split(boardStr, "/")
	for i, rankStr := range rankStrs {
		j := 0
		for _, c := range rankStr {
			if c == 'p' {
				board.pieces = append(board.pieces, NewPawn(NewSquare(i, j), black))
			}
			else if c == 'r' {
				board.pieces = append(board.pieces, NewRook(NewSquare(i, j), black))
			}
			else if c == 'n' {
				board.pieces = append(board.pieces, NewKnight(NewSquare(i, j), black))
			}
			else if c == 'b' {
				board.pieces = append(board.pieces, NewBishop(NewSquare(i, j), black))
			}
			else if c == 'q' {
				board.pieces = append(board.pieces, NewQueen(NewSquare(i, j), black))
			}
			else if c == 'k' {
				board.pieces = append(board.pieces, NewKing(NewSquare(i, j), black))
			}
			else if c == 'P' {
				board.pieces = append(board.pieces, NewPawn(NewSquare(i, j), white))
			}
			else if c == 'R' {
				board.pieces = append(board.pieces, NewRook(NewSquare(i, j), white))
			}
			else if c == 'N' {
				board.pieces = append(board.pieces, NewKnight(NewSquare(i, j), white))
			}
			else if c == 'B' {
				board.pieces = append(board.pieces, NewBishop(NewSquare(i, j), white))
			}
			else if c == 'Q' {
				board.pieces = append(board.pieces, NewQueen(NewSquare(i, j), white))
			}
			else if c == 'K' {
				board.pieces = append(board.pieces, NewKing(NewSquare(i, j), white))
			}
			else i, err := strconv.Atoi(s); err == nil {
				for k := 0; k < i; k++ {
					j += 1
				}
			}
			j += 1
		}
	}

}	