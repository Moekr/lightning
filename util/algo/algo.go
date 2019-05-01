package algo

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Inc(a int) int {
	return a + 1
}

func Dec(a int) int {
	return a - 1
}
