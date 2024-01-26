package chess

import (
	"fmt"
	"slices"
	"strings"
)

func PrintError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetMoveParams(move string, p Position) (string, Square, Square, bool) {

	pieces := []string{"R", "N", "B", "Q", "K"}
	fmt.Println("p.playerTurn", p.playerTurn)
	var pieceRepr string
	var start Square
	var end Square
	var rank string
	var file string
	capture := false
	if len(move) == 2 {
		pieceRepr = GetPieceRepr("p", p.playerTurn)
		rank = string(move[1])
		file = string(move[0])
	} else if len(move) == 3 {
		pieceRepr = GetPieceRepr(move[:1], p.playerTurn)
		rank = string(move[2])
		file = string(move[1])
		fmt.Println("move len 3 , piece ", pieceRepr, " rank ", rank, " file ", file)
	} else if len(move) == 4 && strings.Contains(move, "x") {
		if slices.Contains(pieces, move[:1]) {
			pieceRepr = GetPieceRepr(move[:1], p.playerTurn)
		} else {
			pieceRepr = GetPieceRepr("p", p.playerTurn)
		}
		rank = string(move[3])
		file = string(move[2])
		capture = true
	}
	end = *NewSquare(int(rank[0] - '0'), int(file[0] - 'a') + 1)
	for _ , piece := range p.board.Pieces {
		if piece.repr == pieceRepr {
			fmt.Println("comparing with ", piece.repr, "in pos" , piece.pos , "with moves ", piece.availableMoves)
			if slices.Contains(piece.availableMoves, end) {
				start = piece.pos
				break
			}
		}
	}
	
	fmt.Println("piece ", pieceRepr, " start ", start, " end ", end)
	return pieceRepr, start, end, capture
}


func DiagLimits(s Square, position Position) []Square {
	availableMoves := make([]Square, 0)
	// fmt.Println("CALCULATING DIAG MOVES FOR ", s)
	rank := s.rank
	file := s.file
	for {
		rank++
		file++
		pOnBoard := position.board.GetPiece(rank, file)
		if rank > 8 || file > 8 {
			break
		}
		if pOnBoard == nil {
			availableMoves = append(availableMoves, Square{rank, file})
		} else {
			if pOnBoard.color != position.playerTurn {
				availableMoves = append(availableMoves, Square{rank, file})
			}
			break
		}
	}
	rank = s.rank
	file = s.file
	for {
		rank++
		file--
		if rank > 8 || file < 1 {
			break
		}
		pOnBoard := position.board.GetPiece(rank, file)
		if pOnBoard == nil {
			availableMoves = append(availableMoves, Square{rank, file})
		} else {
			if pOnBoard.color != position.playerTurn {
				availableMoves = append(availableMoves, Square{rank, file})
			}
			break
		}
	}

	rank = s.rank
	file = s.file
	for {
		rank--
		file++
		if rank < 1 || file > 8 {
			break
		}
		pOnBoard := position.board.GetPiece(rank, file)
		if pOnBoard == nil {
			availableMoves = append(availableMoves, Square{rank, file})
		} else {
			if pOnBoard.color != position.playerTurn {
				availableMoves = append(availableMoves, Square{rank, file})
			}
			break
		}
	}
	rank = s.rank
	file = s.file
	for {
		rank--
		file--
		if rank < 1 || file < 1 {
			break
		}
		pOnBoard := position.board.GetPiece(rank, file)
		if pOnBoard == nil {
			availableMoves = append(availableMoves, Square{rank, file})
		} else {
			if pOnBoard.color != position.playerTurn {
				availableMoves = append(availableMoves, Square{rank, file})
			}
			break
		}
	}

	return availableMoves
}


func GetPieceRepr(s string, color Color) string {
	if color == White {
		return strings.ToLower(s)
	} else {
		return strings.ToUpper(s)
	}
}