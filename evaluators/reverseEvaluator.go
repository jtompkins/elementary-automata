package evaluators

import "joshtompkins.com/elementary-automata/neighborhood"

type ReverseEvaluator struct{}

func NewReverseEvaluator() *ReverseEvaluator {
	return &ReverseEvaluator{}
}

func (e ReverseEvaluator) Evaluate(n neighborhood.Neighborhood) bool {
	return n.Equals(neighborhood.Neighborhood{false, false, false})
}
