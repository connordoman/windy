package colors

import "math"

type OklabColor struct {
	L float64
	A float64
	B float64
}

func (c OklabColor) ToOklch() OklchColor {
	var h float64
	if math.Abs(c.A) < 0.0002 && math.Abs(c.B) < 0.0002 {
		h = math.NaN()
	} else {
		deg := degrees(math.Atan2(c.B, c.A))
		h = math.Mod(deg+360, 360)
	}
	return OklchColor{
		L: c.L,
		C: math.Sqrt(c.A*c.A + c.B*c.B),
		H: h,
	}
}

func (c OklabColor) ToXyz() XyzColor {
	lmsG := multiplyVectors(
		vector{
			1, 0.3963377773761749, 0.2158037573099136,
			1, -0.1055613458156586, -0.0638541728258133,
			1, -0.0894841775298119, -1.2914855480194092,
		},
		vector{c.L, c.A, c.B},
	)

	lms := vector{
		lmsG[0] * lmsG[0] * lmsG[0],
		lmsG[1] * lmsG[1] * lmsG[1],
		lmsG[2] * lmsG[2] * lmsG[2],
	}

	xyz := multiplyVectors(
		vector{
			1.2268798758459243, -0.5578149944602171, 0.2813910456659647,
			-0.0405757452148008, 1.1122868032803170, -0.0717110580655164,
			-0.0763729366746601, -0.4214933324022432, 1.5869240198367816,
		},
		lms,
	)

	return XyzColor{
		X: xyz[0],
		Y: xyz[1],
		Z: xyz[2],
	}
}
