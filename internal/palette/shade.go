package palette

import (
	"fmt"
)

func ShadeOf(index int) string {
	switch index {
	case 0:
		return "50"
	case 10:
		return "950"
	default:
		return fmt.Sprintf("%d00", index)
	}
}
