package chess

import (
	"math/rand"
)

var pieceSquareZobristC [12][64]uint64
var castleRightsZobristC [4]uint64
var enPassantZobristC [16]uint64
var whiteToMoveZobristC uint64

func initZobrist() {
	whiteToMoveZobristC = rand.Uint64()
	for i := 0; i < 12; i++ {
		for j := 0; j < 64; j++ {
			pieceSquareZobristC[i][j] = rand.Uint64()
		}
	}
	for i := 0; i < 4; i++ {
		castleRightsZobristC[i] = rand.Uint64()
	}
	for i := 0; i < 16; i++ {
		enPassantZobristC[i] = rand.Uint64()
	}
}

func generateZobristHash(pos *Position) uint64 {
	var hash uint64 = 0
	turn := pos.turn

	/* Turn */
	if turn == White {
		hash ^= whiteToMoveZobristC
	}

	/* Castle */
	cc := pos.castleRights.CanCastle
	if cc(White, KingSide) {
		hash ^= castleRightsZobristC[0]
	}
	if cc(White, QueenSide) {
		hash ^= castleRightsZobristC[1]
	}
	if cc(Black, KingSide) {
		hash ^= castleRightsZobristC[2]
	}
	if cc(Black, QueenSide) {
		hash ^= castleRightsZobristC[3]
	}

	/* En passant */
	enPassant := pos.enPassantSquare
	if enPassant != NoSquare {
		if turn == Black {
			/* Next mov Black -> Current pos White -> White en passant square */
			hash ^= enPassantZobristC[enPassant-16]
		} else {
			/* Next mov White -> Current pos Black -> Black en passant square */
			hash ^= enPassantZobristC[enPassant-40+8]
		}
	}

	/* Board */
	piece := pos.board.Piece
	for sq := 0; sq < 64; sq++ {
		p := piece(Square(sq))
		if p != NoPiece {
			hash ^= pieceSquareZobristC[int8(p)-1][sq]
		}
	}

	return hash
}

func updateZobristHash(pos *Position, mov *Move) uint64 {
	hash := pos.hash
	turn := pos.turn
	srcSq := mov.s1
	dstSq := mov.s2
	piece := pos.board.Piece
	hasTag := mov.HasTag
	posEnPassantSquare := pos.enPassantSquare
	movEnPassantSquare := pos.updateEnPassantSquare(mov)

	/* Switch turn */
	hash ^= whiteToMoveZobristC

	/* Remove our piece in S1 */
	ourP := piece(srcSq)
	hash ^= pieceSquareZobristC[int8(ourP)-1][srcSq]

	/* Add our promoted piece in S2 */
	var ourPromoP Piece
	promo := mov.promo
	if promo != NoPieceType {
		ourPromoP = getPiece(promo, turn)
	} else {
		ourPromoP = ourP
	}
	hash ^= pieceSquareZobristC[int8(ourPromoP)-1][dstSq]

	/* Capture */
	if hasTag(Capture) {
		/* Remove captured piece */
		hash ^= pieceSquareZobristC[int8(piece(dstSq))-1][dstSq]
	}

	if oldCR, newCR := pos.castleRights, pos.updateCastleRights(mov); newCR != oldCR {
		/* Remove old castle rights */
		oldCanCastle := oldCR.CanCastle
		if oldCanCastle(White, KingSide) {
			hash ^= castleRightsZobristC[0]
		}
		if oldCanCastle(White, QueenSide) {
			hash ^= castleRightsZobristC[1]
		}
		if oldCanCastle(Black, KingSide) {
			hash ^= castleRightsZobristC[2]
		}
		if oldCanCastle(Black, QueenSide) {
			hash ^= castleRightsZobristC[3]
		}
		/* Add new castle rights */
		newCanCastle := newCR.CanCastle
		if newCanCastle(White, KingSide) {
			hash ^= castleRightsZobristC[0]
		}
		if newCanCastle(White, QueenSide) {
			hash ^= castleRightsZobristC[1]
		}
		if newCanCastle(Black, KingSide) {
			hash ^= castleRightsZobristC[2]
		}
		if newCanCastle(Black, QueenSide) {
			hash ^= castleRightsZobristC[3]
		}
	}

	/* Castle */
	hasKingSideCastle := hasTag(KingSideCastle)
	hasQueenSideCastle := hasTag(QueenSideCastle)
	if turn == White {
		rookBoard := pieceSquareZobristC[int8(WhiteRook)-1]
		if hasKingSideCastle {
			/* Remove rook in H1 */
			hash ^= rookBoard[H1]
			/* Add rook in F1 */
			hash ^= rookBoard[F1]
		} else if hasQueenSideCastle {
			/* Remove rook in A1 */
			hash ^= rookBoard[A1]
			/* Add rook in D1 */
			hash ^= rookBoard[D1]
		}
	} else {
		rookBoard := pieceSquareZobristC[int8(BlackRook)-1]
		if hasKingSideCastle {
			/* Remove rook in H8 */
			hash ^= rookBoard[H8]
			/* Add rook in F8 */
			hash ^= rookBoard[F8]
		} else if hasQueenSideCastle {
			/* Remove rook in A8 */
			hash ^= rookBoard[A8]
			/* Add rook in D8 */
			hash ^= rookBoard[D8]
		}
	}

	/* Remove old en passant square*/
	if posEnPassantSquare != NoSquare {
		if turn == White {
			/* White mov -> Black old en passant square */
			hash ^= enPassantZobristC[posEnPassantSquare-40+8]
		} else {
			/* Black mov -> White old en passant square */
			hash ^= enPassantZobristC[posEnPassantSquare-16]
		}
	}

	/* Add new en passant square */
	if movEnPassantSquare != NoSquare {
		if turn == Black {
			/* Black mov -> new Black en passant square */
			hash ^= enPassantZobristC[movEnPassantSquare-40+8]
		} else {
			/* White mov -> new White en passant square */
			hash ^= enPassantZobristC[movEnPassantSquare-16]
		}
	}

	/* En passant */
	if hasTag(EnPassant) {
		if turn == White {
			/* White mov ->
			Remove black pawn in same file as en passant square but previous rank */
			hash ^= pieceSquareZobristC[int8(BlackPawn)-1][posEnPassantSquare-8]
		} else {
			/* Black mov ->
			Remove white pawn in same file as en passant square but next rank */
			hash ^= pieceSquareZobristC[int8(WhitePawn)-1][posEnPassantSquare+8]
		}
	}

	return hash
}

func init() {
	initZobrist()
}
