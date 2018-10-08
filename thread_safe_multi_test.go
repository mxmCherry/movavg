package movavg_test

import (
	"sync"

	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ThreadSafeMulti", func() {
	It("should use underlying MultiMA object", func() {
		multi := MAMulti{
			NewSMA(2),
			NewSMA(4),
			NewSMA(8),
		}
		subject := ThreadSafeMulti(multi)

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

		Expect(subject.Add(2)).To(Equal([]float64{1.5, 1.25, 1.5}))
		Expect(subject.Avg()).To(Equal([]float64{1.5, 1.25, 1.5}))
	})
})
