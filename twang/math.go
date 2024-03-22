package twang

func inRange(x, a, b int) bool {
	return x >= a && x <= b
}

func minMax(x, a, b int) int {
	return min(max(x, a), b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
