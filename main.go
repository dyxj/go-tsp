package main

import (
	"fmt"
	"go-tsp/base"
	ga "go-tsp/geneticAlgorithm"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	// For development purpose
	enablelogging = false

	// To store results
	rootpath = "tsp"
	// Define seed for default rand
	seed = int64(1504372704)
	//seed = time.Now().Unix()

	// Number of generation to loop through
	noGen = 100
	// Population Size
	popSize = 200
)

func main() {
	fmt.Println("Traveling sales person")
	// Enable/disable logger
	if enablelogging {
		//f, err := os.OpenFile("tsplog", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		f, err := os.OpenFile("tsplog.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v\n", err)
		}
		defer f.Close()
		log.SetOutput(f)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	// Set seed value for default source
	rand.Seed(seed)
	fmt.Println("seed: ", seed)

	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	var cities []base.City
	cities = *initRandomCities(20)

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	//tspRandom()
	log.Println("Initialization completed")
	log.Println("Begin genetic algorithm")
	tspGA(&tm, noGen)
}

// tspGA : Travelling sales person with genetic algorithm
// input :- TourManager, Number of generations
func tspGA(tm *base.TourManager, gen int) {
	p := base.Population{}
	p.InitPopulation(popSize, *tm)

	// Get initial fittest tour and it's tour distance
	fmt.Println("Start....")
	iFit := p.GetFittest()
	iTourDistance := iFit.TourDistance()
	//fmt.Println("Initial tour distance: ", iTourDistance)

	// Map to store fittest tours
	fittestTours := make([]base.Tour, 0, gen+1)
	fittestTours = append(fittestTours, *iFit)
	// Evolve population "gen" number of times
	for i := 1; i < gen+1; i++ {
		log.Println("Generation ", i)
		p = ga.EvolvePopulation(p)
		// Store fittest for each generation
		fittestTours = append(fittestTours, *p.GetFittest())
	}
	// Get final fittest tour and tour distance
	fmt.Println("Evolution completed")
	fFit := p.GetFittest()
	fTourDistance := fFit.TourDistance()

	fmt.Println("Print and save image of fittest by generation-----------")
	// Remove old data
	dname := fmt.Sprintf("%d", seed)
	dname = filepath.Join(rootpath, dname)
	os.RemoveAll(dname)

	// Store current best distance
	lastBestTourDistance := iTourDistance
	// Plot Generation 0
	visualization(iFit, 0, seed)
	for gn, t := range fittestTours {
		if t.TourDistance() < lastBestTourDistance {
			lastBestTourDistance = t.TourDistance()
			fmt.Printf("Generation %v: %v\n", gn, lastBestTourDistance)
			// Plot graph of points
			visualization(&t, gn, seed)
		}
	}

	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)

	log.Println("Evolution completed")
	log.Println("Initial tour distance: ", iTourDistance)
	log.Println("Final tour distance: ", fTourDistance)
}

func initRandomCities(cityCount int) *[]base.City {
	cities := make([]base.City, 0, cityCount)

	for i := 0; i < cityCount; i++ {
		cities = append(cities, base.GenerateRandomCity())
	}
	return &cities
}

// Save tour as graph
func visualization(t *base.Tour, gen int, rseed int64) {
	// Init plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	// Set plot styles
	p.Title.Text = fmt.Sprintf("Seed: %d, Generation %d, Tour Distance: %f", rseed, gen, t.TourDistance())
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	// Construct points with labels
	pts_labels := TourToPoints(t)
	// Plot points
	err = plotutil.AddLinePoints(p, pts_labels)
	if err != nil {
		panic(err)
	}
	// Create Labels plotter
	plabels, err := plotter.NewLabels(pts_labels)
	if err != nil {
		panic(err)
	}
	p.Add(plabels)

	// Create Directory (based on seed value)
	dname := fmt.Sprintf("%d", rseed)
	dname = filepath.Join(rootpath, dname)
	if err := os.MkdirAll(dname, 0755); err != nil {
		panic(err)
	}
	// Define file path
	fpath := filepath.Join(dname, fmt.Sprintf("%d.png", gen))
	// Save plot to png
	if err := p.Save(30*vg.Centimeter, 30*vg.Centimeter, fpath); err != nil {
		panic(err)
	}
}

func TourToPoints(t *base.Tour) plotter.XYLabels {
	tLen := t.TourSize()
	pts := make(plotter.XYs, tLen+1)
	labels := make([]string, tLen+1)

	c0 := t.GetCity(0)
	pts[0].X = float64(c0.X())
	pts[0].Y = float64(c0.Y())
	pts[tLen].X = float64(c0.X())
	pts[tLen].Y = float64(c0.Y())
	labels[0] = fmt.Sprintf("%d, %d, %d", 0, c0.X(), c0.Y())
	for i := 1; i < tLen; i++ {
		c := t.GetCity(i)
		pts[i].X = float64(c.X())
		pts[i].Y = float64(c.Y())
		labels[i] = fmt.Sprintf("%d, %d, %d", i, c.X(), c.Y())
	}
	xylabels := plotter.XYLabels{pts, labels}
	return xylabels
}
