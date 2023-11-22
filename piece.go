package main

import (
	"fmt"
)

const (
	nopiece = iota
	pawn 
	knight 
	bishop
	rook
	queen
	king
	white 
	black
)

type PieceType uint16

type Color uint8

type Square struct {
	file int
	rank int
}

func NewSquare(file int, rank int) *Square {	
	return &Square{
		file: file,
		rank: rank,
	}
}

type IPiece interface{
	CalcAvailableMoves() []Square
	Move(i Square, j Square)
}

type Piece struct {
	availableMoves []Square
	color Color
	pos Square
}

type Pawn struct {
	*Piece
}


func NewPawn(pos Square, color Color) *Pawn {
	// NewPawn creates a new Pawn with the given position and color.
	//
	// pos: The position of the Pawn on the chessboard.
	// color: The color of the Pawn (either White or Black).
	// Returns a pointer to the newly created Pawn.
	p := &Pawn{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}


func (pawn *Pawn) CalcAvailableMoves() []Square {
	// CalcAvailableMoves calculates the available legal moves for the Pawn.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	legalMoves := make([]Square, 0)
	var increment int
	
	if pawn.Piece.color == black {
		increment = -1
	} else {
		increment = 1 
	}

	if pawn.Piece.pos.file == 1 && pawn.Piece.pos.rank == 1 {
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.file, pawn.Piece.pos.rank + 2*increment})
	}
	legalMoves = append(legalMoves, Square{pawn.Piece.pos.file, pawn.Piece.pos.rank + 1*increment})
	legalMoves = append(legalMoves, Square{pawn.Piece.pos.file + 1*increment, pawn.Piece.pos.rank + 1*increment})

	pawn.Piece.availableMoves = legalMoves
	return legalMoves
}


func (pawn *Pawn) Move(destination Square) {
	
}


func (piece *Piece) HandlePieceMovement( destination Square) {
	// HandlePawnMovement handles the movement of a piece to a given destination square.
	//
	// The function takes a pointer to a Piece struct and a Square struct as its parameters.
	// It checks if the destination is a valid move for the piece by iterating through the
	// available moves of the piece. If a valid move is found, the piece's position is updated
	// to the destination square and the function returns. If the destination is not a valid
	// Check if the destination is a valid move for the piece
	for _, move := range piece.availableMoves {
		if move == destination {
			// Move the piece to the destination square
			piece.pos = destination
			fmt.Println("Piece moved to ", destination)
			return
		}
	}
	
	// If the destination is not a valid move, handle the error or invalid move here
	// ...
}

// Temporary main functin for testing purposes
func main () {
 	p := NewPawn(Square{1, 1}, white)
	fmt.Println("pawn is ", p)
	fmt.Println("pawn squares are ", p.Piece.availableMoves)
	p.HandlePieceMovement(p.availableMoves[1])
	p.CalcAvailableMoves()
	fmt.Println("pawn squares are ", p.Piece.availableMoves)
}

