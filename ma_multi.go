package movavg

// MA defines Moving Average calculator:
// https://en.wikipedia.org/wiki/Moving_average
type MultiMA interface {
	// Add recalculates Moving Average values and returns them.
	Add(value float64) (newAvgs []float64)
	// Avg returns current Moving Average values.
	Avg() (curAvgs []float64)
}
