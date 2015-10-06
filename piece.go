package chess

type Color int

const (
	NoColor Color = iota
	White
	Black
)

func (c Color) Other() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	return NoColor
}

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	}
	return "-"
}

type PieceType int

const (
	King PieceType = iota + 1
	Queen
	Rook
	Bishop
	Knight
	Pawn
)

func (p PieceType) Ptr() *PieceType {
	return &p
}

func (p PieceType) isPromotable() bool {
	switch p {
	case Queen, Rook, Bishop, Knight:
		return true
	}
	return false
}

type Piece struct {
	t   PieceType
	c   Color
	uni string
}

var (
	WhiteKing   = &Piece{t: King, c: White, uni: "♔"}
	WhiteQueen  = &Piece{t: Queen, c: White, uni: "♕"}
	WhiteRook   = &Piece{t: Rook, c: White, uni: "♖"}
	WhiteBishop = &Piece{t: Bishop, c: White, uni: "♗"}
	WhiteKnight = &Piece{t: Knight, c: White, uni: "♘"}
	WhitePawn   = &Piece{t: Pawn, c: White, uni: "♙"}
	BlackKing   = &Piece{t: King, c: Black, uni: "♚"}
	BlackQueen  = &Piece{t: Queen, c: Black, uni: "♛"}
	BlackRook   = &Piece{t: Rook, c: Black, uni: "♜"}
	BlackBishop = &Piece{t: Bishop, c: Black, uni: "♝"}
	BlackKnight = &Piece{t: Knight, c: Black, uni: "♞"}
	BlackPawn   = &Piece{t: Pawn, c: Black, uni: "♟"}
)

var (
	allPieces = []*Piece{
		WhiteKing, WhiteQueen, WhiteRook, WhiteBishop, WhiteKnight, WhitePawn,
		BlackKing, BlackQueen, BlackRook, BlackBishop, BlackKnight, BlackPawn,
	}
)

func getPiece(t PieceType, c Color) *Piece {
	for _, p := range allPieces {
		if p.Color() == c && p.Type() == t {
			return p
		}
	}
	panic("unreachable")
}

func (p *Piece) Type() PieceType {
	return p.t
}

func (p *Piece) Color() Color {
	return p.c
}

func (p *Piece) String() string {
	return p.uni
}

func (p *Piece) getFENChar() string {
	for key, piece := range fenPieceMap {
		if piece == p {
			return key
		}
	}
	panic("unreachable")
}
