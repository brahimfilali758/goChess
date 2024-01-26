package chess

import (
	"errors"
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

func (p *Position) CopyPosition() *Position {
	return &Position{
		playerTurn: p.playerTurn,
		board:       p.board.CopyBoard(),
		validMoves: p.validMoves,
	}
}


func (p *Position) InCheck() bool {
	return p.board.isChecked
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
					if p.board.GetPiece(square.rank, square.file) == nil  {
						validMoves = append(validMoves, Move{piece, piece.pos, square, false})
					} else {
						if p.board.GetPiece(square.rank, square.file).color != piece.color && p.board.GetPiece(square.rank, square.file).pieceType != king {
							validMoves = append(validMoves, Move{piece, piece.pos, square, true})
						}
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
	}
	return legalMoves
}

func (p *Position) UpdatePosition(move Move) error {
	var err error
	fmt.Println(" ------------- Update position ", "with move ", move)
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
	// fmt.Println("Move is ,", move, "searched in ", legalMoves)
	if tmp {
		p.board.UpdateBoard(move)
		fmt.Println("Move done !")
		err = nil
	} else {
		err = errors.New("invalid move")
	}
		
	p.validMoves = p.CalcaLegalMoves(p.ValidMoves())
	p.board.isChecked = p.CalcIsChecked()
	fmt.Println("Is checked ", p.board.isChecked)
	// calc is checked with legelmoves, before recalculating valid moves of updated position
	// for _, move := range p.validMoves {
	// 	fmt.Println("Valid move ", move.piece.repr, move.start, move.end)
	// }

	return err
}


func (p *Position) SwapTurn() {
	fmt.Println("Swapping player turn from ", p.playerTurn, " to ", p.playerTurn.Swap())
	p.playerTurn = p.playerTurn.Swap()
}


func (p *Position) CalcIsChecked() bool {
	fmt.Println("------- CalcIsChecked----------")
	var TheKing *Piece
	if p.playerTurn == White {
		TheKing = p.board.blackKing
	} else {
		TheKing = p.board.whiteKing
	}
	for _, move :=range p.validMoves {
		if move.piece.color != TheKing.color {
			// fmt.Println("Checking if ", move.piece.repr, move.piece.color, " is checking king with ", move.start, move.end ," and king is ", TheKing.String())
			if move.end == TheKing.pos {
				fmt.Println("King checked by ", move.piece.repr, move.start, move.end)
				return true
			}
		}
	}
	return false
}


func (p *Position) CalcMovesToStopCheck() {
	fmt.Println("------- CalcMovesToStopCheck  -----------")
	// validMoves := new_pos.ValidMoves()
	fmt.Println("valid moves are ", p.validMoves)
	// p.validMoves = make([]Move, 0)
	for _, move := range p.ValidMoves() {
		if move.piece.color != p.playerTurn {
			new_pos := p.CopyPosition()
			fmt.Println("new pos created with black king ", new_pos.board.blackKing.pos)
			new_pos.UpdatePosition(move)
			new_pos.board.PrintBoard()
			if !new_pos.board.isChecked {
				fmt.Println("Move ", move, "prevents checking the king")
				// p.validMoves = append(p.validMoves, move)
			}
		}
	}
}