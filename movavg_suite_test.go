package movavg_test

import (
	"testing"

	. "github.com/mxmCherry/movavg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMovavg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movavg Suite")
}

// ----------------------------------------------------------------------------

var _ MA = (*mockMA)(nil)

type mockMA struct {
	addArgs []float64
	addRes  float64
	avgRes  float64
}

func (a *mockMA) Add(v float64) float64 {
	a.addArgs = append(a.addArgs, v)
	return a.addRes
}

func (a *mockMA) Avg() float64 {
	return a.avgRes
}

// ----------------------------------------------------------------------------

var _ MultiMA = (*mockSet)(nil)

type mockSet struct {
	addArgs []float64
	addRes  []float64
	avgRes  []float64
}

func (s *mockSet) Add(v float64) []float64 {
	s.addArgs = append(s.addArgs, v)
	return s.addRes
}

func (s *mockSet) Avg() []float64 {
	return s.avgRes
}
