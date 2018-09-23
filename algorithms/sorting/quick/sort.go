package quick

func Sort(s []int) {
	sort(s, 0, len(s)-1)
}

func sort(s []int, left, right int) {
	if right < left {
		return
	}

	// We need to place our first pivot on the right place.
	// Function will return us new pivots position
	p := sortPartition(s, left, right)

	// After sorting main partition we need to split it.
	// Left and right side from the pivot will be sorted separately.
	// That is because any value on the left is smaller than any value on the right.
	sort(s, left, p-1)
	sort(s, p+1, right)
}

// sortPartition will put pivot on its final position.
// Also it will make sure that any value on the left from the pivot is smaller than any value on the right.
func sortPartition(s []int, first, last int) int {
	// Iterating over the specified piece
	for cur := first; cur < last; cur++ {
		// If current value is smaller than pivot, we need to put it on the left from the current cursor.
		if s[cur] <= s[last] {
			// It is okay to swap values in place. The main idea is to put smaller values on place
			s[cur], s[first] = s[first], s[cur]
			// Updating "possible" position for pivot
			first++
		}
	}

	// After finishing iteration we can put our pivot to the correct position
	s[last], s[first] = s[first], s[last]
	return first
}
