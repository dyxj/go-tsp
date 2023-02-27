package base_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-clothing-loop/go-tsp/base"
)

func TestTourDistance(t *testing.T) {
	fmt.Println("Traveling sales person - Standard Random")
	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := *initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	// Init population
	p := base.Population{}
	p.InitPopulation(50, tm)
	fmt.Println("Find........")
	allDistances := make([]float64, 0, 20)
	shortest := p.GetTour(0).TourDistance()
	for i := 0; i < 20; i++ {
		d := p.GetTour(i).TourDistance()
		allDistances = append(allDistances, d)
		if shortest > d {
			shortest = d
		}
	}

	fmt.Println("Initial best distance: ", p.GetFittest().TourDistance())
	assert.Equalf(t, shortest, p.GetFittest().TourDistance(), "find the shortest distance from %+v", allDistances)
}

func cityListContain() {
	c1 := base.City{}
	c1.SetLocation(10, 20)
	c2 := base.City{}
	c2.SetLocation(30, 40)
	c3 := base.City{}
	c3.SetLocation(50, 60)

	c4 := base.City{}
	c4.SetLocation(30, 40)

	cslice := make([]base.City, 0, 20)

	cslice = append(cslice, c1)
	cslice = append(cslice, c2)
	cslice = append(cslice, c3)

	fmt.Println(cslice)
	fmt.Println(cslice[0])

	for _, c := range cslice {
		if c == c4 {
			fmt.Println("found same", c)
		}
		if reflect.DeepEqual(c, c4) {
			fmt.Println("deep equal true", c)
		}
	}

}

func initializeSampleCities() *[]base.City {
	cities := make([]base.City, 0, 20)
	// Sample
	cities = append(cities, base.GenerateCity(60, 200)) // c1
	cities = append(cities, base.GenerateCity(180, 200))
	cities = append(cities, base.GenerateCity(80, 180))
	cities = append(cities, base.GenerateCity(140, 180))
	cities = append(cities, base.GenerateCity(20, 160)) // c5
	cities = append(cities, base.GenerateCity(100, 160))
	cities = append(cities, base.GenerateCity(200, 160))
	cities = append(cities, base.GenerateCity(140, 140))
	cities = append(cities, base.GenerateCity(40, 120))
	cities = append(cities, base.GenerateCity(100, 120)) // c10
	cities = append(cities, base.GenerateCity(180, 100))
	cities = append(cities, base.GenerateCity(60, 80))
	cities = append(cities, base.GenerateCity(120, 80))
	cities = append(cities, base.GenerateCity(180, 60))
	cities = append(cities, base.GenerateCity(20, 40)) // c15
	cities = append(cities, base.GenerateCity(100, 40))
	cities = append(cities, base.GenerateCity(200, 40))
	cities = append(cities, base.GenerateCity(20, 20))
	cities = append(cities, base.GenerateCity(60, 20))
	cities = append(cities, base.GenerateCity(160, 20)) // c20

	// Sample using random seed
	// Completed testing
	return &cities
}
