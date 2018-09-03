package neighborhood_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "joshtompkins.com/elementary-automata/neighborhood"
)

var _ = Describe("Neighborhood", func() {
	var (
		left Neighborhood = Neighborhood{true, true, true}
	)

	Describe("Equals", func() {
		Context("when the neighborhoods are equivalent", func() {
			It("returns true", func() {
				Expect(left.Equals(Neighborhood{true, true, true})).To(BeTrue())
			})
		})

		Context("when the neighborhoods are not equivalent", func() {
			It("returns false", func() {
				Expect(left.Equals(Neighborhood{false, false, false})).To(BeFalse())
			})
		})

		Context("when the neighborhood isn't the right length", func() {
			It("returns false", func() {
				Expect(left.Equals(Neighborhood{true, true, true, true})).To(BeFalse())
			})
		})
	})
})
