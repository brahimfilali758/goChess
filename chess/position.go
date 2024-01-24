package chess

import (
	"fmt"
	"slices"
)

type Position struct {
	playerTurn Color
	board      *Board
	validMoves []Move
}

func NewPosition(board *Board) *Position {
	posi :=  &Position{
		playerTurn: White,
		board:      board,
		validMoves: nil,
	}
	validMoves := posi.ValidMoves()
	posi.validMoves = posi.CalcaLegalMoves(validMoves)
	return posi
}

func (p *Position) GetPlayerTurn() Color {
	return p.playerTurn
}

func (p *Position) GetBoard() *Board {
	return p.board
}

func (p *Position) GetValidMoves() []Move {
	// for _, move := range p.ValidMoves() {
	// 	fmt.Println("Valid move ", move.piece.repr, move.start, move.end)
	// }
	return p.validMoves
}

func (p *Position) ValidMoves() []Move {
	validMoves := make([]Move, 0)
	for _, piece := range p.board.Pieces {
		if piece.pieceType == knight {
			for _, square := range CalcAvailableMovesKnight(piece.pos) {
				if square.InBoard() {
					if p.board.GetPiece(square.rank, square.file) == nil {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
					} else if p.board.GetPiece(square.rank, square.file).color != piece.color {
						validMoves = append(validMoves, Move{piece, piece.pos, square, true})
					}
				}
			}
		} else if piece.pieceType == rook {
			moves := p.board.VerHorLimits(piece.pos, piece.color)
			for _, square := range moves {
				if p.board.GetPiece(square.rank, square.file) == nil {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
				} else if p.board.GetPiece(square.rank, square.file).color != piece.color {
					validMoves = append(validMoves, Move{piece, piece.pos, square, true})
				}
			}
		} else if piece.pieceType == queen {
			moves := append(p.board.VerHorLimits(piece.pos, piece.color), p.board.DiagLimits(piece.pos, piece.color)...)
			for _, square := range moves {
				if p.board.GetPiece(square.rank, square.file) == nil {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
				} else if p.board.GetPiece(square.rank, square.file).color != piece.color {
					validMoves = append(validMoves, Move{piece, piece.pos, square, true})
				}
			}
		} else if piece.pieceType == bishop {
			moves := p.board.DiagLimits(piece.pos, piece.color)
			for _, square := range moves {
				if p.board.GetPiece(square.rank, square.file) == nil {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
				} else if p.board.GetPiece(square.rank, square.file).color != piece.color {
					validMoves = append(validMoves, Move{piece, piece.pos, square, true})
				}
			}
		} else if piece.pieceType == king {
			for _, square := range CalcAvailableMovesKing(piece.pos) {
				if square.InBoard() {
					if p.board.GetPiece(square.rank, square.file) == nil {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
					} else if p.board.GetPiece(square.rank, square.file).color != piece.color {
						validMoves = append(validMoves, Move{piece, piece.pos, square, true})
					}
				}
			}
		} else if piece.pieceType == pawn {
			// fmt.Println("Pawn moves are ", p.board.PawnMoves(piece))
			validMoves = append(validMoves, p.board.PawnMoves(piece)...)
		}
	}
	return validMoves
}

func (p *Position) CalcaLegalMoves(validMoves []Move) []Move {
	legalMoves := make([]Move, 0)
	// fmt.Println("piece valid moves are :")
	// for _, move := range validMoves {
	// 	if move.piece.pieceType == pawn {
	// 		fmt.Println(move.piece.repr, move.start, move.end)
	// 	}
	// }
	// empty pieces avaiblable moves
	for _, piece := range p.board.Pieces {
		piece.availableMoves = make([]Square, 0)
	}
	for _, move := range validMoves {
		
		pOnBoard := p.board.GetPiece(move.end.rank, move.end.file)
		if pOnBoard == nil || pOnBoard.color != move.piece.color {
			move.piece.availableMoves = append(move.piece.availableMoves, move.end)
			legalMoves = append(legalMoves, move)
			// fmt.Println("piece ", move.piece.repr, "available moves are ", move.piece.availableMoves)
		}
		
		// } else {
		// 	limits := p.board.VerHorLimits(move.piece.pos, move.piece.color)
		// 	if slices.Contains(limits, move.end) {
		// 		move.piece.availableMoves = append(move.piece.availableMoves, move.end)
		// 		legalMoves = append(legalMoves, move)
		// 	}
		// }
	}
	return legalMoves
}

func (p *Position) UpdatePosition(move Move) {

	legalMoves := p.validMoves
	// fmt.Println("UpdatePosition Legal moves are ", legalMoves)
	tmp := false
	if move.piece.pieceType == knight {
		fmt.Println("move is a knight")
		tmp = true
	} else {
		if slices.Contains(legalMoves, move) {
			tmp = true
		}
	}
	fmt.Println("Move is ,", move, "searched in ", legalMoves)
	if tmp {
		p.board.UpdateBoard(move)
		fmt.Println("Move done !")
		if p.playerTurn == White {
			p.playerTurn = Black
		} else {
			p.playerTurn = White
		}
	}
	
	p.validMoves = p.CalcaLegalMoves(p.ValidMoves())
}
