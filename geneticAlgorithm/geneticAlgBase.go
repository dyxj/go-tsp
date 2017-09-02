package geneticAlgorithm

import (
	"go-tsp/base"
	"math/rand"
	"fmt"
)

// Genetic Algorithm Parameters
var (
	mutationRate        float64 = 0.015
	tournamentSize      int     = 5
	elitism             bool    = true
	randomCrossoverRate         = true
	defCrossoverRate    float32 = 0.7
)

func CrossoverRate() float32 {
	if randomCrossoverRate {
		return rand.Float32()
	}
	return defCrossoverRate
}

// crossover : performs cross over with 2 parents
// Assumption - parents have equal size
func Crossover(p1 base.Tour, p2 base.Tour) base.Tour{
	// Size
	size := p1.TourSize()
	// Child Tour
	c := base.Tour{}
	c.InitTour(size)

	// Number of crossover
	nc := int(CrossoverRate() * float32(size))
	if nc == 0 {
		fmt.Println("no crossover")
		return p1
	}
	// Start positions of cross over for parent 1
	sp := int(rand.Float32() * float32(size))
	// End position of cross over for parent 1
	ep := (sp+nc) % size
	// Parent 2 slots
	p2s := make([]int, 0, size-nc)
	fmt.Println(size,sp,nc,ep)
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
	fmt.Println(msCity)
	fmt.Println(p2s)
	fmt.Println(len(p2s))
	return c
}

func Crossover_Alternative(){
	//for i := 0; i < nc; i++ {
	//	var p int = (i + sp) % size
	//	c.SetCity(p, p1.GetCity(p))
	//}
}
