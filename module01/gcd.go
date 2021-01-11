package module01

// GCD is the graitest common divider
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
