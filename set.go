package movavg

// Set is a group of Moving Average calculators:
// https://en.wikipedia.org/wiki/Moving_average
type Set []MA

// Add recalculates Simple Moving Average values and returns them.
func (mas Set) Add(v float64) (newAvgs []float64) {
	newAvgs = make([]float64, len(mas))
	for i, ma := range mas {
		newAvgs[i] = ma.Add(v)
	}
	return newAvgs
}

// Avg returns current Simple Moving Average values.
func (mas Set) Avg() (curAvg []float64) {
	curAvg = make([]float64, len(mas))
	for i, ma := range mas {
		curAvg[i] = ma.Avg()
	}
	return curAvg
}
