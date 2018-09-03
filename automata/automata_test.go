package automata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"joshtompkins.com/elementary-automata/neighborhood"

	. "joshtompkins.com/elementary-automata/automata"
	"joshtompkins.com/elementary-automata/evaluators/evaluatorsfakes"
	"joshtompkins.com/elementary-automata/generation"
	"joshtompkins.com/elementary-automata/renderers/renderersfakes"
)

var _ = Describe("Automata", func() {
	var (
		initialGeneration generation.Generation
		fakeRenderer      *renderersfakes.FakeRenderer
		fakeEvaluator     *evaluatorsfakes.FakeEvaluator
		a                 *Automata
	)

	BeforeEach(func() {
		initialGeneration = generation.Generation{true, false, true}
		fakeEvaluator = new(evaluatorsfakes.FakeEvaluator)
		fakeRenderer = new(renderersfakes.FakeRenderer)
	})

	Describe("Step", func() {
		BeforeEach(func() {
			a = New(initialGeneration, fakeEvaluator, fakeRenderer)
			a.Step()
		})

		It("evaluates the neighborhoods in the generation", func() {
			Expect(fakeEvaluator.EvaluateCallCount()).To(Equal(3))
		})

		It("passes the neighborhoods from the previous generation to the evaluator", func() {
			n := fakeEvaluator.EvaluateArgsForCall(0)
			Expect(n.Equals(neighborhood.Neighborhood{false, true, false})).To(BeTrue())

			n = fakeEvaluator.EvaluateArgsForCall(1)
			Expect(n.Equals(neighborhood.Neighborhood{true, false, true})).To(BeTrue())

			n = fakeEvaluator.EvaluateArgsForCall(2)
			Expect(n.Equals(neighborhood.Neighborhood{false, true, false})).To(BeTrue())
		})
	})

	Describe("Render", func() {
		BeforeEach(func() {
			a = New(initialGeneration, fakeEvaluator, fakeRenderer)
			a.Render()
		})

		It("delegates rendering to the renderer", func() {
			Expect(fakeRenderer.RenderCallCount()).To(Equal(1))
		})

		It("passes the generations to the renderer", func() {
			gens := fakeRenderer.RenderArgsForCall(0)
			Expect(gens[0]).To(Equal(initialGeneration))
		})
	})
})
