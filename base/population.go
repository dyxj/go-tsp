package base

import "log"

type Population struct {
	tours []Tour
}

func (a *Population) InitEmpty(pSize int) {
	a.tours = make([]Tour, pSize)
}

func (a *Population) InitPopulation(pSize int, tm TourManager) {
	a.tours = make([]Tour, pSize)
	for i := 0; i < pSize; i++ {
		nT := Tour{}
		nT.InitTourCities(tm)
		a.SaveTour(i, nT)
	}
}

func (a *Population) SaveTour(i int, t Tour) {
	a.tours[i] = t
}

func (a *Population) GetTour(i int) *Tour {
	return &a.tours[i]
}

func (a *Population) PopulationSize() int {
	return len(a.tours)
}

func (a *Population) GetFittest() *Tour {
	fittest := a.tours[0]
	// Loop through all tours taken by population and determine the fittest
	for i := 0; i < a.PopulationSize(); i++ {
		log.Println("Current Tour: ", i)
		if fittest.Fitness() <= a.GetTour(i).Fitness() {
			fittest = *a.GetTour(i)
		}
	}
	return &fittest
}
