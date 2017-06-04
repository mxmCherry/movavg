package movavg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMovavg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movavg Suite")
}

// ----------------------------------------------------------------------------

type mockMA struct {
	addArg float64
	addRes float64
	avgRes float64
}

func (a *mockMA) Add(v float64) float64 {
	a.addArg = v
	return a.addRes
}

func (a *mockMA) Avg() float64 {
	return a.avgRes
}
