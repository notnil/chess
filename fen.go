package chess

import (
	"fmt"
	"strconv"
	"strings"
)

// https://en.wikipedia.org/wiki/Forsyth–Edwards_Notation
// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
func FEN(fen string) (*GameState, error) {
	fen = strings.TrimSpace(fen)
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return nil, fmt.Errorf("chess: fen invalid notiation %s must have 6 sections", fen)
	}
	b, err := fenBoard(parts[0])
	if err != nil {
		return nil, err
	}
	turn, ok := fenTurnMap[parts[1]]
	if !ok {
		return nil, fmt.Errorf("chess: fen invalid turn %s", parts[1])
	}
	rights, err := formCastleRights(parts[2])
	if err != nil {
		return nil, err
	}
	sq, err := formEnPassant(parts[3])
	if err != nil {
		return nil, err
	}
	halfMoveClock, err := strconv.Atoi(parts[4])
	if err != nil || halfMoveClock < 0 {
		return nil, fmt.Errorf("chess: fen invalid half move clock %s", parts[4])
	}
	moveCount, err := strconv.Atoi(parts[5])
	if err != nil || moveCount < 1 {
		return nil, fmt.Errorf("chess: fen invalid move count %s", parts[5])
	}
	return &GameState{
		board:           b,
		turn:            turn,
		castleRights:    rights,
		enPassantSquare: sq,
		halfMoveClock:   halfMoveClock,
		moveCount:       moveCount,
	}, nil
}

func fenBoard(boardStr string) (Board, error) {
	rankStrs := strings.Split(boardStr, "/")
	if len(rankStrs) != 8 {
		return nil, fmt.Errorf("chess: fen invalid board %s", boardStr)
	}
	b := map[*Square]*Piece{}
	for i, rankStr := range rankStrs {
		rank := Rank(8 - i)
		fileMap, err := fenFormRank(rankStr)
		if err != nil {
			return nil, err
		}
		for file, piece := range fileMap {
			b[getSquare(file, rank)] = piece
		}
	}
	return b, nil
}

func fenFormRank(rankStr string) (map[File]*Piece, error) {
	count := 0
	m := map[File]*Piece{}
	err := fmt.Errorf("chess: fen invalid rank %s", rankStr)
	for _, r := range rankStr {
		c := fmt.Sprintf("%c", r)
		piece := fenPieceMap[c]
		if piece == nil {
			skip, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			count += skip
			continue
		}
		count++
		m[File(count)] = piece
	}
	if count != 8 {
		return nil, err
	}
	return m, nil
}

func formCastleRights(castleStr string) (CastleRights, error) {
	// check for duplicates aka. KKkq right now is valid
	for _, s := range []string{"K", "Q", "k", "q", "-"} {
		if strings.Count(castleStr, s) > 1 {
			return "-", fmt.Errorf("chess: fen invalid castle rights %s", castleStr)
		}
	}
	for _, r := range castleStr {
		c := fmt.Sprintf("%c", r)
		switch c {
		case "K", "Q", "k", "q", "-":
		default:
			return "-", fmt.Errorf("chess: fen invalid castle rights %s", castleStr)
		}
	}
	return CastleRights(castleStr), nil
}

func formEnPassant(enPassant string) (*Square, error) {
	if enPassant == "-" {
		return nil, nil
	}
	sq := squareFromStr(enPassant)
	if sq == nil || !(sq.rank == R3 || sq.rank == R6) {
		return nil, fmt.Errorf("chess: fen invalid En Passant square %s", enPassant)
	}
	return sq, nil
}

var (
	fenPieceMap = map[string]*Piece{
		"K": WhiteKing,
		"Q": WhiteQueen,
		"R": WhiteRook,
		"B": WhiteBishop,
		"N": WhiteKnight,
		"P": WhitePawn,
		"k": BlackKing,
		"q": BlackQueen,
		"r": BlackRook,
		"b": BlackBishop,
		"n": BlackKnight,
		"p": BlackPawn,
	}

	fenTurnMap = map[string]Color{
		"w": White,
		"b": Black,
	}
)
