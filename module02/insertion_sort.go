package module02

import "sort"

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	sorted := make([]int, 0, len(list))

	if len(list) > 0 {
		sorted = append(sorted, list[0])
	}

	for j := 1; j < len(list); j++ {
		v := list[j]

		if v >= sorted[len(sorted)-1] {
			sorted = append(sorted, v)
			continue
		}

		for i := 0; i < len(sorted); i++ {
			if sorted[i] >= v {
				sorted = append(sorted, 0)
				copy(sorted[i+1:], sorted[i:])
				sorted[i] = v
				break
			}
		}
	}

	for i, v := range sorted {
		list[i] = v
	}
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
}
