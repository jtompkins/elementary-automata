package evaluators

import "joshtompkins.com/elementary-automata/neighborhood"

type Evaluator interface {
	Evaluate(neighborhood.Neighborhood) bool
}
