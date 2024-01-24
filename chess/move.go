package chess

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
