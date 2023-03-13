package runFunctions

import (
	"github.com/the-clothing-loop/go-tsp/base"
	ga "github.com/the-clothing-loop/go-tsp/geneticAlgorithm"
)




func TspGA(tm *base.TourManager, gen int) *base.Tour {
	p := base.Population{}
	p.InitPopulation(100, *tm)

	
	iFit := p.GetFittest()

	fittestTours := make([]base.Tour, 0, gen+1)
	fittestTours = append(fittestTours, *iFit)

	for i := 1; i < gen+1; i++ {
		
		p = ga.EvolvePopulation(p)
	
		fittestTours = append(fittestTours, *p.GetFittest())
	}

	fFit := p.GetFittest()

	return fFit
}