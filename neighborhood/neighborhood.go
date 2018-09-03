package neighborhood

type Neighborhood []bool

func (n Neighborhood) Equals(m Neighborhood) bool {
	if len(m) != 3 {
		return false
	}

	for i := range n {
		if n[i] != m[i] {
			return false
		}
	}

	return true
}
