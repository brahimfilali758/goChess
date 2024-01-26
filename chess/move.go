package chess

import "fmt"

type Move struct {
	piece   *Piece
	start   Square
	end     Square
	capture bool
}

func NewMove(piece *Piece, start Square, end Square, capture bool) *Move {
	return &Move{
		piece:   piece,
		start:   start,
		end:     end,
		capture: capture,
	}
}

func (m *Move) String() string { 
	return fmt.Sprintf("%s : start:{%d} end:{%d}", m.piece.repr, m.start, m.end)
}
