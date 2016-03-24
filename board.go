package chess

import (
	"strconv"
	"strings"
)

// A Board represents a chess board and its relationship between squares and pieces.
type Board struct {
	bbWhiteKing   bitboard
	bbWhiteQueen  bitboard
	bbWhiteRook   bitboard
	bbWhiteBishop bitboard
	bbWhiteKnight bitboard
	bbWhitePawn   bitboard
	bbBlackKing   bitboard
	bbBlackQueen  bitboard
	bbBlackRook   bitboard
	bbBlackBishop bitboard
	bbBlackKnight bitboard
	bbBlackPawn   bitboard
	whiteSqs      bitboard
	blackSqs      bitboard
	emptySqs      bitboard
	whiteKingSq   Square
	blackKingSq   Square
}

// SquareMap returns a mapping of squares to pieces.  A square is only added to the map if it is occupied.
func (b *Board) SquareMap() map[Square]*Piece {
	m := map[Square]*Piece{}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		p := b.piece(Square(sq))
		if p != nil {
			m[Square(sq)] = p
		}
	}
	return m
}

// Draw returns visual representation of the board useful for debugging.  Ex.
// A B C D E F G H
// 8- - - - - - - -
// 7- ♟ - - - - - ♝
// 6- - - - ♘ - - -
// 5- - - - ♟ - - -
// 4- - - - - - - -
// 3- - - - - - - -
// 2- ♖ - - - - - -
// 1- ♗ - - - - - -
func (b *Board) Draw() string {
	s := "\n A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += rank(r).String()
		for f := 0; f < numOfSquaresInRow; f++ {
			p := b.piece(getSquare(file(f), rank(r)))
			if p == nil {
				s += "-"
			} else {
				s += p.String()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

// String implements the fmt.Stringer interface and returns
// a string in the FEN board format: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
func (b *Board) String() string {
	fen := ""
	for r := 7; r >= 0; r-- {
		for f := 0; f < numOfSquaresInRow; f++ {
			sq := getSquare(file(f), rank(r))
			p := b.piece(sq)
			if p != nil {
				fen += p.getFENChar()
			} else {
				fen += "1"
			}
		}
		if r != 0 {
			fen += "/"
		}
	}
	for i := 8; i > 1; i-- {
		repeatStr := strings.Repeat("1", i)
		countStr := strconv.Itoa(i)
		fen = strings.Replace(fen, repeatStr, countStr, -1)
	}
	return fen
}

func newBoard(m map[Square]*Piece) *Board {
	b := &Board{}
	for _, p1 := range allPieces {
		bm := map[Square]bool{}
		for sq, p2 := range m {
			if p1 == p2 {
				bm[sq] = true
			}
		}
		bb := newBitboard(bm)
		b.setBBForPiece(p1, bb)
	}
	b.calcConvienceBBs()
	return b
}

func (b *Board) update(m *Move) {
	p1 := b.piece(m.s1)
	s1BB := bbSquares[m.s1]
	s2BB := bbSquares[m.s2]
	// move s1 piece to s2
	for _, p := range allPieces {
		bb := b.bbForPiece(p)
		// remove what was at s2
		b.setBBForPiece(p, bb & ^s2BB)
		// move what was at s1 to s2
		if bb.Occupied(m.s1) {
			bb = b.bbForPiece(p)
			b.setBBForPiece(p, (bb & ^s1BB)|s2BB)
		}
	}
	// check promotion
	if m.promo != NoPieceType {
		newPiece := getPiece(m.promo, p1.Color())
		// remove pawn
		bbPawn := b.bbForPiece(p1)
		b.setBBForPiece(p1, bbPawn & ^s2BB)
		// add promo piece
		bbPromo := b.bbForPiece(newPiece)
		b.setBBForPiece(newPiece, bbPromo|s2BB)
	}
	// remove captured en passant piece
	if m.HasTag(EnPassant) {
		if p1.Color() == White {
			b.bbBlackPawn = ^(bbSquares[m.s2] << 8) & b.bbBlackPawn
		} else {
			b.bbWhitePawn = ^(bbSquares[m.s2] >> 8) & b.bbWhitePawn
		}
	}
	// move rook for castle
	if p1.Color() == White && m.HasTag(KingSideCastle) {
		b.bbWhiteRook = (b.bbWhiteRook & ^bbSquares[H1]) | bbSquares[F1]
	} else if p1.Color() == White && m.HasTag(QueenSideCastle) {
		b.bbWhiteRook = (b.bbWhiteRook & ^bbSquares[A1]) | bbSquares[D1]
	} else if p1.Color() == Black && m.HasTag(KingSideCastle) {
		b.bbBlackRook = (b.bbBlackRook & ^bbSquares[H8]) | bbSquares[F8]
	} else if p1.Color() == Black && m.HasTag(QueenSideCastle) {
		b.bbBlackRook = (b.bbBlackRook & ^bbSquares[A8]) | bbSquares[D8]
	}
	b.calcConvienceBBs()
}

func (b *Board) calcConvienceBBs() {
	whiteSqs := b.bbWhiteKing | b.bbWhiteQueen | b.bbWhiteRook | b.bbWhiteBishop | b.bbWhiteKnight | b.bbWhitePawn
	blackSqs := b.bbBlackKing | b.bbBlackQueen | b.bbBlackRook | b.bbBlackBishop | b.bbBlackKnight | b.bbBlackPawn
	emptySqs := ^(whiteSqs | blackSqs)
	b.whiteSqs = whiteSqs
	b.blackSqs = blackSqs
	b.emptySqs = emptySqs
	b.whiteKingSq = NoSquare
	b.blackKingSq = NoSquare

	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		sqr := Square(sq)
		if b.bbWhiteKing.Occupied(sqr) {
			b.whiteKingSq = sqr
		} else if b.bbBlackKing.Occupied(sqr) {
			b.blackKingSq = sqr
		}
	}
}

func (b *Board) copy() *Board {
	return &Board{
		whiteSqs:      b.whiteSqs,
		blackSqs:      b.blackSqs,
		emptySqs:      b.emptySqs,
		whiteKingSq:   b.whiteKingSq,
		blackKingSq:   b.blackKingSq,
		bbWhiteKing:   b.bbWhiteKing,
		bbWhiteQueen:  b.bbWhiteQueen,
		bbWhiteRook:   b.bbWhiteRook,
		bbWhiteBishop: b.bbWhiteBishop,
		bbWhiteKnight: b.bbWhiteKnight,
		bbWhitePawn:   b.bbWhitePawn,
		bbBlackKing:   b.bbBlackKing,
		bbBlackQueen:  b.bbBlackQueen,
		bbBlackRook:   b.bbBlackRook,
		bbBlackBishop: b.bbBlackBishop,
		bbBlackKnight: b.bbBlackKnight,
		bbBlackPawn:   b.bbBlackPawn,
	}
}

func (b *Board) isOccupied(sq Square) bool {
	return !b.emptySqs.Occupied(sq)
}

func (b *Board) piece(sq Square) *Piece {
	for _, p := range allPieces {
		bb := b.bbForPiece(p)
		if bb.Occupied(sq) {
			return p
		}
	}
	return nil
}

func (b *Board) hasSufficientMaterial() bool {
	// queen, rook, or pawn exist
	if (b.bbWhiteQueen | b.bbWhiteRook | b.bbWhitePawn |
		b.bbBlackQueen | b.bbBlackRook | b.bbBlackPawn) > 0 {
		return true
	}
	// if king is missing then it is a test
	if b.bbWhiteKing == 0 || b.bbBlackKing == 0 {
		return true
	}
	count := map[PieceType]int{}
	pieceMap := b.SquareMap()
	for _, p := range pieceMap {
		count[p.Type()]++
	}
	// 	king versus king
	if count[Bishop] == 0 && count[Knight] == 0 {
		return false
	}
	// king and bishop versus king
	if count[Bishop] == 1 && count[Knight] == 0 {
		return false
	}
	// king and knight versus king
	if count[Bishop] == 0 && count[Knight] == 1 {
		return false
	}
	// king and bishop(s) versus king and bishop(s) with the bishops on the same colour.
	if count[Knight] == 0 {
		whiteCount := 0
		blackCount := 0
		for sq, p := range pieceMap {
			if p.Type() == Bishop {
				switch sq.color() {
				case White:
					whiteCount++
				case Black:
					blackCount++
				}
			}
		}
		if whiteCount == 0 || blackCount == 0 {
			return false
		}
	}
	return true
}

func (b *Board) bbForPiece(p *Piece) bitboard {
	switch p {
	case WhiteKing:
		return b.bbWhiteKing
	case WhiteQueen:
		return b.bbWhiteQueen
	case WhiteRook:
		return b.bbWhiteRook
	case WhiteBishop:
		return b.bbWhiteBishop
	case WhiteKnight:
		return b.bbWhiteKnight
	case WhitePawn:
		return b.bbWhitePawn
	case BlackKing:
		return b.bbBlackKing
	case BlackQueen:
		return b.bbBlackQueen
	case BlackRook:
		return b.bbBlackRook
	case BlackBishop:
		return b.bbBlackBishop
	case BlackKnight:
		return b.bbBlackKnight
	case BlackPawn:
		return b.bbBlackPawn
	}
	return bitboard(0)
}

func (b *Board) setBBForPiece(p *Piece, bb bitboard) {
	switch p {
	case WhiteKing:
		b.bbWhiteKing = bb
	case WhiteQueen:
		b.bbWhiteQueen = bb
	case WhiteRook:
		b.bbWhiteRook = bb
	case WhiteBishop:
		b.bbWhiteBishop = bb
	case WhiteKnight:
		b.bbWhiteKnight = bb
	case WhitePawn:
		b.bbWhitePawn = bb
	case BlackKing:
		b.bbBlackKing = bb
	case BlackQueen:
		b.bbBlackQueen = bb
	case BlackRook:
		b.bbBlackRook = bb
	case BlackBishop:
		b.bbBlackBishop = bb
	case BlackKnight:
		b.bbBlackKnight = bb
	case BlackPawn:
		b.bbBlackPawn = bb
	default:
		panic("HERE")
	}
}

var (
	// file bitboards
	bbFileA bitboard
	bbFileB bitboard
	bbFileC bitboard
	bbFileD bitboard
	bbFileE bitboard
	bbFileF bitboard
	bbFileG bitboard
	bbFileH bitboard

	// rank bitboards
	bbRank1 bitboard
	bbRank2 bitboard
	bbRank3 bitboard
	bbRank4 bitboard
	bbRank5 bitboard
	bbRank6 bitboard
	bbRank7 bitboard
	bbRank8 bitboard

	bbFiles        [8]bitboard
	bbRanks        [8]bitboard
	bbSquares      [64]bitboard
	bbDiagonal     [64]bitboard
	bbAntiDiagonal [64]bitboard
)

func init() {
	bbFileA = bbForFile(fileA)
	bbFileB = bbForFile(fileB)
	bbFileC = bbForFile(fileC)
	bbFileD = bbForFile(fileD)
	bbFileE = bbForFile(fileE)
	bbFileF = bbForFile(fileF)
	bbFileG = bbForFile(fileG)
	bbFileH = bbForFile(fileH)

	bbRank1 = bbForRank(rank1)
	bbRank2 = bbForRank(rank2)
	bbRank3 = bbForRank(rank3)
	bbRank4 = bbForRank(rank4)
	bbRank5 = bbForRank(rank5)
	bbRank6 = bbForRank(rank6)
	bbRank7 = bbForRank(rank7)
	bbRank8 = bbForRank(rank8)

	bbFiles = [8]bitboard{bbFileA, bbFileB, bbFileC, bbFileD, bbFileE, bbFileF, bbFileG, bbFileH}
	bbRanks = [8]bitboard{bbRank1, bbRank2, bbRank3, bbRank4, bbRank5, bbRank6, bbRank7, bbRank8}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		sqr := Square(sq)
		bbSquares[sq] = bbRanks[sqr.rank()] & bbFiles[sqr.file()]
	}

	// init diagonal and anti-diagonal bitboards
	bbDiagonal = [64]bitboard{}
	bbAntiDiagonal = [64]bitboard{}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		sqr := Square(sq)
		bbDiagonal[sqr] = bbForDiagonal(sqr)
		bbAntiDiagonal[sqr] = bbForAntiDiagonal(sqr)
	}
}

func bbForFile(f file) bitboard {
	m := map[Square]bool{}
	var sq Square
	for ; sq < numOfSquaresInBoard; sq++ {
		if sq.file() == f {
			m[sq] = true
		}
	}
	return newBitboard(m)
}

func bbForRank(r rank) bitboard {
	m := map[Square]bool{}
	var sq Square
	for ; sq < numOfSquaresInBoard; sq++ {
		if sq.rank() == r {
			m[sq] = true
		}
	}
	return newBitboard(m)
}

func bbForDiagonal(sq Square) bitboard {
	v := int(sq.file()) - int(sq.rank())
	m := map[Square]bool{}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		sqr := Square(sq)
		if int(sqr.file())-int(sqr.rank()) == v {
			m[sqr] = true
		}
	}
	return newBitboard(m)
}

func bbForAntiDiagonal(sq Square) bitboard {
	v := int(sq.rank()) + int(sq.file())
	m := map[Square]bool{}
	for sq := 0; sq < numOfSquaresInBoard; sq++ {
		sqr := Square(sq)
		if int(sqr.rank())+int(sqr.file()) == v {
			m[sqr] = true
		}
	}
	return newBitboard(m)
}
