package chess

import (
	"fmt"	
)

type Board struct {
	// The pieces on the board
	Pieces []*Piece
	isChecked bool
	whiteKing *Piece
	blackKing *Piece
}

// create a new board
func NewBoard() *Board {
	return &Board{
		Pieces: make([]*Piece, 0),
		isChecked: false,
		whiteKing: nil,
		blackKing: nil,	
	}
}

// add a piece to the board
func (b *Board) AddPiece(p Piece) {
	b.Pieces = append(b.Pieces, &p)
}

// remove a piece from the board
func (b *Board) RemovePiece(p *Piece) {
	for i, piece := range b.Pieces {
		if piece == p {
			b.Pieces = append(b.Pieces[:i], b.Pieces[i+1:]...)
			return
		}
	}
}

// get a piece from the board
func (b *Board) GetPiece(i int, j int) *Piece {
	for _, piece := range b.Pieces {
		if piece.pos.rank == i && piece.pos.file == j {
			return piece
		}
	}
	return nil
}

func (b *Board) CalcIsChecked(move Move) {
	if move.piece.color == White {
		if move.end == b.blackKing.pos {
			b.isChecked = true
		}
	} else if move.piece.color == Black {
		if move.end == b.whiteKing.pos {
			b.isChecked = true
		}
	}	
}

func (b *Board) PrintBoard() {
	for i := 8; i > 0; i-- {
		for j := 1; j <= 8; j++ {
			piece := b.GetPiece(i, j)
			if piece != nil {
				fmt.Printf(piece.repr)
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
}


func (b *Board) UpdateBoard(move Move) {
	pieceDestination := b.GetPiece(move.end.rank, move.end.file)
	fmt.Println("Move is , ", move, " and destination is ", pieceDestination)
	if pieceDestination == nil {
		fmt.Println("Move done with destination square empty")
		move.piece.HandlePieceMovement(move.end)
	} else if pieceDestination != nil && move.piece.color != pieceDestination.color{
		fmt.Println("Move done with capture")
		// capture destination piece
		b.RemovePiece(pieceDestination)
		move.piece.HandlePieceMovement(move.end)
	} else {
		fmt.Println("Move impossible !")
		return
	}
	// b.CalcIsChecked(move)
}

func (b *Board) GetPieceByRepr(repr string, pos Square) *Piece {
	for _, piece := range b.Pieces {
		// fmt.Println("Checking piece ", piece.repr, piece.pos)
		if piece.repr == repr && piece.pos == pos {
			return piece
		}
	}
	return nil
}

