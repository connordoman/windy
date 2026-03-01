package colors

import "math"

type XyzColor struct {
	X float64
	Y float64
	Z float64
}

func (c XyzColor) ToOklab() OklabColor {
	lms := multiplyVectors(
		vector{
			0.8190224379967030, 0.3619062600528904, -0.1288737815209879,
			0.0329836539323885, 0.9292868615863434, 0.0361446663506424,
			0.0481771893596242, 0.2642395317527308, 0.6335478284694309,
		},
		vector{c.X, c.Y, c.Z},
	)

	lmsG := vector{
		math.Cbrt(lms[0]),
		math.Cbrt(lms[1]),
		math.Cbrt(lms[2]),
	}

	lab := multiplyVectors(
		vector{
			0.2104542683093140, 0.7936177747023054, -0.0040720430116193,
			1.9779985324311684, -2.4285922420485799, 0.4505937096174110,
			0.0259040424655478, 0.7827717124575296, -0.8086757549230774,
		},
		lmsG,
	)

	return OklabColor{
		L: lab[0],
		A: lab[1],
		B: lab[2],
	}
}

func (c XyzColor) ToSrgbLinear() SrgbColor {
	rgb := multiplyVectors(
		vector{
			3.2409699419045226, -1.537383177570094, -0.4986107602930034,
			-0.9692436362808796, 1.8759675015077202, 0.04155505740717559,
			0.05563007969699366, -0.20397695888897652, 1.0569715142428786,
		},
		vector{c.X, c.Y, c.Z},
	)

	return SrgbColor{
		R: rgb[0],
		G: rgb[1],
		B: rgb[2],
	}
}
