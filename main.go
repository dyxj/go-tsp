package main

import (
	"fmt"
	"go-tsp/base"
	ga "go-tsp/geneticAlgorithm"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

var enablelogging = true
var randomCityBool = true

func main() {
	fmt.Println("Traveling sales person")
	// Disable logger
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

	// Seed default random
	//seed := time.Now().Unix()

	// Seed 1504372704 for 18 0 18
	seed := int64(1504372704)
	fmt.Println("seed: ", seed)
	rand.Seed(seed)

	// Init TourManager
	tm := base.TourManager{}
	tm.NewTourManager()

	// Generate Cities
	var cities []base.City
	if randomCityBool {
		cities = *initRandomCities(20)
	} else {
		cities = *initializeSampleCities()
	}

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
	iFit := p.GetFittest()
	iTourDistance := iFit.TourDistance()
	//fmt.Println("Initial tour distance: ", iTourDistance)

	// Evolve population "gen" number of times
	for i := 0; i < gen; i++ {
		log.Println("Generation ", i+1)
		p = ga.EvolvePopulation(p)
	}
	// Get final fittest tour and tour distance
	fmt.Println("Evolution completed")
	fFit := p.GetFittest()
	fTourDistance := fFit.TourDistance()

	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)

	log.Println("Evolution completed")
	log.Println("Initial tour distance: ", iTourDistance)
	log.Println("Final tour distance: ", fTourDistance)
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

func initRandomCities(cityCount int) *[]base.City {
	cities := make([]base.City, 0, cityCount)

	for i := 0; i < cityCount; i++ {
		cities = append(cities, base.GenerateRandomCity())
	}

	return &cities
}
