package main

import (
	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
)

func main() {
	a := automata.New(
		generation.NewWithCenterCellOn(25),
		evaluators.NewRuleOneTwentySixEvaluator(),
		renderers.NewCliRenderer(),
	)

	for i := 1; i < 25; i++ {
		a.Step()
	}

	a.Render()
}
