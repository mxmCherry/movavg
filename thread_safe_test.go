package movavg_test

import (
	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ThreadSafe", func() {
	It("should use underlying MA object", func() {
		ma := &mockMA{
			addRes: 0.1,
			avgRes: 0.2,
		}
		subject := ThreadSafe(ma)

		Expect(subject.Add(0.3)).To(Equal(float64(0.1)))
		Expect(subject.Avg()).To(Equal(float64(0.2)))
	})
})
