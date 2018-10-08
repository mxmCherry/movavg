package movavg

import "sync"

// MultiThreadSafe synchronizes access to Moving Average calculators group.
func MultiThreadSafe(m MultiMA) MultiMA {
	return &multiThreadSafe{
		multiMA: m,
	}
}

// ----------------------------------------------------------------------------

type multiThreadSafe struct {
	mx      sync.RWMutex
	multiMA MultiMA
}

func (a *multiThreadSafe) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.multiMA.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *multiThreadSafe) Avg() []float64 {
	a.mx.RLock()
	avgs := a.multiMA.Avg()
	a.mx.RUnlock()
	return avgs
}
