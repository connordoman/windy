package windy4

import (
	"fmt"
	"testing"

	"github.com/connordoman/windy/pkg/colors"
)

func TestWindy4OklchConversion(t *testing.T) {
	tests := []struct {
		name  string
		color TailwindColor
		hex   string
		want  colors.RgbColor
	}{
		{
			name:  "Indigo500",
			color: Indigo500,
			hex:   "#615fff",
			want:  colors.RgbColor{R: 97, G: 95, B: 255},
		},
		{
			name:  "Slate950",
			color: Slate950,
			hex:   "#020618",
			want:  colors.RgbColor{R: 2, G: 6, B: 24},
		},
		{
			name:  "Mauve500",
			color: Mauve500,
			hex:   "#79697b",
			want:  colors.RgbColor{R: 121, G: 105, B: 123},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oklch, err := colors.ParseOklch(string(test.color))
			if err != nil {
				t.Fatalf("failed to parse oklch color: %v", err)
			}
			rgb := oklch.ToRgb().To255()
			if rgb != test.want {
				t.Fatalf("got rgb color %v, want %v", rgb, test.want)
			}
			hex := oklch.ToRgb().ToHex()
			if hex != test.hex {
				t.Fatalf("got hex color %v, want %v", hex, test.hex)
			}
			t.Logf("got rgb color %v, hex color %v", rgb, hex)
		})
	}
}

func TestWindy4OklchParsing(t *testing.T) {

	ForEachShadeByFamily(func(family TailwindFamily, color TailwindColor, shade string) {
		c := fmt.Sprintf("%s%s", family, shade)
		t.Run(c, func(t *testing.T) {
			oklch, err := colors.ParseOklch(string(color))
			if err != nil {
				t.Fatalf("failed to parse oklch color: %v", err)
			}

			t.Logf("got oklch color %v", oklch)
		})
	})
}
