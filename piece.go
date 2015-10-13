package chess

// Color represents the color of a chess piece.
type Color int

const (
	// NoColor represents no color
	NoColor Color = iota
	// White represents the color white
	White
	// Black represents the color black
	Black
)

// Other returns the opposie color of the receiver.
func (c Color) Other() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	return NoColor
}

// String implements the fmt.Stringer interface and returns
// the color's FEN compatible notation.
func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	}
	return "-"
}

// PieceType is the type of a piece.
type PieceType int

const (
	// NoPiece represents a lack of piece type
	NoPiece PieceType = iota
	// King represents a king
	King
	// Queen represents a queen
	Queen
	// Rook represents a rook
	Rook
	// Bishop represents a bishop
	Bishop
	// Knight represents a knight
	Knight
	// Pawn represents a pawn
	Pawn
)

// PieceTypes returns a slice of all piece types.
func PieceTypes() []PieceType {
	return []PieceType{King, Queen, Rook, Bishop, Knight, Pawn}
}

func (p PieceType) isPromotable() bool {
	switch p {
	case Queen, Rook, Bishop, Knight:
		return true
	}
	return false
}

// Piece is a piece type with a color.
type Piece struct {
	t   PieceType
	c   Color
	uni string
}

var (
	// WhiteKing is a white king
	WhiteKing = &Piece{t: King, c: White, uni: "♔"}
	// WhiteQueen is a white queen
	WhiteQueen = &Piece{t: Queen, c: White, uni: "♕"}
	// WhiteRook is a white rook
	WhiteRook = &Piece{t: Rook, c: White, uni: "♖"}
	// WhiteBishop is a white bishop
	WhiteBishop = &Piece{t: Bishop, c: White, uni: "♗"}
	// WhiteKnight is a white knight
	WhiteKnight = &Piece{t: Knight, c: White, uni: "♘"}
	// WhitePawn is a white pawn
	WhitePawn = &Piece{t: Pawn, c: White, uni: "♙"}
	// BlackKing is a black king
	BlackKing = &Piece{t: King, c: Black, uni: "♚"}
	// BlackQueen is a black queen
	BlackQueen = &Piece{t: Queen, c: Black, uni: "♛"}
	// BlackRook is a black rook
	BlackRook = &Piece{t: Rook, c: Black, uni: "♜"}
	// BlackBishop is a black bishop
	BlackBishop = &Piece{t: Bishop, c: Black, uni: "♝"}
	// BlackKnight is a black knight
	BlackKnight = &Piece{t: Knight, c: Black, uni: "♞"}
	// BlackPawn is a black pawn
	BlackPawn = &Piece{t: Pawn, c: Black, uni: "♟"}
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
	return nil
}

// Type returns the type of the piece.
func (p *Piece) Type() PieceType {
	return p.t
}

// Color returns the color of the piece.
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
	return ""
}
