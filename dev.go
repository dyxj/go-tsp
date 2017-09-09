package main

import (
	"fmt"
	"go-tsp/base"
	ga "go-tsp/geneticAlgorithm"
	"math/rand"
	"reflect"
	"time"
)

// Functions below are for experimentation
func tspRandom() {
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
	fmt.Println("Initial best distance: ", p.GetFittest().TourDistance())
}

func gatester() {
	fmt.Println("")
	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := *initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}
	t1 := base.Tour{}
	t1.InitTourCities(tm)
	t2 := base.Tour{}
	t2.InitTourCities(tm)
	t3 := ga.Crossover(t1, t2)
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
