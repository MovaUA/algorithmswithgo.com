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

	for _, item := range list {
		sorted = insert(sorted, item)
	}

	copy(list, sorted)
}

func insert(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			// sorted[:i] + item + sorted[i:]
			sorted = append(sorted, 0)
			copy(sorted[i+1:], sorted[i:])
			sorted[i] = item

			return sorted
		}
	}
	return append(sorted, item)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
	sorted := make([]string, 0, len(list))

	for _, item := range list {
		sorted = insertString(sorted, item)
	}

	copy(list, sorted)
}
func insertString(sorted []string, item string) []string {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			sorted = append(sorted, "")
			copy(sorted[i+1:], sorted[i:])
			sorted[i] = item
			return sorted
		}
	}
	return append(sorted, item)
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
	sorted := make([]Person, 0, len(people))

	for _, person := range people {
		sorted = insertPerson(sorted, person)
	}

	copy(people, sorted)
}

func insertPerson(sorted []Person, person Person) []Person {
	for i, sortedPerson := range sorted {
		if greater(sortedPerson, person) {
			sorted = append(sorted, Person{})
			copy(sorted[i+1:], sorted[i:])
			sorted[i] = person
			return sorted
		}
	}
	return append(sorted, person)
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	for i := 1; i < list.Len(); i++ {
		insertInterface(list, i)
	}
}

func insertInterface(list sort.Interface, index int) {
	for i := index - 1; i >= 0; i-- {
		if list.Less(i+1, i) {
			list.Swap(i+1, i)
			continue
		}
		return
	}
}
