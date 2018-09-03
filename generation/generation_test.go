package generation_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "joshtompkins.com/elementary-automata/generation"
	. "joshtompkins.com/elementary-automata/neighborhood"
)

var _ = Describe("Generation", func() {
	var generation Generation

	Describe("New", func() {
		BeforeEach(func() {
			generation = New(3)
		})

		It("returns a generation with correct length and capacity", func() {
			Expect(len(generation)).To(Equal(3))
			Expect(cap(generation)).To(Equal(3))
		})
	})

	Describe("NewFromCenter", func() {
		BeforeEach(func() {
			generation = NewFromCenter(4)
		})

		It("returns a generation with correct length and capacity", func() {
			Expect(len(generation)).To(Equal(4))
			Expect(cap(generation)).To(Equal(4))
		})

		It("returns a generation with the center cell turned on", func() {
			Expect(generation).To(Equal(Generation{false, false, true, false}))
		})

		Context("when the length is an odd number", func() {
			BeforeEach(func() {
				generation = NewFromCenter(5)
			})

			It("rounds down when determining the center cell", func() {
				Expect(generation[2]).To(Equal(true))
			})
		})
	})

	Describe("NewFromRandom", func() {
		BeforeEach(func() {
			rand.Seed(1)
			generation = NewFromRandom(3)
		})

		It("fills in the cells randomly", func() {
			Expect(generation).To(Equal(Generation{true, true, true}))
		})
	})

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
