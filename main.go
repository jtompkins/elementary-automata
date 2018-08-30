package main

import (
	"fmt"

	"joshtompkins.com/elementary-automata/types"
)

func main() {
	fmt.Println("1-dimensional automata are cool!")

	generations := make([]types.Generation, 0)
	generations = append(generations, types.NewGeneration(10))

	lastGeneration := generations[len(generations)-1]

	lastGeneration.Print()

	nextGeneration := types.NewGeneration(10)

	for locale := range lastGeneration {
		nextGeneration[locale] = evaluate(types.NewNeighborhood(lastGeneration, locale))
	}

	nextGeneration.Print()

	generations = append(generations, nextGeneration)
}

func evaluate(neighborhood types.Neighborhood) bool {
	return true
}
