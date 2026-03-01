package colors

import (
	"fmt"
	"math"
)

type RgbColor struct {
	R float64
	G float64
	B float64
}

func (c RgbColor) ToOklch() OklchColor {
	return c.ToSrgbLinear().ToXyz().ToOklab().ToOklch()
}

func rgbToSrgbLinear(c float64) float64 {
	if math.Abs(c) <= 0.04045 {
		return c / 12.92
	} else {
		return sign(c) * math.Pow((math.Abs(c)+0.055)/1.055, 2.4)
	}
}

func (c RgbColor) ToSrgbLinear() SrgbColor {
	return SrgbColor{
		R: rgbToSrgbLinear(c.R),
		G: rgbToSrgbLinear(c.G),
		B: rgbToSrgbLinear(c.B),
	}
}

func (c RgbColor) To255() RgbColor {
	return RgbColor{
		R: math.Round(clamp(c.R*255, 0, 255)),
		G: math.Round(clamp(c.G*255, 0, 255)),
		B: math.Round(clamp(c.B*255, 0, 255)),
	}
}

func (c RgbColor) Clamp() RgbColor {
	return RgbColor{
		R: clamp(c.R, 0, 1),
		G: clamp(c.G, 0, 1),
		B: clamp(c.B, 0, 1),
	}
}

func (c RgbColor) ToHex() string {
	// Treat channels in [0,1] as normalized RGB and convert to [0,255] first.
	if c.R <= 1 && c.G <= 1 && c.B <= 1 {
		c = c.To255()
	}

	return fmt.Sprintf(
		"#%02x%02x%02x",
		uint8(math.Round(clamp(c.R, 0, 255))),
		uint8(math.Round(clamp(c.G, 0, 255))),
		uint8(math.Round(clamp(c.B, 0, 255))),
	)
}
