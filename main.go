package main

import (
	"fmt"
	"go-tsp/base"
	"io/ioutil"
	"log"
	"math/rand"
	"reflect"
	"time"
	ga "go-tsp/geneticAlgorithm"
)

var enablelogging = true

func main() {
	fmt.Println("Traveling sales person")
	// Disable logger
	if !enablelogging {
		log.SetOutput(ioutil.Discard)
	}

	// Seed default random
	//seed := time.Now().Unix()

	// Seed 1504372704 for 18 0 18
	seed := int64(1504372704)
	fmt.Println("seed: ",seed)
	rand.Seed(seed)

	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	// Fixed cities
	cities := initializeSampleCities()
	// Random Citites
	// TODO : random cities generator

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	//tspRandom()
	log.Println("Initialization completed")
	log.Println("Begin genetic algorithm")
	tspGA(&tm, 100)
}


// tspGA : Travelling sales person with genetic algorithm
// input :- TourManager, Number of generations
func tspGA(tm *base.TourManager, gen int) {
	p := base.Population{}
	p.InitPopulation(50, *tm)

	// Get initial fittest tour and it's tour distance
	fmt.Println("Start....")
	iFit:= p.GetFittest()
	iTourDistance := iFit.TourDistance()
	//fmt.Println("Initial tour distance: ", iTourDistance)

	// Evolve population "gen" number of times
	for i := 0; i < gen; i++ {
		log.Println("Generation ",i+1)
		p = ga.EvolvePopulation(p)
	}
	// Get final fittest tour and tour distance
	fmt.Println("Evolution completed")
	fFit := p.GetFittest()
	fTourDistance := fFit.TourDistance()

	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)
}

func tspRandom() {
	fmt.Println("Traveling sales person - Standard Random")
	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	// Init population
	p := base.Population{}
	p.InitPopulation(50, tm)
	fmt.Println("Find........")
	fmt.Println("Initial best distance: ", p.GetFittest().TourDistance())

}

func initializeSampleCities() []base.City {
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
	return cities
}

// Functions below are for experimentation
func gatester() {
	fmt.Println("")
	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}
	t1 := base.Tour{}
	t1.InitTourCities(tm)
	t2 := base.Tour{}
	t2.InitTourCities(tm)
	t3 := ga.Crossover(t1,t2)
	fmt.Println("t1---")
	fmt.Println(t1)
	fmt.Println("t2---")
	fmt.Println(t2)
	fmt.Println("t3---")
	fmt.Println(t3)
	fmt.Println("Mutation to t3")
	ga.Mutation(&t3)
	fmt.Println(t3)
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

func xfunc() {
	//rand.Seed(time.Now().Unix())

	x := 5 - 10
	fmt.Println(x)
	if x < 0 {
		x = -x
	}
	fmt.Println(x)

	c1 := base.GenerateRandomCity()
	fmt.Println(c1)

	c2 := base.GenerateRandomCity()
	fmt.Println(c2)

	fmt.Println(c1.DistanceTo(c2))

	random1()
	random2()
}

func random1() {
	fmt.Println("random1")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src))
	perm := r.Perm(len(src))
	fmt.Println(perm)
	for i, v := range perm {
		dest[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest)

	dest2 := make([]int, len(src))
	perm2 := r.Perm(len(src))
	fmt.Println(perm2)
	for i, v := range perm2 {
		dest2[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest2)
}

func random2() {
	fmt.Println("random2")
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	fmt.Println(perm)
	for i, v := range perm {
		dest[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest)

	dest2 := make([]int, len(src))
	perm2 := rand.Perm(len(src))
	fmt.Println(perm2)
	for i, v := range perm2 {
		dest2[v] = src[i]
	}
	fmt.Println(src)
	fmt.Println(dest2)
}
