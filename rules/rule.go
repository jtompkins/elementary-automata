package rules

import "errors"

type Rule [8]bool

var registeredRules = make(map[int]Rule)

func register(num int, r Rule) {
	registeredRules[num] = r
}

func Get(r int) (Rule, error) {
	rule, ok := registeredRules[r]

	if !ok {
		return Rule{}, errors.New("rule not found")
	}

	return rule, nil
}
