package movavg

import "sync"

// ThreadSafeMulti synchronizes access to Moving Average calculator.
func ThreadSafeMulti(multiMA MultiMA) MultiMA {
	return &ThreadSafeMultiMA{
		multiMA: multiMA,
	}
}

// ----------------------------------------------------------------------------

type ThreadSafeMultiMA struct {
	mx      sync.RWMutex
	multiMA MultiMA
}

func (a *ThreadSafeMultiMA) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.multiMA.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *ThreadSafeMultiMA) Avg() []float64 {
	a.mx.RLock()
	avgs := a.multiMA.Avg()
	a.mx.RUnlock()
	return avgs
}
