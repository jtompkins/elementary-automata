package main

import (
	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
)

func main() {

	a := automata.New(
		generation.New(10),
		evaluators.NewReverseEvaluator(),
		renderers.NewCliRenderer(),
	)

	a.Step()
	a.Step()
	a.Step()

	a.Render()
}
