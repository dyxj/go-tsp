package tsp

import (
	"fmt"
	"log"

	"github.com/the-clothing-loop/go-tsp/base"
	ga "github.com/the-clothing-loop/go-tsp/geneticAlgorithm"
)

// tspGA : Travelling sales person with genetic algorithm
// input :- TourManager, Number of generations


func TspGA(tm *base.TourManager, gen int)(*base.Tour) {
	p := base.Population{}
	p.InitPopulation(200, *tm)

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


	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)

	log.Println("Evolution completed")
	log.Println("Initial tour distance: ", iTourDistance)
	log.Println("Final tour distance: ", fTourDistance)

	return fFit;

}


func nTspGA(tm *base.TourManager, gen int)(*base.Tour){
	p := base.Population{}
	p.InitPopulation(200, *tm)

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

	// Store current best distance
	lastBestTourDistance := iTourDistance
	// Plot Generation 0
	for gn, t := range fittestTours {
		if t.TourDistance() < lastBestTourDistance {
			lastBestTourDistance = t.TourDistance()
			fmt.Printf("Generation %v: %v\n", gn, lastBestTourDistance)
			// Plot graph of points
		}
	}

	fmt.Println("Initial tour distance: ", iTourDistance)
	fmt.Println("Final tour distance: ", fTourDistance)

	log.Println("Evolution completed")
	log.Println("Initial tour distance: ", iTourDistance)
	log.Println("Final tour distance: ", fTourDistance)

	return fFit;
}
