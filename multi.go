package movavg

// Multi is a group of Moving Average calculators.
type Multi []MA

// Add recalculates Moving Average values and returns them.
func (m Multi) Add(v float64) (newAvgs []float64) {
	newAvgs = make([]float64, len(m))
	for i, ma := range m {
		newAvgs[i] = ma.Add(v)
	}
	return newAvgs
}

// Avg returns current Moving Average values.
func (m Multi) Avg() (curAvg []float64) {
	curAvg = make([]float64, len(m))
	for i, ma := range m {
		curAvg[i] = ma.Avg()
	}
	return curAvg
}
