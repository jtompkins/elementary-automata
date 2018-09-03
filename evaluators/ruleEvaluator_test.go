package evaluators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"joshtompkins.com/elementary-automata/neighborhood"

	. "joshtompkins.com/elementary-automata/evaluators"
	. "joshtompkins.com/elementary-automata/rules"
)

var _ = Describe("RuleEvaluator", func() {
	var (
		r = Rule{
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
		}

		e = NewRuleEvaluator(r)
	)

	Describe("Evaluate", func() {
		It("evaluates the neighborhood based on the rule", func() {
			res := e.Evaluate(neighborhood.Neighborhood{false, false, false})
			Expect(res).To(BeTrue())
		})

		// TODO: add more tests with neighborhoods that match actual rules
	})
})
