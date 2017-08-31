package base

import (
	"math/rand"
	"math"
	"fmt"
)

// City
type City struct {
	X int
	Y int
}

func (a *City) NewRandomCity() {
	a.X = rand.Int() * 100
	a.Y = rand.Int() * 100
}

func (a *City) NewCity(x int, y int) {
	a.X = x
	a.Y = y
}

func (a *City) DistanceTo(c City) float64{
	idx := a.X - c.X
	idy := a.Y - c.Y

	if idx < 0 {
		idx = -idx
	}
	if idy < 0 {
		idy = -idy
	}

	fdx := float64(idx)
	fdy := float64(idy)

	fd := math.Sqrt((fdx*fdx) + (fdy*fdy))
	return fd
}

func (a *City) String() string {
	return fmt.Sprint(a.X, " ", a.Y)
}