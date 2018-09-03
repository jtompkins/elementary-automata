package evaluators_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "joshtompkins.com/elementary-automata/evaluators"
	. "joshtompkins.com/elementary-automata/neighborhood"
)

var _ = Describe("ReverseEvaluator", func() {
	var (
		e = NewReverseEvaluator()
	)

	Describe("Evaluate", func() {
		Context("when the neighborhood are all 'off' cells", func() {
			It("returns true", func() {
				Expect(e.Evaluate(Neighborhood{false, false, false})).To(BeTrue())
			})
		})

		Context("when the neighborhood is anything else", func() {
			It("returns false", func() {
				Expect(e.Evaluate(Neighborhood{true, false, true})).To(BeFalse())
			})
		})
	})
})
