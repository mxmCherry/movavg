package movavg

import "sync"

// ThreadSafe synchronizes access to Moving Average calculator.
func ThreadSafeSet(ma_set MASet) MASet {
	return &threadSafeSet{
		ma_set: ma_set,
	}
}

// ----------------------------------------------------------------------------

type threadSafeSet struct {
	mx sync.RWMutex
	ma_set MASet
}

func (a *threadSafeSet) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.ma_set.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *threadSafeSet) Avg() []float64 {
	a.mx.RLock()
	avgs := a.ma_set.Avg()
	a.mx.RUnlock()
	return avgs
}
