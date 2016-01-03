package chess

import (
	"strconv"
	"strings"
)

type Board struct {
	bbs      map[*Piece]bitboard
	whiteSqs bitboard
	blackSqs bitboard
	emptySqs bitboard
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
	bbs := map[*Piece]bitboard{}
	for _, p1 := range allPieces {
		bm := map[Square]bool{}
		for sq, p2 := range m {
			if p1 == p2 {
				bm[sq] = true
			}
		}
		bb := newBitboard(bm)
		bbs[p1] = bb
	}
	whiteSqs := bbs[WhiteKing] | bbs[WhiteQueen] | bbs[WhiteRook] | bbs[WhiteBishop] | bbs[WhiteKnight] | bbs[WhitePawn]
	blackSqs := bbs[BlackKing] | bbs[BlackQueen] | bbs[BlackRook] | bbs[BlackBishop] | bbs[BlackKnight] | bbs[BlackPawn]
	emptySqs := ^(whiteSqs | blackSqs)
	return &Board{
		bbs:      bbs,
		whiteSqs: whiteSqs,
		blackSqs: blackSqs,
		emptySqs: emptySqs,
	}
}

func (b *Board) isOccupied(sq Square) bool {
	return !b.emptySqs.Occupied(sq)
}

func (b *Board) piece(sq Square) *Piece {
	for p, bb := range b.bbs {
		if bb.Occupied(sq) {
			return p
		}
	}
	return nil
}

func (b *Board) squareMap() map[Square]*Piece {
	m := map[Square]*Piece{}
	for _, p := range allPieces {
		bbMap := b.bbs[p].Mapping()
		for sq := range bbMap {
			m[sq] = p
		}
	}
	return m
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
