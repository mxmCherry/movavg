package movavg_test

import (
	"sync"

	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultiThreadSafe", func() {
	It("should use underlying MultiMA object", func() {
		underlying := &mockMultiMA{
			addRes: []float64{0.1, 0.2, 0.3}, // dummy
			avgRes: []float64{0.4, 0.5, 0.6}, // dummy
		}
		subject := MultiThreadSafe(underlying)

		vsets := [][]float64{
			{1, 1, 1, 1},
			{2, 2, 2, 2},
		}

		var wg sync.WaitGroup

		for _, vset := range vsets {
			wg.Add(1)
			go func(vset []float64) {
				for _, v := range vset {
					subject.Add(v)
				}
				wg.Done()
			}(vset)
		}

		wg.Wait()

		Expect(subject.Add(3)).To(Equal([]float64{0.1, 0.2, 0.3})) // dummy
		Expect(subject.Avg()).To(Equal([]float64{0.4, 0.5, 0.6}))  // dummy

		Expect(underlying.addArgs).To(ConsistOf([]float64{
			1, 1, 1, 1, // four ones (added concurrently)
			2, 2, 2, 2, // four twos (added concurrently)
			3, // added last
		}))
	})
})
