package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
	"joshtompkins.com/elementary-automata/rules"
)

var (
	size  = kingpin.Flag("size", "Size of each generation").Short('s').Default("1000").Int()
	gens  = kingpin.Flag("generations", "Number of generations to simulate").Short('g').Default("1000").Int()
	scale = kingpin.Flag("scale", "Size of each cell in the output in pixels").Default("1").Int()
	rule  = kingpin.Arg("rule", "Wolfram rule to use").Required().Int()
	file  = kingpin.Arg("file", "Path to output PNG").Required().String()
)

func main() {
	kingpin.Version("0.1")
	kingpin.Parse()

	r, _ := rules.Get(*rule)

	a := automata.New(
		generation.NewFromRandom(*size),
		evaluators.NewRuleEvaluator(r),
		renderers.NewPngRenderer(&renderers.RenderOptions{
			File:  *file,
			Scale: *scale,
		}),
	)

	log.Info("Evaluating rules")

	for i := 1; i <= *size; i++ {
		a.Step()
	}

	log.Info("Rendering")

	a.Render()

	log.Info("Done")
}
