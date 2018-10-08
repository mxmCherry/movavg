package movavg

import "sync"

// MultiThreadSafe synchronizes access to Moving Average calculator.
func MultiThreadSafe(mas MultiMA) MultiMA {
	return &threadSafeMulti{
		mas: mas,
	}
}

// ----------------------------------------------------------------------------

type threadSafeMulti struct {
	mx  sync.RWMutex
	mas MultiMA
}

func (a *threadSafeMulti) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.mas.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *threadSafeMulti) Avg() []float64 {
	a.mx.RLock()
	avgs := a.mas.Avg()
	a.mx.RUnlock()
	return avgs
}
