package generation

import (
	"math/rand"
	"time"

	. "joshtompkins.com/elementary-automata/neighborhood"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Generation []bool

func New(size int) Generation {
	return make(Generation, size, size)
}

func NewFromCenter(size int) Generation {
	g := make(Generation, size, size)
	g[size/2] = true

	return g
}

func NewFromRandom(size int) Generation {
	g := make(Generation, size, size)

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
		n = append(Neighborhood{false}, g[:locale+2]...)
	case locale == len(g)-1:
		n = make(Neighborhood, 2, 3)
		copy(n, g[locale-1:])
		n = append(n, false)
	default:
		n = make(Neighborhood, 3, 3)
		copy(n, g[locale-1:locale+2])
	}

	return n
}
