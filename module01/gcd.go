package module01

// GCD is the graitest common divider
func GCD(a, b int) int {
	a, b = minMax(a, b)
	for {
		if a == 0 {
			return b
		}
		if b == 0 {
			return a
		}
		a, b = minMax(b%a, a)
	}
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
