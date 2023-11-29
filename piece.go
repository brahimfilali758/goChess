package chess

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

type Move struct {
	piece *Piece
	start Square
	end Square
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

func (pawn *Pawn) GetAvailableStringMoves() []Square {
	m = make(map[string]int)

	for _, move := range pawn.Piece.availableMoves :
		
}

func (pawn *Pawn) Move(destination Square) {
	
}


type Knight struct {
	*Piece
}

func NewKnight(pos Square, color Color) *Knight {
	// NewKnight creates a new Knight with the given position and color.
	//
	// pos: The position of the Knight on the chessboard.
	// color: The color of the Knight (either White or Black).
	// Returns a pointer to the newly created Knight.
	p := &Knight{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	// p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}


func (knight *Knight) CalcAvailableMoves() []Square {
	// CalcAvailableMoves calculates the available legal moves for the Knight.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	legalMoves := make([]Square, 0)
}

func (knight *Knight) Move(destination Square) {
	
}


type Bishop struct {
	*Piece
}

func NewBishop(pos Square, color Color) *Bishop {
	// NewBishop creates a new Bishop with the given position and color.
	//
	// pos: The position of the Bishop on the chessboard.
	// color: The color of the Bishop (either White or Black).
	// Returns a pointer to the newly created Bishop.
	p := &Bishop{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	// p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}

type Rook struct {
	*Piece
}

func NewRook(pos Square, color Color) *Rook {
	// NewRook creates a new Rook with the given position and color.
	//
	// pos: The position of the Rook on the chessboard.
	// color: The color of the Rook (either White or Black).
	// Returns a pointer to the newly created Rook.
	p := &Rook{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	// p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}

type Queen struct {
	*Piece
}

func NewQueen(pos Square, color Color) *Queen {
	// NewQueen creates a new Queen with the given position and color.
	//
	// pos: The position of the Queen on the chessboard.
	// color: The color of the Queen (either White or Black).
	// Returns a pointer to the newly created Queen.
	p := &Queen{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	// p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}

type King struct {
	*Piece
}

func NewKing(pos Square, color Color) *King {
	// NewKing creates a new King with the given position and color.
	//
	// pos: The position of the King on the chessboard.
	// color: The color of the King (either White or Black).
	// Returns a pointer to the newly created King.
	p := &King{
		Piece: &Piece{
			color: color,
			pos:   pos,
		},
	}
	// p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
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



