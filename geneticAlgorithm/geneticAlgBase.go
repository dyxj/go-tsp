package geneticAlgorithm

import (
	"log"
	"go-tsp/base"
	"math/rand"
)

// Genetic Algorithm Parameters
var (
	mutationRate        float64 = 0.015
	tournamentSize      int     = 10
	elitism             bool    = true
	randomCrossoverRate         = false
	defCrossoverRate    float32 = 0.7
)

func CrossoverRate() float32 {
	if randomCrossoverRate {
		return rand.Float32()
	}
	return defCrossoverRate
}

// Crossover : performs multi point cross over with 2 parents
// Assumption - parents have equal size
func Crossover(p1 base.Tour, p2 base.Tour) base.Tour {
	// Size
	size := p1.TourSize()
	// Child Tour
	c := base.Tour{}
	c.InitTour(size)

	// Number of crossover
	nc := int(CrossoverRate() * float32(size))
	if nc == 0 {
		log.Println("no crossover")
		return p1
	}
	// Start positions of cross over for parent 1
	sp := int(rand.Float32() * float32(size))
	// End position of cross over for parent 1
	ep := (sp + nc) % size
	// Parent 2 slots
	p2s := make([]int, 0, size-nc)
	log.Println(size, sp, nc, ep) // For debugging
	// Populate child with parent 1
	if sp < ep {
		for i := 0; i < size; i++ {
			if i >= sp && i < ep {
				c.SetCity(i, p1.GetCity(i))
			} else {
				p2s = append(p2s, i)
			}
		}
	} else if sp > ep {
		for i := 0; i < size; i++ {
			if !(i >= ep && i < sp) {
				c.SetCity(i, p1.GetCity(i))
			} else {
				p2s = append(p2s, i)
			}
		}
	}

	// For debugging
	msCity := ""
	j := 0
	// Populate child with parent 2 cities that are missing
	for i := 0; i < size; i++ {
		// Check if child contains city
		if !c.ContainCity(p2.GetCity(i)) {
			c.SetCity(p2s[j], p2.GetCity(i))
			j++
			// For debugging
			msCity += p2.GetCity(i).String() + " "
		}
	}
	log.Println(msCity)
	log.Println(p2s)
	log.Println(len(p2s))
	return c
}

// Mutation : Performs swap mutation
// Chance of mutation for each City based on mutation rate
func Mutation(in *base.Tour) {
	// for each city
	for p1 := 0; p1 < in.TourSize(); p1++ {
		if rand.Float64() < mutationRate {
			// Select 2nd city to perform swap
			p2 := int(float64(in.TourSize()) * rand.Float64())
			log.Println("Mutation occured", p1, "swap", p2)
			// Temp store city
			c1 := in.GetCity(p1)
			c2 := in.GetCity(p2)
			// Swap Cities
			in.SetCity(p1, c2)
			in.SetCity(p2, c1)
		}
	}
}

// TournamentSelection : select a group at random and pick the best parent
func TournamentSelection(pop base.Population) base.Tour {
	tourny := base.Population{}
	tourny.InitEmpty(tournamentSize)

	for i := 0; i < tournamentSize; i++ {
		r := int(rand.Float64() * float64(pop.PopulationSize()))
		tourny.SaveTour(i, *pop.GetTour(r))
	}
	// fittest tour
	fTour := tourny.GetFittest()
	return *fTour
}

// EvolvePopulation : evolves population by :-
/*
	- Selecting 2 parents using tournament selection
	- Perform crossover to obtain child
	- Mutate child based on probability
	- return new population
*/
func EvolvePopulation(pop base.Population) base.Population{
	npop := base.Population{}
	npop.InitEmpty(pop.PopulationSize())

	popOffset := 0
	if elitism {
		npop.SaveTour(0, *pop.GetFittest())
		popOffset = 1
	}

	for i := popOffset; i < npop.PopulationSize(); i++ {
		p1 := TournamentSelection(pop)
		p2 := TournamentSelection(pop)		
		child := Crossover(p1,p2)
		Mutation(&child)
		npop.SaveTour(i,child)
	}
	return npop
}
