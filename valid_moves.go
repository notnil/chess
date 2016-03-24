package chess

type moveFunc func(*Position, bitboard, searchType) []*Move

var (
	moveFuncs []moveFunc
)

func init() {
	moveFuncs = []moveFunc{
		kingMoves, pawnMoves, knightMoves, rookMoves, bishopMoves, queenMoves,
	}
}

func (pos *Position) getValidMoves(s2BB bitboard, st searchType, allowChecks bool) []*Move {
	moves := []*Move{}
	for _, f := range moveFuncs {
		result := f(pos, s2BB, st)
		for _, m := range result {
			if addMove(pos, m, allowChecks) {
				moves = append(moves, m)
				if st == getFirst {
					return moves
				}
			}
		}
	}
	// have to seperate castle moves because of callback loop
	if !allowChecks {
		result := castleMoves(pos, s2BB, st)
		for _, m := range result {
			if addMove(pos, m, allowChecks) {
				moves = append(moves, m)
				if st == getFirst {
					return moves
				}
			}
		}
	}
	return moves
}

func addMove(pos *Position, m *Move, allowChecks bool) bool {
	if !allowChecks {
		pos.addTags(m)
		if m.HasTag(inCheck) {
			return false
		}
	}
	return true
}

type validMoveBB struct {
	bb    bitboard
	shift int
	msb   uint8
	lsb   uint8
}

func (pos *Position) addTags(m *Move) {
	p := pos.board.piece(m.s1)
	if pos.board.isOccupied(m.s2) {
		m.addTag(Capture)
	} else if m.s2 == pos.enPassantSquare && p.Type() == Pawn {
		m.addTag(EnPassant)
	}
	// determine if in check after move (makes move invalid)
	cp := pos.copy()
	cp.board.update(m)
	if cp.inCheck() {
		m.addTag(inCheck)
	}
	// determine if opponent in check after move
	cp.turn = cp.turn.Other()
	if cp.inCheck() {
		m.addTag(Check)
	}
}

func (pos *Position) inCheck() bool {
	kingSq := pos.board.whiteKingSq
	if pos.Turn() == Black {
		kingSq = pos.board.blackKingSq
	}
	// king should only be missing in tests / examples
	if kingSq == NoSquare {
		return false
	}
	// if no piece is on a attacking square, there can't be a check
	kingBB := bbSquares[kingSq]
	s2BB := pos.board.blackSqs
	if pos.Turn() == Black {
		s2BB = pos.board.whiteSqs
	}
	occ := ^pos.board.emptySqs
	// check queen attack vector
	bb := (diaAttack(occ, kingSq) | hvAttack(occ, kingSq)) & s2BB
	if bb != 0 {
		return pos.squaresAreAttacked(kingSq)
	}
	// check knight attack vector
	bb = ((kingBB & ^(bbRank7 | bbRank8) & ^bbFileH) >> 17) & s2BB
	bb = (((kingBB & ^(bbRank7 | bbRank8) & ^bbFileA) >> 15) & s2BB) | bb
	bb = ((kingBB & ^bbRank8 & ^(bbFileG | bbFileH) >> 10) & s2BB) | bb
	bb = ((kingBB & ^bbRank1 & ^(bbFileG | bbFileH) << 6) & s2BB) | bb
	bb = (((kingBB & ^(bbRank1 | bbRank2) & ^bbFileH) << 15) & s2BB) | bb
	bb = (((kingBB & ^(bbRank1 | bbRank2) & ^bbFileA) << 17) & s2BB) | bb
	bb = ((kingBB & ^bbRank8 & ^(bbFileA | bbFileB) >> 6) & s2BB) | bb
	bb = ((kingBB & ^bbRank1 & ^(bbFileA | bbFileB) << 10) & s2BB) | bb
	if bb != 0 {
		return pos.squaresAreAttacked(kingSq)
	}
	return pos.squaresAreAttacked(kingSq)
}

func (pos *Position) squaresAreAttacked(sqs ...Square) bool {
	// make s2bb for attacked squares
	var s2BB bitboard
	for _, sq := range sqs {
		s2BB = s2BB | bbSquares[sq]
	}
	// toggle turn to get other colors moves
	cp := pos.copy()
	cp.turn = cp.turn.Other()
	cp.enPassantSquare = NoSquare
	moves := cp.getValidMoves(s2BB, getFirst, true)
	return len(moves) > 0
}

func castleMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	moves := []*Move{}
	kingSide := pos.castleRights.CanCastle(pos.Turn(), KingSide)
	queenSide := pos.castleRights.CanCastle(pos.Turn(), QueenSide)
	// white king side
	if pos.turn == White && kingSide &&
		(^pos.board.emptySqs&(bbSquares[F1]|bbSquares[G1])) == 0 &&
		!pos.squaresAreAttacked(F1, G1) &&
		!pos.inCheck() {
		m := &Move{s1: E1, s2: G1}
		m.addTag(KingSideCastle)
		moves = append(moves, m)
	}
	// white queen side
	if pos.turn == White && queenSide &&
		(^pos.board.emptySqs&(bbSquares[B1]|bbSquares[C1]|bbSquares[D1])) == 0 &&
		!pos.squaresAreAttacked(C1, D1) &&
		!pos.inCheck() {
		m := &Move{s1: E1, s2: C1}
		m.addTag(QueenSideCastle)
		moves = append(moves, m)
	}
	// black king side
	if pos.turn == Black && kingSide &&
		(^pos.board.emptySqs&(bbSquares[F8]|bbSquares[G8])) == 0 &&
		!pos.squaresAreAttacked(F8, G8) &&
		!pos.inCheck() {
		m := &Move{s1: E8, s2: G8}
		m.addTag(KingSideCastle)
		moves = append(moves, m)
	}
	// black queen side
	if pos.turn == Black && queenSide &&
		(^pos.board.emptySqs&(bbSquares[B8]|bbSquares[C8]|bbSquares[D8])) == 0 &&
		!pos.squaresAreAttacked(C8, D8) &&
		!pos.inCheck() {
		m := &Move{s1: E8, s2: C8}
		m.addTag(QueenSideCastle)
		moves = append(moves, m)
	}
	return moves
}

func pawnMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	var bbs []*validMoveBB
	var bbEnPassant bitboard
	if pos.enPassantSquare != NoSquare {
		bbEnPassant = bbSquares[pos.enPassantSquare]
	}
	p := NoPiece
	if pos.Turn() == White {
		p = WhitePawn
		bbWhite := pos.board.bbWhitePawn
		if bbWhite == 0 {
			return []*Move{}
		}
		bbWhiteCapRight := ((bbWhite & ^bbFileH & ^bbRank8) >> 9) & ((pos.board.blackSqs & s2BB) | bbEnPassant)
		bbWhiteCapLeft := ((bbWhite & ^bbFileA & ^bbRank8) >> 7) & ((pos.board.blackSqs & s2BB) | bbEnPassant)
		bbWhiteUpOne := ((bbWhite & ^bbRank8) >> 8) & (pos.board.emptySqs & s2BB)
		bbWhiteUpTwo := ((bbWhiteUpOne & bbRank3) >> 8) & (pos.board.emptySqs & s2BB)
		bbs = []*validMoveBB{
			{bb: bbWhiteCapRight, shift: 9, lsb: 17, msb: 63},
			{bb: bbWhiteCapLeft, shift: 7, lsb: 16, msb: 62},
			{bb: bbWhiteUpOne, shift: 8, lsb: 16, msb: 63},
			{bb: bbWhiteUpTwo, shift: 16, lsb: 24, msb: 31},
		}
	} else {
		p = BlackPawn
		bbBlack := pos.board.bbBlackPawn
		if bbBlack == 0 {
			return []*Move{}
		}
		bbBlackCapRight := ((bbBlack & ^bbFileH & ^bbRank1) << 7) & ((pos.board.whiteSqs & s2BB) | bbEnPassant)
		bbBlackCapLeft := ((bbBlack & ^bbFileA & ^bbRank1) << 9) & ((pos.board.whiteSqs & s2BB) | bbEnPassant)
		bbBlackUpOne := ((bbBlack & ^bbRank1) << 8) & (pos.board.emptySqs & s2BB)
		bbBlackUpTwo := ((bbBlackUpOne & bbRank6) << 8) & (pos.board.emptySqs & s2BB)
		bbs = []*validMoveBB{
			{bb: bbBlackCapRight, shift: -7, lsb: 1, msb: 47},
			{bb: bbBlackCapLeft, shift: -9, lsb: 0, msb: 46},
			{bb: bbBlackUpOne, shift: -8, lsb: 0, msb: 47},
			{bb: bbBlackUpTwo, shift: -16, lsb: 32, msb: 39},
		}
	}
	return steppingMoves(p, bbs, st)
}

func knightMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	bb := pos.board.bbWhiteKnight
	p := WhiteKnight
	if pos.Turn() == Black {
		bb = pos.board.bbBlackKnight
		p = BlackKnight
	}
	if bb == 0 {
		return []*Move{}
	}
	bbUpRight := ((bb & ^(bbRank7 | bbRank8) & ^bbFileH) >> 17) & s2BB
	bbUpLeft := ((bb & ^(bbRank7 | bbRank8) & ^bbFileA) >> 15) & s2BB
	bbRightUp := (bb & ^bbRank8 & ^(bbFileG | bbFileH) >> 10) & s2BB
	bbRightDown := (bb & ^bbRank1 & ^(bbFileG | bbFileH) << 6) & s2BB
	bbDownRight := ((bb & ^(bbRank1 | bbRank2) & ^bbFileH) << 15) & s2BB
	bbDownLeft := ((bb & ^(bbRank1 | bbRank2) & ^bbFileA) << 17) & s2BB
	bbLeftUp := (bb & ^bbRank8 & ^(bbFileA | bbFileB) >> 6) & s2BB
	bbLeftDown := (bb & ^bbRank1 & ^(bbFileA | bbFileB) << 10) & s2BB
	bbs := []*validMoveBB{
		{bb: bbUpRight, shift: 17, lsb: 17, msb: 63},
		{bb: bbUpLeft, shift: 15, lsb: 15, msb: 63},
		{bb: bbRightUp, shift: 10, lsb: 10, msb: 63},
		{bb: bbRightDown, shift: -6, lsb: 0, msb: 57},
		{bb: bbDownRight, shift: -15, lsb: 0, msb: 48},
		{bb: bbDownLeft, shift: -17, lsb: 0, msb: 46},
		{bb: bbLeftUp, shift: 6, lsb: 6, msb: 63},
		{bb: bbLeftDown, shift: -10, lsb: 0, msb: 53},
	}
	return steppingMoves(p, bbs, st)
}

func kingMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	bb := pos.board.bbWhiteKing
	p := WhiteKing
	if pos.Turn() == Black {
		bb = pos.board.bbBlackKing
		p = BlackKing
	}
	if bb == 0 {
		return []*Move{}
	}
	bbUp := ((bb & ^bbRank8) >> 8) & s2BB
	bbUpRight := ((bb & ^bbRank8 & ^bbFileH) >> 9) & s2BB
	bbRight := ((bb & ^bbFileH) >> 1) & s2BB
	bbDownRight := ((bb & ^bbRank1 & ^bbFileH) << 7) & s2BB
	bbDown := ((bb & ^bbRank1) << 8) & s2BB
	bbDownLeft := ((bb & ^bbRank1 & ^bbFileA) << 9) & s2BB
	bbLeft := ((bb & ^bbFileA) << 1) & s2BB
	bbLeftUp := ((bb & ^bbRank8 & ^bbFileA) >> 7) & s2BB
	bbs := []*validMoveBB{
		{bb: bbUp, shift: 8, lsb: 8, msb: 63},
		{bb: bbUpRight, shift: 9, lsb: 9, msb: 63},
		{bb: bbRight, shift: 1, lsb: 1, msb: 63},
		{bb: bbDownRight, shift: -7, lsb: 0, msb: 56},
		{bb: bbDown, shift: -8, lsb: 0, msb: 55},
		{bb: bbDownLeft, shift: -9, lsb: 0, msb: 54},
		{bb: bbLeft, shift: -1, lsb: 0, msb: 62},
		{bb: bbLeftUp, shift: 7, lsb: 7, msb: 63},
	}
	// todo use predetermined king square
	return steppingMoves(p, bbs, st)
}

type slidingBB struct {
	s1 Square
	bb bitboard
}

var (
	rookFunc = func(pos *Position, sq Square, validBB bitboard) *slidingBB {
		bb := hvAttack(^pos.board.emptySqs, Square(sq)) & validBB
		return &slidingBB{bb: bb, s1: Square(sq)}
	}

	bishopFunc = func(pos *Position, sq Square, validBB bitboard) *slidingBB {
		bb := diaAttack(^pos.board.emptySqs, Square(sq)) & validBB
		return &slidingBB{bb: bb, s1: Square(sq)}
	}

	queenFunc = func(pos *Position, sq Square, validBB bitboard) *slidingBB {
		occ := ^pos.board.emptySqs
		bb := (diaAttack(occ, Square(sq)) | hvAttack(occ, Square(sq))) & validBB
		return &slidingBB{bb: bb, s1: Square(sq)}
	}
)

func queenMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	return pos.slidingMoves(Queen, s2BB, st, queenFunc)
}

func rookMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	return pos.slidingMoves(Rook, s2BB, st, rookFunc)
}

func bishopMoves(pos *Position, s2BB bitboard, st searchType) []*Move {
	return pos.slidingMoves(Bishop, s2BB, st, bishopFunc)
}

func (pos *Position) slidingMoves(pt PieceType, s2BB bitboard, st searchType, f func(*Position, Square, bitboard) *slidingBB) []*Move {
	moves := []*Move{}
	p := getPiece(pt, pos.Turn())
	bb := pos.board.bbForPiece(p)
	if bb == 0 {
		return moves
	}
	bbs := []*slidingBB{}
	for sq := 0; sq < 64; sq++ {
		if bb.Occupied(Square(sq)) {
			bbs = append(bbs, f(pos, Square(sq), s2BB))
		}
	}
	for _, sBB := range bbs {
		for sq := 0; sq < 64; sq++ {
			if sBB.bb.Occupied(Square(sq)) {
				moves = append(moves, &Move{s1: sBB.s1, s2: Square(sq)})
				if st == getFirst {
					return moves
				}
			}
		}
	}
	return moves
}

func diaAttack(occupied bitboard, sq Square) bitboard {
	pos := bbSquares[sq]
	dMask := bbDiagonal[sq]
	adMask := bbAntiDiagonal[sq]
	return linearAttack(occupied, pos, dMask) | linearAttack(occupied, pos, adMask)
}

func hvAttack(occupied bitboard, sq Square) bitboard {
	pos := bbSquares[sq]
	rankMask := bbRanks[Square(sq).Rank()]
	fileMask := bbFiles[Square(sq).File()]
	return linearAttack(occupied, pos, rankMask) | linearAttack(occupied, pos, fileMask)
}

func linearAttack(occupied, pos, mask bitboard) bitboard {
	oInMask := occupied & mask
	return ((oInMask - 2*pos) ^ (oInMask.Reverse() - 2*pos.Reverse()).Reverse()) & mask
}

type searchType uint8

const (
	getAll searchType = iota
	getFirst
)

func steppingMoves(p Piece, bbs []*validMoveBB, st searchType) []*Move {
	moves := []*Move{}
	for _, validBB := range bbs {
		if validBB.bb == 0 {
			continue
		}
		// TODO reduce number of squares by range of LSB to MSB per bitboard
		for sq := validBB.lsb; sq <= validBB.msb; sq++ {
			if validBB.bb.Occupied(Square(sq)) {
				s1 := Square(int(sq) - validBB.shift)
				s2 := Square(sq)
				if (p == WhitePawn && s2.Rank() == Rank8) || (p == BlackPawn && s2.Rank() == Rank1) {
					qm := &Move{s1: s1, s2: s2, promo: Queen}
					rm := &Move{s1: s1, s2: s2, promo: Rook}
					bm := &Move{s1: s1, s2: s2, promo: Bishop}
					nm := &Move{s1: s1, s2: s2, promo: Knight}
					moves = append(moves, qm, rm, bm, nm)
					if st == getFirst {
						return moves
					}
				} else {
					moves = append(moves, &Move{s1: s1, s2: s2})
					if st == getFirst {
						return moves
					}
				}
			}
		}
	}
	return moves
}
