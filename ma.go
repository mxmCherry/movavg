package movavg

// MA defines Moving Average calculator:
// https://en.wikipedia.org/wiki/Moving_average
type MA interface {
	// Add recalculates Moving Average value and returns it.
	Add(value float64) (newAvg float64)
	// Avg returns current Moving Average value.
	Avg() (curAvg float64)
}
