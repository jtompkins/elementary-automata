package generation

import (
	"math/rand"

	. "joshtompkins.com/elementary-automata/neighborhood"
)

type Generation []bool

func New(size int) Generation {
	return make(Generation, size)
}

func NewFromCells(cells ...bool) Generation {
	return append(make(Generation, 0), cells...)
}

func NewFromCenter(size int) Generation {
	g := make(Generation, size)
	g[size/2] = true

	return g
}

func NewFromRandom(size int) Generation {
	g := make(Generation, size)

	for i := range g {
		if rand.Intn(2) == 1 {
			g[i] = true
		}
	}

	return g
}

func (g Generation) NeighborhoodAtLocale(locale int) Neighborhood {
	var n Neighborhood

	switch {
	case locale == 0:
		n = Neighborhood{false}
		n = append(n, g[:locale+2]...)
	case locale == len(g)-1:
		n = make(Neighborhood, 2)
		copy(n, g[locale-1:])
		n = append(n, false)
	default:
		n = make(Neighborhood, 3)
		copy(n, g[locale-1:locale+2])
	}

	return n
}
