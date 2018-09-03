package evaluators

import "joshtompkins.com/elementary-automata/neighborhood"

//go:generate counterfeiter . Evaluator

type Evaluator interface {
	Evaluate(neighborhood.Neighborhood) bool
}
