package evaluators

import (
	"errors"

	"joshtompkins.com/elementary-automata/neighborhood"
	"joshtompkins.com/elementary-automata/rules"
)

// these are the eight possible elementary automata neighborhoods,
// in the "canonical" order. We expect each "Rule" to match the order
// as listed here
var neighborhoods = []neighborhood.Neighborhood{
	neighborhood.Neighborhood{true, true, true},
	neighborhood.Neighborhood{true, true, false},
	neighborhood.Neighborhood{true, false, true},
	neighborhood.Neighborhood{true, false, false},
	neighborhood.Neighborhood{false, true, true},
	neighborhood.Neighborhood{false, true, false},
	neighborhood.Neighborhood{false, false, true},
	neighborhood.Neighborhood{false, false, false},
}

type RuleEvaluator struct {
	rule rules.Rule
}

func NewRuleEvaluator(r rules.Rule) *RuleEvaluator {
	return &RuleEvaluator{rule: r}
}

func (e *RuleEvaluator) Evaluate(n neighborhood.Neighborhood) bool {
	// find the matching neighborhood
	// drop the error case since all possible neighborhoods are in the slice
	match, _ := e.findMatchingNeighborhood(n)

	return e.rule[match]
}

func (e *RuleEvaluator) findMatchingNeighborhood(n neighborhood.Neighborhood) (int, error) {
	for i := range neighborhoods {
		if neighborhoods[i].Equals(n) {
			return i, nil
		}
	}

	return 0, errors.New("Invalid neighborhood")
}
