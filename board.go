package chess

import (
	"fmt"	
)

type Player struct {
	// The color of the player
	Color Color
	username string
}

type Board struct {
	// The pieces on the board
	uuid string
	Pieces []*Piece
	isChecked bool
	whiteKing *Piece
	blackKing *Piece
}

// create a new board
func NewBoard() *Board {
	return &Board{}
}

// add a piece to the board
func (b *Board) AddPiece(p Piece) {
	b.Pieces = append(b.Pieces, p)
}

// remove a piece from the board
func (b *Board) RemovePiece(p *Piece) {
	for i, piece := range b.*Pieces {
		if piece == p {
			b.Pieces = append(b.Pieces[:i], b.Pieces[i+1:]...)
			return
		}
	}
}

func (b *Board) CalcIsChecked(move Move) {
	if move.piece.Piece.color == white {
		if move.end == b.blackKing.pos {
			b.isChecked = true
		}
	}
	else if move.piece.Piece.color == black {
		if move.end == b.whiteKing.pos {
			b.isChecked = true
		}
	}
}

func (b *Board) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(b.Pieces[i*8+j].Piece.color)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
