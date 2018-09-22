package insertion

func Sort(s []int) {
	for i := 1; i < len(s); i++ {
		pos := -1
		for j := i - 1; j >= 0; j-- { // Checking previously "sorted" values to find place for current one
			if s[i] < s[j] {
				pos = j
			}
		}

		if pos < 0 { // If new position was not found - element is on right place
			continue
		}

		val := s[i]                                            // save element before delete from slice
		s = append(s[:i], s[i+1:]...)                          // delete element from old position
		s = append(s[:pos], append([]int{val}, s[pos:]...)...) // put element on new position
	}
}
