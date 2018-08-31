package automata

import (
	"joshtompkins.com/elementary-automata/evaluators"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers"
)

type Automata struct {
	generations []generation.Generation
	evaluator   evaluators.Evaluator
	renderer    renderers.Renderer
}

func New(initial generation.Generation, e evaluators.Evaluator, r renderers.Renderer) *Automata {
	return &Automata{
		generations: []generation.Generation{initial},
		evaluator:   e,
		renderer:    r,
	}
}

func (a *Automata) Step() {
	lastGeneration := a.generations[len(a.generations)-1]
	nextGeneration := generation.New(10)

	for locale := range lastGeneration {
		nextGeneration[locale] = a.evaluator.Evaluate(lastGeneration.NeighborhoodAtLocale(locale))
	}

	a.generations = append(a.generations, nextGeneration)
}

func (a *Automata) Render() {
	a.renderer.Render(a.generations)
}
