package movavg

// MAMulti is a set of related calculators of Moving Average:
// https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average
//
// It is as precise, as float64 is: https://golang.org/pkg/builtin/#float64 (good performance, not-so-good precision).
//
// It is not thread-safe, if you need to access it from multiple goroutines - secure it with ThreadSafeMulti(...).
//
// It is safe to use even without adding values (then Avg() will return zeroes),
// or without filling entire window (then average of currently added values will be returned).
//
// To initialize calculator with current average value, simply Add(...) this value.
type MAMulti struct {
	mas []MA
}

// NewMulti cronstructs a new Simple Moving Average calculator set from
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
