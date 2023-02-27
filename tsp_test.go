package tsp

import (
	"testing"

	"github.com/dyxj/go-tsp/base"
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
