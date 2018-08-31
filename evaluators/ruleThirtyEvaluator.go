package evaluators

import (
	"joshtompkins.com/elementary-automata/neighborhood"
)

type RuleThirtyEvaluator struct{}

func NewRuleThirtyEvaluator() *RuleThirtyEvaluator {
	return &RuleThirtyEvaluator{}
}

func (r *RuleThirtyEvaluator) Evaluate(n neighborhood.Neighborhood) bool {
	switch {
	case n.Equals(neighborhood.Neighborhood{true, true, true}):
		return false
	case n.Equals(neighborhood.Neighborhood{true, true, false}):
		return false
	case n.Equals(neighborhood.Neighborhood{true, false, true}):
		return false
	case n.Equals(neighborhood.Neighborhood{false, false, false}):
		return false
	case n.Equals(neighborhood.Neighborhood{true, false, false}):
		return true
	case n.Equals(neighborhood.Neighborhood{false, true, true}):
		return true
	case n.Equals(neighborhood.Neighborhood{false, true, false}):
		return true
	case n.Equals(neighborhood.Neighborhood{false, false, true}):
		return true
	}

	return false
}
