package evaluators

import (
	"joshtompkins.com/elementary-automata/neighborhood"
)

type RuleOneTwentySixEvaluator struct{}

func NewRuleOneTwentySixEvaluator() *RuleOneTwentySixEvaluator {
	return &RuleOneTwentySixEvaluator{}
}

func (r *RuleOneTwentySixEvaluator) Evaluate(n neighborhood.Neighborhood) bool {
	switch {
	case n.Equals(neighborhood.Neighborhood{true, true, true}):
		return false
	case n.Equals(neighborhood.Neighborhood{false, false, false}):
		return false
	case n.Equals(neighborhood.Neighborhood{true, true, false}):
		return true
	case n.Equals(neighborhood.Neighborhood{true, false, true}):
		return true
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
