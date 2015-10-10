package chess

func rankStep(c Color) int {
	if c == White {
		return 1
	}
	return -1
}

func backRank(c Color) Rank {
	if c == White {
		return R1
	}
	return R8
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func sign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	}
	return 0
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
