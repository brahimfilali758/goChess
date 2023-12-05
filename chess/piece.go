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
	White 
	Black
)

type PieceType uint16

type Color uint8

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

func (s *Square) InBoard() bool {
	if s.rank < 1 || s.rank > 8 || s.file < 1 || s.file > 8 {
		return false
	}
	return true
}

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

type IPiece interface{
	CalcAvailableMoves() []Square
	CalcaLegalMoves(b *Board) []Square
}

type Piece struct {
	availableMoves []Square
	color Color
	pos Square
	repr string
	jumps bool
}

type Pawn struct {
	*Piece
}


func (p *Piece) String() string {
	return p.repr + " " + p.color.String() + fmt.Sprintf("(%d, %d)", p.pos.rank, p.pos.file)
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
			jumps: false,
		},
	}
	if color == White {
		p.repr = "♙"
	} else if color == Black {
		p.repr = "♟"
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

	if pawn.Piece.color == Black {
		increment = -1
		// Starting position
		if pawn.Piece.pos.rank == 7 {
			legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 2*increment, pawn.Piece.pos.file})
		}
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file})
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file + 1*increment})
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file - 1*increment})
	} else {
		increment = 1 
		// Starting position
		if  pawn.Piece.pos.rank == 2 {
			legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank+2*increment , pawn.Piece.pos.file})
		}
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file})
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file + 1*increment})
		legalMoves = append(legalMoves, Square{pawn.Piece.pos.rank + 1*increment, pawn.Piece.pos.file - 1*increment})
	}

	pawn.Piece.availableMoves = legalMoves
	return legalMoves
}

func (pawn *Pawn) GetAvailableStringMoves() []Square {
	// m := make(map[string]int)
	return 	[]Square{}
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
	if color == White {
		p.repr = "♘"
	} else if color == Black {
		p.repr = "♞"
	}
	p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}


func (knight *Knight) CalcAvailableMoves() []Square {
	// CalcAvailableMoves calculates the available legal moves for the Knight.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	legalMoves := make([]Square, 0)
	for rank := -2; rank <= 2; rank += 4 {
		for file := -1; file <= 1; file += 2 {
			legalMoves = append(legalMoves, Square{knight.Piece.pos.rank + rank, knight.Piece.pos.file + file})
		}
	}
	for rank := -1; rank <= 1; rank += 2 {
		for file := -2; file <= 2; file += 4 {
			legalMoves = append(legalMoves, Square{knight.Piece.pos.rank + rank, knight.Piece.pos.file + file})
		}
	}
	knight.Piece.availableMoves = legalMoves
	return legalMoves
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
	if color == White {
		p.repr = "♗"
	} else if color == Black {
		p.repr = "♝"
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
	if color == White {
		p.repr = "♖"
	} else if color == Black {
		p.repr = "♜"
	}
	p.Piece.availableMoves = p.CalcAvailableMoves()
	return p
}

func (r *Rook) CalcAvailableMoves() []Square {
	// CalcAvailableMoves calculates the available legal moves for the Rook.
	//
	// It does not take any parameters.
	// It returns a slice of Square.
	legalMoves := make([]Square, 0)
	for file := 1; file < 9; file++ {
		legalMoves = append(legalMoves, Square{r.Piece.pos.rank, file})
	}
	for rank := 1; rank < 9; rank++ {
		legalMoves = append(legalMoves, Square{rank, r.Piece.pos.file})
	}
	r.Piece.availableMoves = legalMoves
	return legalMoves
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
	if color == White {
		p.repr = "♕"
	} else if color == Black {
		p.repr = "♛"
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
	if color == White {
		p.repr = "♔"
	} else if color == Black {
		p.repr = "♚"
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
	fmt.Println("Piece available moves ", piece.availableMoves)
	for _, move := range piece.availableMoves {
		if move == destination {
			// Move the piece to the destination square
			piece.pos = destination
			fmt.Println("Piece moved to ", destination)
			// Return from the function
			return
		}
	}
	
	// If the destination is not a valid move, handle the error or invalid move here
	fmt.Println("Invalid move")
}

func (p *Piece) CalcaLegalMoves(b *Board) {
	fmt.Println("CalcaLegalMoves with length ", len(p.availableMoves), " and ", p.availableMoves)
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

