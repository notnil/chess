package chess

type Piece string

const (
	WKing   Piece = "♔"
	WQueen  Piece = "♕"
	WRook   Piece = "♖"
	WBishop Piece = "♗"
	WKnight Piece = "♘"
	WPawn   Piece = "♙"
	BKing   Piece = "♚"
	BQueen  Piece = "♛"
	BRook   Piece = "♜"
	BBishop Piece = "♝"
	BKnight Piece = "♞"
	BPawn   Piece = "♟"
)

func (p Piece) color() color {
	switch p {
	case WKing, WQueen, WRook, WBishop, WKnight, WPawn:
		return white
	}
	return black
}

func (p Piece) pieceType() pieceType {
	switch p {
	case WKing, BKing:
		return king
	case WQueen, BQueen:
		return queen
	case WRook, BRook:
		return rook
	case WBishop, BBishop:
		return bishop
	case WKnight, BKnight:
		return knight
	case WPawn, BPawn:
		return pawn
	}
	panic("unreachable")
}

type color int

const (
	white color = iota + 1
	black
)

type pieceType int

const (
	king pieceType = iota + 1
	queen
	rook
	bishop
	knight
	pawn
)
