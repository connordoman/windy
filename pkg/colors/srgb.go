package colors

import "math"

type SrgbColor struct {
	R float64
	G float64
	B float64
}

func (c SrgbColor) ToXyz() XyzColor {
	xyz := multiplyVectors(
		vector{
			0.41239079926595934, 0.357584339383878, 0.1804807884018343,
			0.21263900587151027, 0.715168678767756, 0.07219231536073371,
			0.01933081871559182, 0.11919477979462598, 0.9505321522496607,
		},
		vector{c.R, c.G, c.B},
	)

	return XyzColor{
		X: xyz[0],
		Y: xyz[1],
		Z: xyz[2],
	}
}

func srgbLinearToRgb(c float64) float64 {
	if math.Abs(c) > 0.0031308 {
		return sign(c) * (1.055*math.Pow(math.Abs(c), 1.0/2.4) - 0.055)
	} else {
		return c * 12.92
	}
}

func (c SrgbColor) ToRgb() RgbColor {
	return RgbColor{
		R: srgbLinearToRgb(c.R),
		G: srgbLinearToRgb(c.G),
		B: srgbLinearToRgb(c.B),
	}
}
