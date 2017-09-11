/*
	THIS FILE IS ONLY FOR DEVELOPMENT AND TESTING FUNCTIONS
*/
package main

import (
	"fmt"
	"go-tsp/base"
	ga "go-tsp/geneticAlgorithm"
	"math/rand"
	"reflect"
	"time"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
	"path/filepath"
	"gonum.org/v1/plot/plotutil"
)

func devTest() {
	fmt.Println("Run dev test")
	//plotTest()
}

// Functions below are for experimentation
func plotTest() {
	//cities := *initializeSampleCities()
	cities := *initRandomCities(20)
	p, err := plot.New()
	if err != nil  {
		panic(err)
	}
	p.Title.Text = "Test Plot ABC"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	pts := CitiesToPoints2(&cities)
	// Plot, Add lines and points
	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		panic(err)
	}
	// Create labels plotter
	labels, err := plotter.NewLabels(pts)
	if err != nil {
		panic(err)
	}
	lp, sp, err := plotter.NewLinePoints(pts)
	if err != nil {
		panic(err)
	}
	fmt.Println(lp)
	fmt.Println(sp)
	//p.Add(lp, sp, labels)
	p.Add(labels)

	// Directory names
	rootpath := "devTest"
	dname := "123/"
	// File path to store images
	dname = filepath.Join(rootpath, dname)
	// Create Directory
	if err := os.MkdirAll(dname, 0666); err != nil {
		panic(err)
	}
	// File name
	imgname := "Map.png"
	// File path
	fpath := filepath.Join(dname, imgname)
	// Save plot to png
	if err := p.Save(30*vg.Centimeter, 30*vg.Centimeter, fpath); err != nil {
		panic(err)
	}
}

func CitiesToPoints(cities *[]base.City) plotter.XYs {
	c := *cities
	l := len(c)
	fmt.Println(l)
	pts := make(plotter.XYs, l+1)
	for i, v := range c {
		pts[i].X = float64(v.X())
		pts[i].Y = float64(v.Y())
	}
	pts[l].X = float64(c[0].X())
	pts[l].Y = float64(c[0].Y())
	return pts
}

func CitiesToPoints2(cities *[]base.City) plotter.XYLabels {
	c := *cities
	l := len(c)
	fmt.Println(l)
	pts := make(plotter.XYs, l+1)
	labels := make([]string, l+1)
	for i, v := range c {
		pts[i].X = float64(v.X())
		pts[i].Y = float64(v.Y())
		labels[i] = fmt.Sprintf("%d, %d, %d", i, v.X(), v.Y())
	}
	pts[l].X = float64(c[0].X())
	pts[l].Y = float64(c[0].Y())
	labels[l] = fmt.Sprintf("%d, %d, %d", 0, c[0].X(), c[0].Y())

	xylabels := plotter.XYLabels{pts,labels}
	return xylabels
}

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