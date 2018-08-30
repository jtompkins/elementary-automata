package types

import (
	"fmt"
)

type Generation []bool

func NewGeneration(size int) Generation {
	return make(Generation, size)
}

func (g Generation) Print() {
	for _, cell := range g {
		if cell {
			fmt.Print("X")
		} else {
			fmt.Print("_")
		}
	}

	fmt.Println("")
}

type Neighborhood []bool

func NewNeighborhood(generation Generation, locale int) Neighborhood {
	var n Neighborhood

	switch {
	case locale == 0:
		n = Neighborhood{false}
		n = append(n, generation[:locale+2]...)
	case locale == len(generation)-1:
		n = make(Neighborhood, 2)
		copy(n, generation[locale-1:])
		n = append(n, false)
	default:
		n = make(Neighborhood, 3)
		copy(n, generation[locale-1:locale+2])
	}

	return n
}
