package cnc

import "math"

func circleAreaFromDiameter(diameter float64) float64 {

	return 2 * math.Pi * (diameter / 2.0)
}
