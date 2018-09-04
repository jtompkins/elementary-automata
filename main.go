package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
	"joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
	"joshtompkins.com/elementary-automata/rules"
)

var (
	size     = kingpin.Flag("size", "Size of each generation").Short('s').Default("1000").Int()
	gens     = kingpin.Flag("generations", "Number of generations to simulate").Short('g').Default("1000").Int()
	scale    = kingpin.Flag("scale", "Size of each cell in the output in pixels").Default("1").Int()
	centered = kingpin.Flag("center", "Center the initial generation").Short('c').Bool()
	rule     = kingpin.Arg("rule", "Wolfram rule to use").Required().Int()
	file     = kingpin.Arg("file", "Path to output PNG").Required().String()
)

func main() {
	kingpin.Version("0.1")
	kingpin.Parse()

	fmt.Print("\n🤖 Reading configuration\t")

	r, _ := rules.Get(*rule)

	var initial generation.Generation

	if *centered {
		initial = generation.NewFromCenter(*size)
	} else {
		initial = generation.NewFromRandom(*size)
	}

	fmt.Println("✅")

	fmt.Print("🤖 Spinning up an automata\t")

	a := automata.New(
		initial,
		evaluators.NewRuleEvaluator(r),
		renderers.NewPngRenderer(&renderers.RenderOptions{
			File:  *file,
			Scale: *scale,
		}),
	)

	fmt.Println("✅")

	fmt.Print("🤖 Evaluating rules\t\t")

	for i := 1; i <= *size; i++ {
		a.Step()
	}

	fmt.Println("✅")

	fmt.Print("🤖 Rendering image\t\t")

	a.Render()

	fmt.Println("✅\n")
}
