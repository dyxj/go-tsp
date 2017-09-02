package main

import (
	"fmt"
	b "go-tsp/base"
	"io/ioutil"
	"log"
	"math/rand"
	"reflect"
	"time"
	ga "go-tsp/geneticAlgorithm"
)

func main() {
	fmt.Println("Traveling sales person")
	// Disable logger
	log.SetOutput(ioutil.Discard)
	// Seed default random
	//seed := time.Now().Unix()
	// 1504372704 for 18 0 18
	seed := int64(1504372704)
	fmt.Println("seed: ",seed)
	rand.Seed(seed)
	//tspRandom()

	// Test Area
	crossovertester()
}

func tspRandom() {
	fmt.Println("Traveling sales person - Standard Random")
	// Init TourManager
	tm := b.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	// Init population
	p := b.Population{}
	p.InitPopulation(50, tm)
	fmt.Println("Find........")
	fmt.Println("Initial best distance: ", p.GetFittest().TourDistance())

}

func initializeSampleCities() []b.City {
	cities := make([]b.City, 0, 20)
	// Sample
	cities = append(cities, b.GenerateCity(60, 200)) // c1
	cities = append(cities, b.GenerateCity(180, 200))
	cities = append(cities, b.GenerateCity(80, 180))
	cities = append(cities, b.GenerateCity(140, 180))
	cities = append(cities, b.GenerateCity(20, 160)) // c5
	cities = append(cities, b.GenerateCity(100, 160))
	cities = append(cities, b.GenerateCity(200, 160))
	cities = append(cities, b.GenerateCity(140, 140))
	cities = append(cities, b.GenerateCity(40, 120))
	cities = append(cities, b.GenerateCity(100, 120)) // c10
	cities = append(cities, b.GenerateCity(180, 100))
	cities = append(cities, b.GenerateCity(60, 80))
	cities = append(cities, b.GenerateCity(120, 80))
	cities = append(cities, b.GenerateCity(180, 60))
	cities = append(cities, b.GenerateCity(20, 40)) // c15
	cities = append(cities, b.GenerateCity(100, 40))
	cities = append(cities, b.GenerateCity(200, 40))
	cities = append(cities, b.GenerateCity(20, 20))
	cities = append(cities, b.GenerateCity(60, 20))
	cities = append(cities, b.GenerateCity(160, 20)) // c20

	// Sample using random seed
	// Completed testing
	return cities
}

// Functions below are for experimentation
func crossovertester() {
	fmt.Println("")
	// Init TourManager
	tm := b.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}
	t1 := b.Tour{}
	t1.InitTourCities(tm)
	t2 := b.Tour{}
	t2.InitTourCities(tm)
	t3 := ga.Crossover(t1,t2)
	fmt.Println("t1---")
	fmt.Println(t1)
	fmt.Println("t2---")
	fmt.Println(t2)
	fmt.Println("t3---")
	fmt.Println(t3)
}

func cityListContain() {
	c1 := b.City{}
	c1.SetLocation(10, 20)
	c2 := b.City{}
	c2.SetLocation(30, 40)
	c3 := b.City{}
	c3.SetLocation(50, 60)

	c4 := b.City{}
	c4.SetLocation(30, 40)

	cslice := make([]b.City, 0, 20)

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

	c1 := b.GenerateRandomCity()
	fmt.Println(c1)

	c2 := b.GenerateRandomCity()
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
