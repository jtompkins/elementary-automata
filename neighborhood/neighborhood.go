package neighborhood

type Neighborhood []bool

func (n Neighborhood) Equals(m Neighborhood) bool {
	if len(n) != len(m) {
		return false
	}

	result := true

	for i := range n {
		if n[i] != m[i] {
			return false
		}
	}

	return result
}
