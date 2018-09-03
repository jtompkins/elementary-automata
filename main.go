package main

import (
	log "github.com/sirupsen/logrus"
	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
	"joshtompkins.com/elementary-automata/rules"
)

func main() {
	a := automata.New(
		generation.NewFromCenter(1000),
		evaluators.NewRuleEvaluator(rules.Thirty),
		renderers.NewPngRenderer(&renderers.RenderOptions{
			File:  "test.png",
			Scale: 1,
		}),
	)

	log.SetLevel(log.DebugLevel)

	log.Info("Evaluating rules")

	for i := 1; i < 1001; i++ {
		a.Step()
	}

	log.Info("Rendering")

	a.Render()

	log.Info("Done")
}
