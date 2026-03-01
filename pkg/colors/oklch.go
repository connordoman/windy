package colors

import (
	"math"
	"strconv"
	"strings"
)

type OklchColor struct {
	L float64
	C float64
	H float64
}

func (c OklchColor) ToOklab() OklabColor {
	var a, b float64
	if math.IsNaN(c.H) {
		a = 0
		b = 0
	} else {
		a = c.C * math.Cos(radians(c.H))
		b = c.C * math.Sin(radians(c.H))
	}

	return OklabColor{
		L: c.L,
		A: a,
		B: b,
	}
}

func (c OklchColor) ToRgb() RgbColor {
	return c.ToOklab().ToXyz().ToSrgbLinear().ToRgb().Clamp()
}

// ParseOklch parses a standard CSS `oklch()` string and returns a pointer to an OklchColor.
func ParseOklch(s string) (*OklchColor, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.TrimPrefix(s, "oklch(")
	s = strings.TrimSuffix(s, ")")

	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return nil, ErrInvalidOklchFormat
	}

	lStr := strings.TrimSuffix(parts[0], "%")
	l, err := strconv.ParseFloat(lStr, 64)
	if err != nil {
		return nil, ErrInvalidOklchFormat
	}
	l /= 100

	c, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return nil, ErrInvalidOklchFormat
	}

	h, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, ErrInvalidOklchFormat
	}
	if h < 0 || h >= 360 {
		return nil, ErrInvalidOklchFormat
	}

	return &OklchColor{
		L: l,
		C: c,
		H: h,
	}, nil
}
