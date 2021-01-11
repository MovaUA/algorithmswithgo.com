package module01

import (
	"fmt"
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	m := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}
	n := 0
	l := len(value)
	for i, v := range value {
		index, ok := m[v]
		if !ok {
			panic(fmt.Sprintf("Unknown digit %q", v))
		}
		n += index * int(math.Pow(float64(base), float64(l-i-1)))
	}
	return n
}
