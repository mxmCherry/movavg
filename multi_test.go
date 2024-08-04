package movavg_test

import (
	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Multi", func() {
	var _ MultiMA = Multi(nil)

	It("groups several moving average calculators", func() {
		subject := Multi{
			NewSMA(2),
			NewSMA(3),
			NewSMA(4),
		}

		// zero until values added:
		Expect(subject.Avg()).To(Equal([]float64{0, 0, 0}))

		// load initial value
		Expect(subject.Add(2)).To(Equal([]float64{2, 2, 2}))
		Expect(subject.Avg()).To(Equal([]float64{2, 2, 2}))

		Expect(subject.Add(4)).To(Equal([]float64{3, 3, 3}))
		Expect(subject.Avg()).To(Equal([]float64{3, 3, 3}))

		Expect(subject.Add(6)).To(Equal([]float64{5, 4, 4}))
		Expect(subject.Avg()).To(Equal([]float64{5, 4, 4}))

		Expect(subject.Add(8)).To(Equal([]float64{7, 6, 5}))
		Expect(subject.Avg()).To(Equal([]float64{7, 6, 5}))

		Expect(subject.Add(10)).To(Equal([]float64{9, 8, 7}))
		Expect(subject.Avg()).To(Equal([]float64{9, 8, 7}))
	})
})
