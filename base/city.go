package base

import (
	"fmt"
	"math"
	"math/rand"
)

// City : coordinates of city
type City struct {
	x float64
	y float64
}

// GenerateRandomCity : Generate city with random coordinates
func GenerateRandomCity() City {
	c := City{}
	c.x = rand.Float64()*(100) * 100
	c.y = rand.Float64()*(100) * 100
	return c
}

// GenerateCity : Generate city with user defined coordinates
func GenerateCity(x float64, y float64) City {
	c := City{}
	c.x = x
	c.y = y
	return c
}

// SetLocation : User defined coordinates for a city
func (a *City) SetLocation(x float64, y float64) {
	a.x = x
	a.y = y
}

// DistanceTo : distance of current city to target city
func (a *City) DistanceTo(c City) float64 {
	idx := a.x - c.x
	idy := a.y - c.y

	if idx < 0 {
		idx = -idx
	}
	if idy < 0 {
		idy = -idy
	}

	fdx := float64(idx)
	fdy := float64(idy)

	fd := math.Sqrt((fdx * fdx) + (fdy * fdy))
	return fd
}

func (a *City) X() float64 {
	return a.x
}

func (a *City) Y() float64 {
	return a.y
}

func (a City) String() string {
	return fmt.Sprintf("{x%v y%v}", a.x, a.y)
}

// ShuffleCities : return a shuffled []City given input []City
func ShuffleCities(in []City) []City {
	out := make([]City, len(in), cap(in))
	perm := rand.Perm(len(in))
	for i, v := range perm {
		out[v] = in[i]
	}
	return out
}
