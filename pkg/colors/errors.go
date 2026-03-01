package colors

import "errors"

var (
	ErrInvalidOklchFormat = errors.New("invalid oklch format")
	ErrInvalidOklchColor  = errors.New("invalid oklch color")

	ErrInvalidRgbFormat = errors.New("invalid rgb format")
	ErrInvalidRgbColor  = errors.New("invalid rgb color")

	ErrInvalidHexFormat = errors.New("invalid hex format")
	ErrInvalidHexColor  = errors.New("invalid hex color")
)
