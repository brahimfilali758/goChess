package chess

import (
	"fmt"
)


var m  = map[string]string {
	"p": "♟",
	"P": "♙",
	"r": "♜",
	"R": "♖",
	"n": "♞",
	"N": "♘",
	"b": "♝",
	"B": "♗",
	"q": "♛",
	"Q": "♕",
	"k": "♚",
	"K": "♔",
}

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

func (b *Board) GetPiecePosition(repr string) []Square {
	positions := []Square{}
	for _, piece := range b.Pieces {
		if piece.repr == repr {
			positions = append(positions, piece.pos)
		}
	}
	return positions
}

func (b *Board) CalcIsChecked(move *Move) {
	color := move.piece.color
	if color == White {
		if move.end == b.blackKing.pos {
			b.isChecked = true
		}
	} else if color == Black {
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
				fmt.Printf(m[piece.repr])
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
}


func (b *Board) UpdateBoard(move Move) {
	// Update piece available moves with board
	pieceDestination := b.GetPiece(move.end.rank, move.end.file)

	// fmt.Println("Move is , ", move, " and destination is ", pieceDestination)
	if pieceDestination == nil {
		// fmt.Println("Move done with destination square empty")
		move.piece.HandlePieceMovement(move.end)
		return
	} else {
		if move.piece.color != pieceDestination.color{
			fmt.Println("Move done with capture")
			// capture destination piece
			b.RemovePiece(pieceDestination)
			move.piece.HandlePieceMovement(move.end)
			return
		} else {
			fmt.Println("Move impossible !") 
			return
		}
	} 
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

func (b *Board) Limits(pos Square, color Color) []Square {
	// fmt.Println("Checking limits for ", pos, color)
	emptySquares := make([]Square, 0)
	
	for i := pos.file+1; i <= 8; i++ {
		pOnBoard:= b.GetPiece(pos.rank, i)
		if pOnBoard == nil {
			emptySquares = append(emptySquares, Square{pos.rank, i})
		} else{
			if pOnBoard.color != color {
				emptySquares = append(emptySquares, Square{pos.rank, i})
				break
			} else if pOnBoard.color == color{
				break
			}
		} 
	}
	
	for i := pos.file-1; i >= 1; i-- {
		pOnBoard:= b.GetPiece(pos.rank, i)
		if pOnBoard == nil {
			emptySquares = append(emptySquares, Square{pos.rank, i})
		} else {
			if pOnBoard.color != color {
				emptySquares = append(emptySquares, Square{pos.rank, i})
				break
			} else if pOnBoard.color == color{
				break
			}
		} 
	}
	
	for i := pos.rank+1; i <= 8; i++ {
		// fmt.Println("i is ", i)
		pOnBoard:= b.GetPiece(i, pos.file)
		// fmt.Println("piece on board is ", pOnBoard)
		if pOnBoard == nil {
			emptySquares = append(emptySquares, Square{i, pos.file})
			// fmt.Println("empty square in " , i , pos.file)
		} else {
			if pOnBoard.color != color {
				emptySquares = append(emptySquares, Square{i, pos.file})
				break
			} else if pOnBoard.color == color{
				// fmt.Println("break in " , i , pos.file, " with ", pOnBoard)
				break
			}
		} 
	}
	
	for i := pos.rank-1; i >= 1; i-- {
		pOnBoard:= b.GetPiece(i, pos.file)
		if  pOnBoard == nil {
			emptySquares = append(emptySquares, Square{i, pos.file})
		} else {
			if pOnBoard.color != color {
				emptySquares = append(emptySquares, Square{i, pos.file})
				break
			} else if pOnBoard.color == color{
				break
			}
		} 
	}
	// fmt.Println("empty squares are ", emptySquares)
	
	return emptySquares
}


