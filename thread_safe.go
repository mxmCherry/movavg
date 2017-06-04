package movavg

import "sync"

// ThreadSafe synchronizes access to Moving Average calculator.
func ThreadSafe(ma MA) MA {
	return &threadSafe{
		ma: ma,
	}
}

// ----------------------------------------------------------------------------

type threadSafe struct {
	mx sync.RWMutex
	ma MA
}

func (a *threadSafe) Add(v float64) float64 {
	a.mx.Lock()
	avg := a.ma.Add(v)
	a.mx.Unlock()
	return avg
}

func (a *threadSafe) Avg() float64 {
	a.mx.RLock()
	avg := a.ma.Avg()
	a.mx.RUnlock()
	return avg
}
