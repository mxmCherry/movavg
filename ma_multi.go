package movavg

// MAMulti is a group of Moving Average calculators:
// https://en.wikipedia.org/wiki/Moving_average
type MAMulti struct {
	mas []MA
}

// NewMulti constructs a new Simple Moving Average calculator set from
// the given Moving Average instances.
func NewMulti(mas ...MA) MultiMA {
	m := &MAMulti{}
	m.mas = make([]MA, len(mas))
	for i, ma := range mas {
		m.mas[i] = ma
	}
	return m
}

// NewMultiSMA constructs new Simple Moving Average calculator set.
// Window arg must have 1 or more elements. For example, windows
// of 1 minute, 5 minutes, 30 minutes, 1 hour for a given measurement.
func NewMultiSMA(windows []int) MultiMA {
	if windows == nil || len(windows) == 0 {
		panic("no windows defined")
	}

	mas := make([]MA, len(windows))
	for i, w := range windows {
		mas[i] = NewSMA(w)
	}
	return NewMulti(mas...)
}

// Add recalculates Simple Moving Average values and returns them.
func (m *MAMulti) Add(value float64) (newAvgs []float64) {
	result := make([]float64, len(m.mas))
	for i, sma := range m.mas {
		result[i] = sma.Add(value)
	}
	return result
}

// Avg returns current Simple Moving Average values.
func (m *MAMulti) Avg() (curAvg []float64) {
	result := make([]float64, len(m.mas))
	for i, sma := range m.mas {
		result[i] = sma.Avg()
	}
	return result
}
