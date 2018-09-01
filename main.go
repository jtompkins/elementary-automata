package main

import (
	"fmt"

	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
)

func main() {
	a := automata.New(
		generation.NewFromCenter(100),
		evaluators.NewRuleThirtyEvaluator(),
		renderers.NewPngRenderer(&renderers.RenderOptions{
			File:  "test.png",
			Scale: 10,
		}),
	)

	fmt.Println("Evaluting rules...")

	for i := 1; i < 101; i++ {
		a.Step()
	}

	fmt.Println("Rendering...")

	a.Render()

	fmt.Println("Done!")
}
