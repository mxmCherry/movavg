package movavg_test

import (
	"sync"

	. "github.com/blocktop/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ThreadSafeSet", func() {
	It("should use underlying MASet object", func() {
		set := NewSMASet([]int{2, 4, 8})
		subject := ThreadSafeSet(set)

		valsets := [][]float64{
			[]float64{1, 1, 1, 1},
			[]float64{2, 2, 2, 2}}

		var wg sync.WaitGroup

		for _, vset := range valsets {
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
