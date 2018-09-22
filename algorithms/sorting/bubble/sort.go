package bubble

func Sort(s []int) {
	sortTill := 0 // stores index of last sorted element. We should start from 0 and increase it on each new cycle
	for {
		left, right := len(s)-1, len(s) // Reseting positions of values to compare to the right end of sequence
		for {
			left--
			right--
			// Checking if we reached left end
			if left < sortTill {
				sortTill++
				break
			}

			// Compare values and swap if value on the left is bigger then on the right
			if s[right] < s[left] {
				s[right], s[left] = s[left], s[right]
			}
		}

		if len(s) < sortTill {
			break
		}
	}
}
