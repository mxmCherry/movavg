package movavg

// MultiMA defines a group of Moving Average calculators.
type MultiMA interface {
	// Add recalculates Moving Average values and returns them.
	Add(value float64) (newAvgs []float64)
	// Avg returns current Moving Average values.
	Avg() (curAvgs []float64)
}
