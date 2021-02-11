package space

//Planet name of the planet
type Planet string

//EarthYearInSeconds number of seconds in a earth-year
const EarthYearInSeconds float64 = 31557600.0

//Age converts the age on a planet given in seconds to earth-years
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return seconds / (0.2408467 * EarthYearInSeconds)
	case "Venus":
		return seconds / (0.61519726 * EarthYearInSeconds)
	default:
		return seconds / EarthYearInSeconds
	case "Mars":
		return seconds / (1.8808158 * EarthYearInSeconds)
	case "Jupiter":
		return seconds / (11.862615 * EarthYearInSeconds)
	case "Saturn":
		return seconds / (29.447498 * EarthYearInSeconds)
	case "Uranus":
		return seconds / (84.016846 * EarthYearInSeconds)
	case "Neptune":
		return seconds / (164.79132 * EarthYearInSeconds)
	}
}
