package generation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "joshtompkins.com/elementary-automata/generation"
	. "joshtompkins.com/elementary-automata/neighborhood"
)

var _ = Describe("Generation", func() {
	var generation Generation

	Describe("NeighborhoodAtLocale", func() {
		BeforeEach(func() {
			generation = Generation{false, true, false}
		})

		It("returns a neighborhood with the locale and two surrounding cells", func() {
			n := generation.NeighborhoodAtLocale(1)

			Expect(len(n)).To(Equal(3))
			Expect(n).To(Equal(Neighborhood{false, true, false}))
		})

		Context("when getting a neighborhood for a locale on the left edge", func() {
			It("returns a neighborhood with a leading 'off' cell", func() {
				n := generation.NeighborhoodAtLocale(0)

				Expect(len(n)).To(Equal(3))
				Expect(n).To(Equal(Neighborhood{false, false, true}))
			})
		})

		Context("when getting a neighborhood for a locale on the right edge", func() {
			It("returns a neighborhood with a trailing 'off' cell", func() {
				n := generation.NeighborhoodAtLocale(2)

				Expect(len(n)).To(Equal(3))
				Expect(n).To(Equal(Neighborhood{true, false, false}))
			})
		})
	})
})
