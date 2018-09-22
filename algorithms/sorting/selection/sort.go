package selection

func Sort(s []int) {
	smallest := 0
	index := 0
	for pos := 0; pos < len(s); pos++ {
		for i := pos; i < len(s); i++ {
			if i == pos {
				smallest = s[i]
				index = i
				continue
			}

			if s[i] < smallest {
				smallest = s[i]
				index = i
			}
		}

		s[pos], s[index] = s[index], s[pos]
	}
}
