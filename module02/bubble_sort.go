package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for sweepNum := 0; sweepNum < len(list)-1; sweepNum++ {
		swapped := false
		for i := 0; i < len(list)-1-sweepNum; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	for sweepNum := 0; sweepNum < len(list)-1; sweepNum++ {
		swapped := false
		for i := 0; i < len(list)-1-sweepNum; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	for sweepNum := 0; sweepNum < len(people)-1; sweepNum++ {
		swapped := false
		for i := 0; i < len(people)-1-sweepNum; i++ {
			if greater(people[i], people[i+1]) {
				people[i], people[i+1] = people[i+1], people[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
func greater(p1, p2 Person) bool {
	if p1.Age > p2.Age {
		return true
	}
	if p1.Age < p2.Age {
		return false
	}
	if p1.LastName > p2.LastName {
		return true
	}
	if p1.LastName < p2.LastName {
		return false
	}
	if p1.FirstName > p2.FirstName {
		return true
	}
	// if p1.FirstName<p2.FirstName{
	// 	return false
	// }
	return false
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
}
