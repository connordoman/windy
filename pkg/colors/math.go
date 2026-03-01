package colors

import "math"

type vector []float64

func multiplyVectors(a, b vector) vector {
	return vector{
		a[0]*b[0] + a[1]*b[1] + a[2]*b[2],
		a[3]*b[0] + a[4]*b[1] + a[5]*b[2],
		a[6]*b[0] + a[7]*b[1] + a[8]*b[2],
	}
}

func radians(deg float64) float64 {
	return deg * math.Pi / 180
}

func degrees(rad float64) float64 {
	return rad * 180 / math.Pi
}

func sign(c float64) float64 {
	if c < 0 {
		return -1
	} else {
		return 1
	}
}

func mapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(value, max))
}
