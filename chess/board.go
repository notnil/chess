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
	bbFileA uint64
	bbFileB uint64
	bbFileC uint64
	bbFileD uint64
	bbFileE uint64
	bbFileF uint64
	bbFileG uint64
	bbFileH uint64

	// rank bitboards
	bbRank1 uint64
	bbRank2 uint64
	bbRank3 uint64
	bbRank4 uint64
	bbRank5 uint64
	bbRank6 uint64
	bbRank7 uint64
	bbRank8 uint64
)

func init() {
	fileBBFunc := func(sq, i int) bool {
		return (sq % numOfSquaresInRow) == i
	}
	bbFileA = initBB(int(fileA), fileBBFunc)
	bbFileB = initBB(int(fileB), fileBBFunc)
	bbFileC = initBB(int(fileC), fileBBFunc)
	bbFileD = initBB(int(fileD), fileBBFunc)
	bbFileE = initBB(int(fileE), fileBBFunc)
	bbFileF = initBB(int(fileF), fileBBFunc)
	bbFileG = initBB(int(fileG), fileBBFunc)
	bbFileH = initBB(int(fileH), fileBBFunc)

	rankBBFunc := func(sq, i int) bool {
		return (sq / numOfSquaresInRow) == i
	}
	bbRank1 = initBB(int(rank1), rankBBFunc)
	bbRank2 = initBB(int(rank2), rankBBFunc)
	bbRank3 = initBB(int(rank3), rankBBFunc)
	bbRank4 = initBB(int(rank4), rankBBFunc)
	bbRank5 = initBB(int(rank5), rankBBFunc)
	bbRank6 = initBB(int(rank6), rankBBFunc)
	bbRank7 = initBB(int(rank7), rankBBFunc)
	bbRank8 = initBB(int(rank8), rankBBFunc)
}

func initBB(i int, f func(sq, i int) bool) uint64 {
	s := ""
	for sq := 0; sq < numOfSquaresInRow; sq++ {
		if f(sq, i) {
			s += "1"
		} else {
			s += "0"
		}
	}
	bb, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return bb
}
