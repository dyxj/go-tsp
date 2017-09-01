package base

// ToueManager : Contains list of of cities to be visited
type TourManager struct {
	destCities []City
}

// NewTourManager : Initialize TourManager
func (a *TourManager) NewTourManager() {
	a.destCities = make([]City, 0, 50)
}

func (a *TourManager) AddCity(c City) {
	a.destCities = append(a.destCities, c)
}

func (a *TourManager) GetCity(i int) City {
	return a.destCities[i]
}

func (a *TourManager) NumberOfCities() int {
	return len(a.destCities)
}
