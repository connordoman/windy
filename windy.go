package windy

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/connordoman/windy/internal/palette"
)

// TailwindFamily is a type that represents a Tailwind V3 family (e.g. Slate, Gray, etc.).
type TailwindFamily string

// TailwindColor is a type that represents a Tailwind V3 color (RGB hexadecimal string).
type TailwindColor string

const (
	Slate   TailwindFamily = "Slate"
	Gray    TailwindFamily = "Gray"
	Zinc    TailwindFamily = "Zinc"
	Neutral TailwindFamily = "Neutral"
	Stone   TailwindFamily = "Stone"
	Red     TailwindFamily = "Red"
	Orange  TailwindFamily = "Orange"
	Amber   TailwindFamily = "Amber"
	Yellow  TailwindFamily = "Yellow"
	Lime    TailwindFamily = "Lime"
	Green   TailwindFamily = "Green"
	Emerald TailwindFamily = "Emerald"
	Teal    TailwindFamily = "Teal"
	Cyan    TailwindFamily = "Cyan"
	Sky     TailwindFamily = "Sky"
	Blue    TailwindFamily = "Blue"
	Indigo  TailwindFamily = "Indigo"
	Violet  TailwindFamily = "Violet"
	Purple  TailwindFamily = "Purple"
	Fuchsia TailwindFamily = "Fuchsia"
	Pink    TailwindFamily = "Pink"
	Rose    TailwindFamily = "Rose"
)

const (
	Slate50  TailwindColor = "#f8fafc"
	Slate100 TailwindColor = "#f1f5f9"
	Slate200 TailwindColor = "#e2e8f0"
	Slate300 TailwindColor = "#cbd5e1"
	Slate400 TailwindColor = "#94a3b8"
	Slate500 TailwindColor = "#64748b"
	Slate600 TailwindColor = "#475569"
	Slate700 TailwindColor = "#334155"
	Slate800 TailwindColor = "#1e293b"
	Slate900 TailwindColor = "#0f172a"
	Slate950 TailwindColor = "#020617"

	Gray50  TailwindColor = "#f9fafb"
	Gray100 TailwindColor = "#f3f4f6"
	Gray200 TailwindColor = "#e5e7eb"
	Gray300 TailwindColor = "#d1d5db"
	Gray400 TailwindColor = "#9ca3af"
	Gray500 TailwindColor = "#6b7280"
	Gray600 TailwindColor = "#4b5563"
	Gray700 TailwindColor = "#374151"
	Gray800 TailwindColor = "#1f2937"
	Gray900 TailwindColor = "#111827"
	Gray950 TailwindColor = "#030712"

	Zinc50  TailwindColor = "#fafafa"
	Zinc100 TailwindColor = "#f4f4f5"
	Zinc200 TailwindColor = "#e4e4e7"
	Zinc300 TailwindColor = "#d4d4d8"
	Zinc400 TailwindColor = "#a1a1aa"
	Zinc500 TailwindColor = "#71717a"
	Zinc600 TailwindColor = "#52525b"
	Zinc700 TailwindColor = "#3f3f46"
	Zinc800 TailwindColor = "#27272a"
	Zinc900 TailwindColor = "#18181b"
	Zinc950 TailwindColor = "#09090b"

	Neutral50  TailwindColor = "#fafafa"
	Neutral100 TailwindColor = "#f5f5f5"
	Neutral200 TailwindColor = "#e5e5e5"
	Neutral300 TailwindColor = "#d4d4d4"
	Neutral400 TailwindColor = "#a3a3a3"
	Neutral500 TailwindColor = "#737373"
	Neutral600 TailwindColor = "#525252"
	Neutral700 TailwindColor = "#404040"
	Neutral800 TailwindColor = "#262626"
	Neutral900 TailwindColor = "#171717"
	Neutral950 TailwindColor = "#0a0a0a"

	Stone50  TailwindColor = "#fafaf9"
	Stone100 TailwindColor = "#f5f5f4"
	Stone200 TailwindColor = "#e7e5e4"
	Stone300 TailwindColor = "#d6d3d1"
	Stone400 TailwindColor = "#a8a29e"
	Stone500 TailwindColor = "#78716c"
	Stone600 TailwindColor = "#57534e"
	Stone700 TailwindColor = "#44403c"
	Stone800 TailwindColor = "#292524"
	Stone900 TailwindColor = "#1c1917"
	Stone950 TailwindColor = "#0c0a09"

	Red50  TailwindColor = "#fef2f2"
	Red100 TailwindColor = "#fee2e2"
	Red200 TailwindColor = "#fecaca"
	Red300 TailwindColor = "#fca5a5"
	Red400 TailwindColor = "#f87171"
	Red500 TailwindColor = "#ef4444"
	Red600 TailwindColor = "#dc2626"
	Red700 TailwindColor = "#b91c1c"
	Red800 TailwindColor = "#991b1b"
	Red900 TailwindColor = "#7f1d1d"
	Red950 TailwindColor = "#450a0a"

	Orange50  TailwindColor = "#fff7ed"
	Orange100 TailwindColor = "#ffedd5"
	Orange200 TailwindColor = "#fed7aa"
	Orange300 TailwindColor = "#fdba74"
	Orange400 TailwindColor = "#fb923c"
	Orange500 TailwindColor = "#f97316"
	Orange600 TailwindColor = "#ea580c"
	Orange700 TailwindColor = "#c2410c"
	Orange800 TailwindColor = "#9a3412"
	Orange900 TailwindColor = "#7c2d12"
	Orange950 TailwindColor = "#431407"

	Amber50  TailwindColor = "#fffbeb"
	Amber100 TailwindColor = "#fef3c7"
	Amber200 TailwindColor = "#fde68a"
	Amber300 TailwindColor = "#fcd34d"
	Amber400 TailwindColor = "#fbbf24"
	Amber500 TailwindColor = "#f59e0b"
	Amber600 TailwindColor = "#d97706"
	Amber700 TailwindColor = "#b45309"
	Amber800 TailwindColor = "#92400e"
	Amber900 TailwindColor = "#78350f"
	Amber950 TailwindColor = "#451a03"

	Yellow50  TailwindColor = "#fefce8"
	Yellow100 TailwindColor = "#fef9c3"
	Yellow200 TailwindColor = "#fef08a"
	Yellow300 TailwindColor = "#fde047"
	Yellow400 TailwindColor = "#facc15"
	Yellow500 TailwindColor = "#eab308"
	Yellow600 TailwindColor = "#ca8a04"
	Yellow700 TailwindColor = "#a16207"
	Yellow800 TailwindColor = "#854d0e"
	Yellow900 TailwindColor = "#713f12"
	Yellow950 TailwindColor = "#422006"

	Lime50  TailwindColor = "#f7fee7"
	Lime100 TailwindColor = "#ecfccb"
	Lime200 TailwindColor = "#d9f99d"
	Lime300 TailwindColor = "#bef264"
	Lime400 TailwindColor = "#a3e635"
	Lime500 TailwindColor = "#84cc16"
	Lime600 TailwindColor = "#65a30d"
	Lime700 TailwindColor = "#4d7c0f"
	Lime800 TailwindColor = "#3f6212"
	Lime900 TailwindColor = "#365314"
	Lime950 TailwindColor = "#1a2e05"

	Green50  TailwindColor = "#f0fdf4"
	Green100 TailwindColor = "#dcfce7"
	Green200 TailwindColor = "#bbf7d0"
	Green300 TailwindColor = "#86efac"
	Green400 TailwindColor = "#4ade80"
	Green500 TailwindColor = "#22c55e"
	Green600 TailwindColor = "#16a34a"
	Green700 TailwindColor = "#15803d"
	Green800 TailwindColor = "#166534"
	Green900 TailwindColor = "#14532d"
	Green950 TailwindColor = "#052e16"

	Emerald50  TailwindColor = "#ecfdf5"
	Emerald100 TailwindColor = "#d1fae5"
	Emerald200 TailwindColor = "#a7f3d0"
	Emerald300 TailwindColor = "#6ee7b7"
	Emerald400 TailwindColor = "#34d399"
	Emerald500 TailwindColor = "#10b981"
	Emerald600 TailwindColor = "#059669"
	Emerald700 TailwindColor = "#047857"
	Emerald800 TailwindColor = "#065f46"
	Emerald900 TailwindColor = "#064e3b"
	Emerald950 TailwindColor = "#022c22"

	Teal50  TailwindColor = "#f0fdfa"
	Teal100 TailwindColor = "#ccfbf1"
	Teal200 TailwindColor = "#99f6e4"
	Teal300 TailwindColor = "#5eead4"
	Teal400 TailwindColor = "#2dd4bf"
	Teal500 TailwindColor = "#14b8a6"
	Teal600 TailwindColor = "#0d9488"
	Teal700 TailwindColor = "#0f766e"
	Teal800 TailwindColor = "#115e59"
	Teal900 TailwindColor = "#134e4a"
	Teal950 TailwindColor = "#042f2e"

	Cyan50  TailwindColor = "#ecfeff"
	Cyan100 TailwindColor = "#cffafe"
	Cyan200 TailwindColor = "#a5f3fc"
	Cyan300 TailwindColor = "#67e8f9"
	Cyan400 TailwindColor = "#22d3ee"
	Cyan500 TailwindColor = "#06b6d4"
	Cyan600 TailwindColor = "#0891b2"
	Cyan700 TailwindColor = "#0e7490"
	Cyan800 TailwindColor = "#155e75"
	Cyan900 TailwindColor = "#164e63"
	Cyan950 TailwindColor = "#083344"

	Sky50  TailwindColor = "#f0f9ff"
	Sky100 TailwindColor = "#e0f2fe"
	Sky200 TailwindColor = "#bae6fd"
	Sky300 TailwindColor = "#7dd3fc"
	Sky400 TailwindColor = "#38bdf8"
	Sky500 TailwindColor = "#0ea5e9"
	Sky600 TailwindColor = "#0284c7"
	Sky700 TailwindColor = "#0369a1"
	Sky800 TailwindColor = "#075985"
	Sky900 TailwindColor = "#0c4a6e"
	Sky950 TailwindColor = "#082f49"

	Blue50  TailwindColor = "#eff6ff"
	Blue100 TailwindColor = "#dbeafe"
	Blue200 TailwindColor = "#bfdbfe"
	Blue300 TailwindColor = "#93c5fd"
	Blue400 TailwindColor = "#60a5fa"
	Blue500 TailwindColor = "#3b82f6"
	Blue600 TailwindColor = "#2563eb"
	Blue700 TailwindColor = "#1d4ed8"
	Blue800 TailwindColor = "#1e40af"
	Blue900 TailwindColor = "#1e3a8a"
	Blue950 TailwindColor = "#172554"

	Indigo50  TailwindColor = "#eef2ff"
	Indigo100 TailwindColor = "#e0e7ff"
	Indigo200 TailwindColor = "#c7d2fe"
	Indigo300 TailwindColor = "#a5b4fc"
	Indigo400 TailwindColor = "#818cf8"
	Indigo500 TailwindColor = "#6366f1"
	Indigo600 TailwindColor = "#4f46e5"
	Indigo700 TailwindColor = "#4338ca"
	Indigo800 TailwindColor = "#3730a3"
	Indigo900 TailwindColor = "#312e81"
	Indigo950 TailwindColor = "#1e1b4b"

	Violet50  TailwindColor = "#f5f3ff"
	Violet100 TailwindColor = "#ede9fe"
	Violet200 TailwindColor = "#ddd6fe"
	Violet300 TailwindColor = "#c4b5fd"
	Violet400 TailwindColor = "#a78bfa"
	Violet500 TailwindColor = "#8b5cf6"
	Violet600 TailwindColor = "#7c3aed"
	Violet700 TailwindColor = "#6d28d9"
	Violet800 TailwindColor = "#5b21b6"
	Violet900 TailwindColor = "#4c1d95"
	Violet950 TailwindColor = "#2e1065"

	Purple50  TailwindColor = "#faf5ff"
	Purple100 TailwindColor = "#f3e8ff"
	Purple200 TailwindColor = "#e9d5ff"
	Purple300 TailwindColor = "#d8b4fe"
	Purple400 TailwindColor = "#c084fc"
	Purple500 TailwindColor = "#a855f7"
	Purple600 TailwindColor = "#9333ea"
	Purple700 TailwindColor = "#7e22ce"
	Purple800 TailwindColor = "#6b21a8"
	Purple900 TailwindColor = "#581c87"
	Purple950 TailwindColor = "#3b0764"

	Fuchsia50  TailwindColor = "#fdf4ff"
	Fuchsia100 TailwindColor = "#fae8ff"
	Fuchsia200 TailwindColor = "#f5d0fe"
	Fuchsia300 TailwindColor = "#f0abfc"
	Fuchsia400 TailwindColor = "#e879f9"
	Fuchsia500 TailwindColor = "#d946ef"
	Fuchsia600 TailwindColor = "#c026d3"
	Fuchsia700 TailwindColor = "#a21caf"
	Fuchsia800 TailwindColor = "#86198f"
	Fuchsia900 TailwindColor = "#701a75"
	Fuchsia950 TailwindColor = "#4a044e"

	Pink50  TailwindColor = "#fdf2f8"
	Pink100 TailwindColor = "#fce7f3"
	Pink200 TailwindColor = "#fbcfe8"
	Pink300 TailwindColor = "#f9a8d4"
	Pink400 TailwindColor = "#f472b6"
	Pink500 TailwindColor = "#ec4899"
	Pink600 TailwindColor = "#db2777"
	Pink700 TailwindColor = "#be185d"
	Pink800 TailwindColor = "#9d174d"
	Pink900 TailwindColor = "#831843"
	Pink950 TailwindColor = "#500724"

	Rose50  TailwindColor = "#fff1f2"
	Rose100 TailwindColor = "#ffe4e6"
	Rose200 TailwindColor = "#fecdd3"
	Rose300 TailwindColor = "#fda4af"
	Rose400 TailwindColor = "#fb7185"
	Rose500 TailwindColor = "#f43f5e"
	Rose600 TailwindColor = "#e11d48"
	Rose700 TailwindColor = "#be123c"
	Rose800 TailwindColor = "#9f1239"
	Rose900 TailwindColor = "#881337"
	Rose950 TailwindColor = "#4c0519"
)

// V3Families is a slice of all Tailwind V3 family names.
var V3Families = []TailwindFamily{
	Slate,
	Gray,
	Zinc,
	Neutral,
	Stone,
	Red,
	Orange,
	Amber,
	Yellow,
	Lime,
	Green,
	Emerald,
	Teal,
	Cyan,
	Sky,
	Blue,
	Indigo,
	Violet,
	Purple,
	Fuchsia,
	Pink,
	Rose,
}

// V3Colors is a slice of all Tailwind V3 colors.
var V3Colors = []TailwindColor{
	Slate50, Slate100, Slate200, Slate300, Slate400, Slate500, Slate600, Slate700, Slate800, Slate900, Slate950,
	Gray50, Gray100, Gray200, Gray300, Gray400, Gray500, Gray600, Gray700, Gray800, Gray900, Gray950,
	Zinc50, Zinc100, Zinc200, Zinc300, Zinc400, Zinc500, Zinc600, Zinc700, Zinc800, Zinc900, Zinc950,
	Neutral50, Neutral100, Neutral200, Neutral300, Neutral400, Neutral500, Neutral600, Neutral700, Neutral800, Neutral900, Neutral950,
	Stone50, Stone100, Stone200, Stone300, Stone400, Stone500, Stone600, Stone700, Stone800, Stone900, Stone950,
	Red50, Red100, Red200, Red300, Red400, Red500, Red600, Red700, Red800, Red900, Red950,
	Orange50, Orange100, Orange200, Orange300, Orange400, Orange500, Orange600, Orange700, Orange800, Orange900, Orange950,
	Amber50, Amber100, Amber200, Amber300, Amber400, Amber500, Amber600, Amber700, Amber800, Amber900, Amber950,
	Yellow50, Yellow100, Yellow200, Yellow300, Yellow400, Yellow500, Yellow600, Yellow700, Yellow800, Yellow900, Yellow950,
	Lime50, Lime100, Lime200, Lime300, Lime400, Lime500, Lime600, Lime700, Lime800, Lime900, Lime950,
	Green50, Green100, Green200, Green300, Green400, Green500, Green600, Green700, Green800, Green900, Green950,
	Emerald50, Emerald100, Emerald200, Emerald300, Emerald400, Emerald500, Emerald600, Emerald700, Emerald800, Emerald900, Emerald950,
	Teal50, Teal100, Teal200, Teal300, Teal400, Teal500, Teal600, Teal700, Teal800, Teal900, Teal950,
	Cyan50, Cyan100, Cyan200, Cyan300, Cyan400, Cyan500, Cyan600, Cyan700, Cyan800, Cyan900, Cyan950,
	Sky50, Sky100, Sky200, Sky300, Sky400, Sky500, Sky600, Sky700, Sky800, Sky900, Sky950,
	Blue50, Blue100, Blue200, Blue300, Blue400, Blue500, Blue600, Blue700, Blue800, Blue900, Blue950,
	Indigo50, Indigo100, Indigo200, Indigo300, Indigo400, Indigo500, Indigo600, Indigo700, Indigo800, Indigo900, Indigo950,
	Violet50, Violet100, Violet200, Violet300, Violet400, Violet500, Violet600, Violet700, Violet800, Violet900, Violet950,
	Purple50, Purple100, Purple200, Purple300, Purple400, Purple500, Purple600, Purple700, Purple800, Purple900, Purple950,
	Fuchsia50, Fuchsia100, Fuchsia200, Fuchsia300, Fuchsia400, Fuchsia500, Fuchsia600, Fuchsia700, Fuchsia800, Fuchsia900, Fuchsia950,
	Pink50, Pink100, Pink200, Pink300, Pink400, Pink500, Pink600, Pink700, Pink800, Pink900, Pink950,
	Rose50, Rose100, Rose200, Rose300, Rose400, Rose500, Rose600, Rose700, Rose800, Rose900, Rose950,
}

func (f TailwindFamily) String() string {
	return string(f)
}

func (c TailwindColor) String() string {
	return string(c)
}

// Glossy converts a Tailwind V3 color to a lipgloss.Color.
func (c TailwindColor) Glossy() lipgloss.Color {
	return lipgloss.Color(string(c))
}

// ForEachShadeByFamily iterates over each shade of each family and calls the provided function.
func ForEachShadeByFamily(fn func(family TailwindFamily, color TailwindColor, shade string)) {
	for f, family := range V3Families {
		for s := range 11 {
			fn(family, V3Colors[f*11+s], palette.ShadeOf(s))
		}
	}
}
