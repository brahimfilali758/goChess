package chess

import (
	"fmt"
)

type PieceType uint16
type Color uint8

func (p Color) Swap() Color {
	if p == White {
		return Black
	}
	return White
}

const (
	nopiece = iota
	pawn 
	knight 
	bishop
	rook
	queen
	king
	White 
	Black
)


func (c Color) String() string {
	if c == White {
		return "white"
	}
	return "black"
}

type Square struct {
	rank int
	file int
}

func NewSquare(rank int, file int) *Square {	
	return &Square{
		rank: rank,
		file: file,
	}
}

func (s Square) GetFile() int {
	return s.file
}

func (s Square) GetRank() int {
	return s.rank
}

func (s Square) InBoard() bool {
	if s.rank < 1 || s.rank > 8 || s.file < 1 || s.file > 8 {
		return false
	}
	return true
}


type Piece struct{
	pieceType PieceType
	color Color
	pos Square
	repr string
	availableMoves []Square
}


func (p Piece) String() string {
	return p.repr + " " + p.color.String() + fmt.Sprintf("(%d, %d)", p.pos.rank, p.pos.file)
}

func NewPiece(pieceType PieceType, pos Square, color Color) *Piece {

	p:= &Piece{
		pieceType: pieceType,
		color: color,
		pos: pos,
	}
	// p.availableMoves = p.CalcAvailableMoves()
	switch color {
		case White:
			switch pieceType {	
				case pawn:
					p.repr = "p"
				case knight:
					p.repr = "n"
				case bishop:
					p.repr = "b"
				case rook:
					p.repr = "r"
				case queen:
					p.repr = "q"
				case king:
					p.repr = "k"
			}
		case Black:
			switch pieceType {	
				case pawn:
					p.repr = "P"
				case knight:
					p.repr = "N"
				case bishop:
					p.repr = "B"
				case rook:
					p.repr = "R"
				case queen:
					p.repr = "Q"
				case king:
					p.repr = "K"
			}
	}
	return p
}

func (p Piece) GetColor() Color {
	return p.color
}

func (p Piece) GetPos() Square {
	return p.pos
}

func (p Piece) GetRepr() string {
	return p.repr
}

func (p Piece) GetAvailableMoves() []Square {
	return p.availableMoves
}

func (p Piece) GetPieceType() PieceType {
	return p.pieceType
}


func CalcAvailableMovesPawn(pos Square, color Color) []Square {
	availableMoves := make([]Square, 0)
	var increment int

	if color == Black {
		increment = -1
	} else {
		increment = 1
	}

	if pos.rank == 7 && color == Black {
		availableMoves = append(availableMoves, Square{pos.rank + 2*increment, pos.file})
	} else if pos.rank == 2 && color == White {
		availableMoves = append(availableMoves, Square{pos.rank + 2*increment, pos.file})
	}

	availableMoves = append(availableMoves, Square{pos.rank + 1*increment, pos.file})
	// en passant
	// availableMoves = append(availableMoves, Square{pos.rank + 1*increment, pos.file + 1*increment})
	// availableMoves = append(availableMoves, Square{pos.rank + 1*increment, pos.file - 1*increment})
	return availableMoves
}


func CalcAvailableMovesKnight(pos Square) []Square {
	// CalcAvailableMoves calculates the available legal moves for the Knight.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	availableMoves := make([]Square, 0)
	for rank := -2; rank <= 2; rank += 4 {
		for file := -1; file <= 1; file += 2 {
			availableMoves = append(availableMoves, Square{pos.rank + rank, pos.file + file})
		}
	}
	for rank := -1; rank <= 1; rank += 2 {
		for file := -2; file <= 2; file += 4 {
			availableMoves = append(availableMoves, Square{pos.rank + rank, pos.file + file})
		}
	}
	return availableMoves
}



func  CalcAvailableMovesBishop(pos Square, position Position) []Square {
	// CalcAvailableMoves calculates the available legal moves for the Bishop.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	return DiagLimits(pos, position)
}


func CalcAvailableMovesRook(pos Square) []Square {
	// CalcAvailableMoves calculates the available legal moves for the Rook.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	legalMoves := make([]Square, 0)
	for file := 1; file < 9; file++ {
		legalMoves = append(legalMoves, Square{pos.rank, file})
	}
	for rank := 1; rank < 9; rank++ {
		legalMoves = append(legalMoves, Square{rank, pos.file})
	}
	return legalMoves
}


func CalcAvailableMovesQueen(pos Square) []Square {
	// CalcAvailableMoves calculates the available legal moves for the Rook.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	availableMoves := make([]Square, 0)
	for file := 1; file < 9; file++ {
		availableMoves = append(availableMoves, Square{pos.rank, file})
	}
	for rank := 1; rank < 9; rank++ {
		availableMoves = append(availableMoves, Square{rank, pos.file})
	}
	return availableMoves
}


func CalcAvailableMovesKing(pos Square) []Square {
	// CalcAvailableMoves calculates the available legal moves for the Rook.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	availableMoves := []Square{
		{pos.rank - 1, pos.file - 1},
		{pos.rank - 1, pos.file},
		{pos.rank - 1, pos.file + 1},
		{pos.rank, pos.file - 1},
		{pos.rank, pos.file + 1},
		{pos.rank + 1, pos.file - 1},
		{pos.rank + 1, pos.file},
		{pos.rank + 1, pos.file + 1},
	}
	return availableMoves
}

func (piece *Piece) HandlePieceMovement(destination Square) {
	fmt.Println("Piece" , piece ," moved to ", destination)		
	piece.pos = destination
}




func (p *Piece) CalcaLegalMoves(b *Board) {
	// fmt.Println("CalcaLegalMoves with length ", len(p.availableMoves), " and ", p.availableMoves)
	legalMoves := make([]Square, 0)
	for _, square := range p.availableMoves {
		if square.InBoard() && (b.GetPiece(square.rank, square.file) == nil || b.GetPiece(square.rank, square.file).color != p.color) {
			fmt.Println("NOT removing square ", square)
			legalMoves = append(legalMoves, square)
		} else {
			fmt.Println("removing square ", square)
		}
	}
	p.availableMoves = legalMoves
	fmt.Println("CalcaLegalMoves After with length ", len(p.availableMoves), " and ", p.availableMoves)
}

