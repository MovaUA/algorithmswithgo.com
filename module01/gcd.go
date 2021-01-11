package module01

// GCD is the graitest common divider
func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}

	min, max := minMax(a, b)

	for {
		r := max % min
		if r == 0 {
			return min
		}

		min, max = minMax(min, r)
		if min == 1 {
			return 1
		}
	}
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
