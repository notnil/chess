package chess

type Piece struct {
	uni string
}

var (
	WKing   = &Piece{uni: "♔"}
	WQueen  = &Piece{uni: "♕"}
	WRook   = &Piece{uni: "♖"}
	WBishop = &Piece{uni: "♗"}
	WKnight = &Piece{uni: "♘"}
	WPawn   = &Piece{uni: "♙"}
	BKing   = &Piece{uni: "♚"}
	BQueen  = &Piece{uni: "♛"}
	BRook   = &Piece{uni: "♜"}
	BBishop = &Piece{uni: "♝"}
	BKnight = &Piece{uni: "♞"}
	BPawn   = &Piece{uni: "♟"}
)

func (p *Piece) String() string {
	return p.uni
}

func (p *Piece) color() color {
	switch p {
	case WKing, WQueen, WRook, WBishop, WKnight, WPawn:
		return white
	}
	return black
}

func (p *Piece) pieceType() pieceType {
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
	panic(p.uni)
}

type color int

const (
	white color = iota + 1
	black
)

func (c color) other() color {
	if c == white {
		return black
	}
	return white
}

func (c color) rankStep() int {
	if c == white {
		return 1
	}
	return -1
}

func (c color) backRank() rank {
	if c == white {
		return rank1
	}
	return rank8
}

type pieceType int

const (
	king pieceType = iota + 1
	queen
	rook
	bishop
	knight
	pawn
)

func (p pieceType) isPromotable() bool {
	switch p {
	case queen, rook, bishop, knight:
		return true
	}
	return false
}
