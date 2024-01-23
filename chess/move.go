package chess 

type Move struct {
	piece *Piece
	start Square
	end Square
}

func NewMove(piece *Piece, start Square, end Square) *Move {
	return &Move{
		piece: piece,
		start: start,
		end: end,
	}
}

