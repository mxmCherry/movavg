package movavg_test

import (
	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SMA", func() {
	var _ = (*SMA)(nil)

	It("should calculate Simple Moving Average", func() {
		subject := NewSMA(3)

		// zero until values added:
		Expect(subject.Avg()).To(BeZero())

		// load initial value
		Expect(subject.Add(1)).To(Equal(float64(1)))
		Expect(subject.Avg()).To(Equal(float64(1)))

		// (1 + 2) / 2
		Expect(subject.Add(2)).To(Equal(float64(1.5)))
		Expect(subject.Avg()).To(Equal(float64(1.5)))

		// (1 + 2 + 3) / 3
		Expect(subject.Add(3)).To(Equal(float64(2)))
		Expect(subject.Avg()).To(Equal(float64(2)))

		// (2 + 3 + 4) / 3
		Expect(subject.Add(4)).To(Equal(float64(3)))
		Expect(subject.Avg()).To(Equal(float64(3)))

		// (3 + 4 + 5) / 3
		Expect(subject.Add(5)).To(Equal(float64(4)))
		Expect(subject.Avg()).To(Equal(float64(4)))
	})
})
