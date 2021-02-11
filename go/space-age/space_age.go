package space

//Planet name of the planet
type Planet string

//EarthYearInSeconds number of seconds in a earth-year
const EarthYearInSeconds float64 = 31557600.0

//Age converts the age on a planet given in seconds to earth-years
func Age(seconds float64, planet Planet) float64 {
	PlanetEarthYears := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (PlanetEarthYears[planet] * EarthYearInSeconds)
}
