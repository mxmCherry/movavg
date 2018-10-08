package movavg

import "sync"

// MultiThreadSafe synchronizes access to Moving Average calculators group.
func MultiThreadSafe(mas MultiMA) MultiMA {
	return &multiThreadSafe{
		mas: mas,
	}
}

// ----------------------------------------------------------------------------

type multiThreadSafe struct {
	mx  sync.RWMutex
	mas MultiMA
}

func (a *multiThreadSafe) Add(v float64) []float64 {
	a.mx.Lock()
	avgs := a.mas.Add(v)
	a.mx.Unlock()
	return avgs
}

func (a *multiThreadSafe) Avg() []float64 {
	a.mx.RLock()
	avgs := a.mas.Avg()
	a.mx.RUnlock()
	return avgs
}
