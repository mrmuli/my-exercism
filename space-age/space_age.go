package space

type Planet string

// Age returns an age value based on seconds input and planet
func Age(seconds float64, planet Planet) float64 {
	// calculate earth days period from input
	var age float64
	switch planet {
	case "Earth":
		age = (seconds / 86400) / 365.25
	case "Mercury":
		age = dayDiff(seconds, 0.2408467)
	case "Venus":
		age = dayDiff(seconds, 0.61519726)
	case "Mars":
		age = dayDiff(seconds, 1.8808158)
	case "Jupiter":
		age = dayDiff(seconds, 11.862615)
	case "Saturn":
		age = dayDiff(seconds, 29.447498)
	case "Uranus":
		age = dayDiff(seconds, 84.016846)
	case "Neptune":
		age = dayDiff(seconds, 164.79132)
	}
	return age
}

func dayDiff(sec float64, diff float64) float64 {
	// not the most friendly way but...
	age := (sec / 86400) / (365.25 * diff)
	return age
}
