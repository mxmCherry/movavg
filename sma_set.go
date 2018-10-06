package movavg

// SMASet is a set of related calculators of Simple Moving Average:
// https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average
//
// It is as precise, as float64 is: https://golang.org/pkg/builtin/#float64 (good performance, not-so-good precision).
//
// It is not thread-safe, if you need to access it from multiple goroutines - secure it with ThreadSafeSet(...).
//
// It is safe to use even without adding values (then Avg() will return zeroes),
// or without filling entire window (then average of currently added values will be returned).
//
// To initialize calculator with current average value, simply Add(...) this value.
type SMASet struct {
	smas []*SMA
}

// NewSMASet constructs new Simple Moving Average calculator set.
// Window arg must have 1 or more elements. For example, windows
// of 1 minute, 5 minutes, 30 minutes, 1 hour for a given measurement.
func NewSMASet(windows []int) *SMASet {
	if windows == nil || len(windows) == 0 {
		panic("no windows defined")
	}

	s := &SMASet{}
	s.smas = make([]*SMA, len(windows))
	for i, w := range windows {
		s.smas[i] = NewSMA(w)
	}
	return s
}

// Add recalculates Simple Moving Average values and returns them. 
func (s *SMASet) Add(value float64) (newAvgs []float64) {
	result := make([]float64, len(s.smas))
	for i, sma := range s.smas {
		result[i] = sma.Add(value)
	}
	return result
}

// Avg returns current Simple Moving Average values.
func (s *SMASet) Avg() (curAvg []float64) {
	result := make([]float64, len(s.smas))
	for i, sma := range s.smas {
		result[i] = sma.Avg()
	}
	return result
}