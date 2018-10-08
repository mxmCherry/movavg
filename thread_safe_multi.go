package movavg

import "sync"

// ThreadSafeMulti synchronizes access to Moving Average calculator.
func ThreadSafeMulti(multiMA MultiMA) MultiMA {
	return &threadSafeMulti{
		multiMA: multiMA,
	}
}

// ----------------------------------------------------------------------------

type threadSafeMulti struct {
	mx      sync.RWMutex
	multiMA MultiMA
}

func (a *threadSafeMulti) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.multiMA.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *threadSafeMulti) Avg() []float64 {
	a.mx.RLock()
	avgs := a.multiMA.Avg()
	a.mx.RUnlock()
	return avgs
}
