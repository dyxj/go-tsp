package tsp

import (
	"testing"

	"github.com/the-clothing-loop/go-tsp/base"
)

func TestTsp(t *testing.T) {
	tm := &base.TourManager{}
	tm.AddCity(base.GenerateRandomCity())
	tm.AddCity(base.GenerateRandomCity())
	tm.AddCity(base.GenerateRandomCity())
	tm.AddCity(base.GenerateRandomCity())
	tm.AddCity(base.GenerateRandomCity())
	tm.AddCity(base.GenerateRandomCity())

	TspGA(tm, 20)
	t.Fail()
}
