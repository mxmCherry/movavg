package movavg

// Set is a group of Moving Average calculators:
// https://en.wikipedia.org/wiki/Moving_average
type Set []MA

// Add recalculates Simple Moving Average values and returns them.
func (mas Set) Add(value float64) (newAvgs []float64) {
	result := make([]float64, len(mas))
	for i, ma := range mas {
		result[i] = ma.Add(value)
	}
	return result
}

// Avg returns current Simple Moving Average values.
func (mas Set) Avg() (curAvg []float64) {
	result := make([]float64, len(mas))
	for i, ma := range mas {
		result[i] = ma.Avg()
	}
	return result
}
