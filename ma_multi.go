package movavg

// MAMulti is a group of Moving Average calculators:
// https://en.wikipedia.org/wiki/Moving_average
type MAMulti []MA

// Add recalculates Simple Moving Average values and returns them.
func (mas MAMulti) Add(value float64) (newAvgs []float64) {
	result := make([]float64, len(mas))
	for i, ma := range mas {
		result[i] = ma.Add(value)
	}
	return result
}

// Avg returns current Simple Moving Average values.
func (mas MAMulti) Avg() (curAvg []float64) {
	result := make([]float64, len(mas))
	for i, ma := range mas {
		result[i] = ma.Avg()
	}
	return result
}
