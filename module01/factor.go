package module01

// Factor takes in a list of primes and a number and factors that number with
// the provided primes.
//
// The returned numbers can be in any order; tests will sort them in increasing
// order to make checking easier.
//
// Bonus: Any remainder (aside from 1) that can't be factored will be treated as
// a prime in the results.
//
// Examples:
//
//   Factor([]int{2,3,5}, 30) // []int{2,3,5}
//   Factor([]int{2,3,5}, 28) // []int{2,2,7}
//   Factor([]int{2,3,5}, 720) // []int{2,2,2,2,3,3,5}
//
// Examples with remainders:
//
//   Factor([2,5], 30) // []int{2,5,3}
//   Factor([3,5], 720) // []int{3,3,5,16}
//   Factor([], 4) // []int{4}
//
func Factor(primes []int, number int) []int {
	result := make([]int, 0, len(primes))

	index := 0
	for {
		if index > len(primes)-1 {
			break
		}
		prime := primes[index]
		if prime > number {
			break
		}

		if rem := number % prime; rem == 0 {
			result = append(result, prime)
			number /= prime
			continue
		}

		index++
	}

	if number > 1 {
		result = append(result, number)
	}

	return result
}
