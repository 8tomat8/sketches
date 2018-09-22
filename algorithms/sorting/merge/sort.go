package merge

func Sort(s []int) {
	if len(s) == 1 {
		return
	}

	res := make([]int, len(s))
	copy(res, s)
	left, right := res[:len(res)/2], res[len(res)/2:] // split sequence on 2 parts to sort them separately
	if len(res) > 2 {
		Sort(left)
		Sort(right)
	}

	// When both parts are sorted, we need to merge them and return them back
	for i, v := range merge(left, right) {
		s[i] = v
	}
}

func merge(a, b []int) []int {
	c := make([]int, 0, len(a)+len(b))
	bi := 0
	ai := 0
	for ai < len(a) {
		if len(b) == bi {
			break
		}
		for bi < len(b) {
			if a[ai] < b[bi] {
				c = append(c, a[ai])
				ai++
				break // Need to pick up next av value
			} else {
				c = append(c, b[bi])
				bi++
			}
		}
	}

	// Checking if any values left and appending them
	if len(b) != bi {
		c = append(c, b[bi:]...)
	} else if len(a) != ai {
		c = append(c, a[ai:]...)
	}
	return c
}
