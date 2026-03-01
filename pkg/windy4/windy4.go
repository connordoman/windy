package windy4

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/connordoman/windy/internal/palette"
	"github.com/connordoman/windy/pkg/colors"
)

// TailwindFamily is a type that represents a Tailwind V4 family (e.g. Slate, Gray, etc.).
type TailwindFamily string

// TailwindColor is a type that represents a Tailwind V4 color (OKLCH string).
type TailwindColor string

const (
	Slate   TailwindFamily = "Slate"
	Gray    TailwindFamily = "Gray"
	Zinc    TailwindFamily = "Zinc"
	Neutral TailwindFamily = "Neutral"
	Stone   TailwindFamily = "Stone"
	Mauve   TailwindFamily = "Mauve"
	Olive   TailwindFamily = "Olive"
	Mist    TailwindFamily = "Mist"
	Taupe   TailwindFamily = "Taupe"
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
	Slate50  TailwindColor = "oklch(98.4% 0.003 247.858)"
	Slate100 TailwindColor = "oklch(96.8% 0.007 247.896)"
	Slate200 TailwindColor = "oklch(92.9% 0.013 255.508)"
	Slate300 TailwindColor = "oklch(86.9% 0.022 252.894)"
	Slate400 TailwindColor = "oklch(70.4% 0.04 256.788)"
	Slate500 TailwindColor = "oklch(55.4% 0.046 257.417)"
	Slate600 TailwindColor = "oklch(44.6% 0.043 257.281)"
	Slate700 TailwindColor = "oklch(37.2% 0.044 257.287)"
	Slate800 TailwindColor = "oklch(27.9% 0.041 260.031)"
	Slate900 TailwindColor = "oklch(20.8% 0.042 265.755)"
	Slate950 TailwindColor = "oklch(12.9% 0.042 264.695)"

	Gray50  TailwindColor = "oklch(98.5% 0.002 247.839)"
	Gray100 TailwindColor = "oklch(96.7% 0.003 264.542)"
	Gray200 TailwindColor = "oklch(92.8% 0.006 264.531)"
	Gray300 TailwindColor = "oklch(87.2% 0.01 258.338)"
	Gray400 TailwindColor = "oklch(70.7% 0.022 261.325)"
	Gray500 TailwindColor = "oklch(55.1% 0.027 264.364)"
	Gray600 TailwindColor = "oklch(44.6% 0.03 256.802)"
	Gray700 TailwindColor = "oklch(37.3% 0.034 259.733)"
	Gray800 TailwindColor = "oklch(27.8% 0.033 256.848)"
	Gray900 TailwindColor = "oklch(21% 0.034 264.665)"
	Gray950 TailwindColor = "oklch(13% 0.028 261.692)"

	Zinc50  TailwindColor = "oklch(98.5% 0 0)"
	Zinc100 TailwindColor = "oklch(96.7% 0.001 286.375)"
	Zinc200 TailwindColor = "oklch(92% 0.004 286.32)"
	Zinc300 TailwindColor = "oklch(87.1% 0.006 286.286)"
	Zinc400 TailwindColor = "oklch(70.5% 0.015 286.067)"
	Zinc500 TailwindColor = "oklch(55.2% 0.016 285.938)"
	Zinc600 TailwindColor = "oklch(44.2% 0.017 285.786)"
	Zinc700 TailwindColor = "oklch(37% 0.013 285.805)"
	Zinc800 TailwindColor = "oklch(27.4% 0.006 286.033)"
	Zinc900 TailwindColor = "oklch(21% 0.006 285.885)"
	Zinc950 TailwindColor = "oklch(14.1% 0.005 285.823)"

	Neutral50  TailwindColor = "oklch(98.5% 0 0)"
	Neutral100 TailwindColor = "oklch(97% 0 0)"
	Neutral200 TailwindColor = "oklch(92.2% 0 0)"
	Neutral300 TailwindColor = "oklch(87% 0 0)"
	Neutral400 TailwindColor = "oklch(70.8% 0 0)"
	Neutral500 TailwindColor = "oklch(55.6% 0 0)"
	Neutral600 TailwindColor = "oklch(43.9% 0 0)"
	Neutral700 TailwindColor = "oklch(37.1% 0 0)"
	Neutral800 TailwindColor = "oklch(26.9% 0 0)"
	Neutral900 TailwindColor = "oklch(20.5% 0 0)"
	Neutral950 TailwindColor = "oklch(14.5% 0 0)"

	Stone50  TailwindColor = "oklch(98.5% 0.001 106.423)"
	Stone100 TailwindColor = "oklch(97% 0.001 106.424)"
	Stone200 TailwindColor = "oklch(92.3% 0.003 48.717)"
	Stone300 TailwindColor = "oklch(86.9% 0.005 56.366)"
	Stone400 TailwindColor = "oklch(70.9% 0.01 56.259)"
	Stone500 TailwindColor = "oklch(55.3% 0.013 58.071)"
	Stone600 TailwindColor = "oklch(44.4% 0.011 73.639)"
	Stone700 TailwindColor = "oklch(37.4% 0.01 67.558)"
	Stone800 TailwindColor = "oklch(26.8% 0.007 34.298)"
	Stone900 TailwindColor = "oklch(21.6% 0.006 56.043)"
	Stone950 TailwindColor = "oklch(14.7% 0.004 49.25)"

	Mauve50  TailwindColor = "oklch(98.5% 0 0)"
	Mauve100 TailwindColor = "oklch(96% 0.003 325.6)"
	Mauve200 TailwindColor = "oklch(92.2% 0.005 325.62)"
	Mauve300 TailwindColor = "oklch(86.5% 0.012 325.68)"
	Mauve400 TailwindColor = "oklch(71.1% 0.019 323.02)"
	Mauve500 TailwindColor = "oklch(54.2% 0.034 322.5)"
	Mauve600 TailwindColor = "oklch(43.5% 0.029 321.78)"
	Mauve700 TailwindColor = "oklch(36.4% 0.029 323.89)"
	Mauve800 TailwindColor = "oklch(26.3% 0.024 320.12)"
	Mauve900 TailwindColor = "oklch(21.2% 0.019 322.12)"
	Mauve950 TailwindColor = "oklch(14.5% 0.008 326)"

	Olive50  TailwindColor = "oklch(98.8% 0.003 106.5)"
	Olive100 TailwindColor = "oklch(96.6% 0.005 106.5)"
	Olive200 TailwindColor = "oklch(93% 0.007 106.5)"
	Olive300 TailwindColor = "oklch(88% 0.011 106.6)"
	Olive400 TailwindColor = "oklch(73.7% 0.021 106.9)"
	Olive500 TailwindColor = "oklch(58% 0.031 107.3)"
	Olive600 TailwindColor = "oklch(46.6% 0.025 107.3)"
	Olive700 TailwindColor = "oklch(39.4% 0.023 107.4)"
	Olive800 TailwindColor = "oklch(28.6% 0.016 107.4)"
	Olive900 TailwindColor = "oklch(22.8% 0.013 107.4)"
	Olive950 TailwindColor = "oklch(15.3% 0.006 107.1)"

	Mist50  TailwindColor = "oklch(98.7% 0.002 197.1)"
	Mist100 TailwindColor = "oklch(96.3% 0.002 197.1)"
	Mist200 TailwindColor = "oklch(92.5% 0.005 214.3)"
	Mist300 TailwindColor = "oklch(87.2% 0.007 219.6)"
	Mist400 TailwindColor = "oklch(72.3% 0.014 214.4)"
	Mist500 TailwindColor = "oklch(56% 0.021 213.5)"
	Mist600 TailwindColor = "oklch(45% 0.017 213.2)"
	Mist700 TailwindColor = "oklch(37.8% 0.015 216)"
	Mist800 TailwindColor = "oklch(27.5% 0.011 216.9)"
	Mist900 TailwindColor = "oklch(21.8% 0.008 223.9)"
	Mist950 TailwindColor = "oklch(14.8% 0.004 228.8)"

	Taupe50  TailwindColor = "oklch(98.6% 0.002 67.8)"
	Taupe100 TailwindColor = "oklch(96% 0.002 17.2)"
	Taupe200 TailwindColor = "oklch(92.2% 0.005 34.3)"
	Taupe300 TailwindColor = "oklch(86.8% 0.007 39.5)"
	Taupe400 TailwindColor = "oklch(71.4% 0.014 41.2)"
	Taupe500 TailwindColor = "oklch(54.7% 0.021 43.1)"
	Taupe600 TailwindColor = "oklch(43.8% 0.017 39.3)"
	Taupe700 TailwindColor = "oklch(36.7% 0.016 35.7)"
	Taupe800 TailwindColor = "oklch(26.8% 0.011 36.5)"
	Taupe900 TailwindColor = "oklch(21.4% 0.009 43.1)"
	Taupe950 TailwindColor = "oklch(14.7% 0.004 49.3)"

	Red50  TailwindColor = "oklch(97.1% 0.013 17.38)"
	Red100 TailwindColor = "oklch(93.6% 0.032 17.717)"
	Red200 TailwindColor = "oklch(88.5% 0.062 18.334)"
	Red300 TailwindColor = "oklch(80.8% 0.114 19.571)"
	Red400 TailwindColor = "oklch(70.4% 0.191 22.216)"
	Red500 TailwindColor = "oklch(63.7% 0.237 25.331)"
	Red600 TailwindColor = "oklch(57.7% 0.245 27.325)"
	Red700 TailwindColor = "oklch(50.5% 0.213 27.518)"
	Red800 TailwindColor = "oklch(44.4% 0.177 26.899)"
	Red900 TailwindColor = "oklch(39.6% 0.141 25.723)"
	Red950 TailwindColor = "oklch(25.8% 0.092 26.042)"

	Orange50  TailwindColor = "oklch(98% 0.016 73.684)"
	Orange100 TailwindColor = "oklch(95.4% 0.038 75.164)"
	Orange200 TailwindColor = "oklch(90.1% 0.076 70.697)"
	Orange300 TailwindColor = "oklch(83.7% 0.128 66.29)"
	Orange400 TailwindColor = "oklch(75% 0.183 55.934)"
	Orange500 TailwindColor = "oklch(70.5% 0.213 47.604)"
	Orange600 TailwindColor = "oklch(64.6% 0.222 41.116)"
	Orange700 TailwindColor = "oklch(55.3% 0.195 38.402)"
	Orange800 TailwindColor = "oklch(47% 0.157 37.304)"
	Orange900 TailwindColor = "oklch(40.8% 0.123 38.172)"
	Orange950 TailwindColor = "oklch(26.6% 0.079 36.259)"

	Amber50  TailwindColor = "oklch(98.7% 0.022 95.277)"
	Amber100 TailwindColor = "oklch(96.2% 0.059 95.617)"
	Amber200 TailwindColor = "oklch(92.4% 0.12 95.746)"
	Amber300 TailwindColor = "oklch(87.9% 0.169 91.605)"
	Amber400 TailwindColor = "oklch(82.8% 0.189 84.429)"
	Amber500 TailwindColor = "oklch(76.9% 0.188 70.08)"
	Amber600 TailwindColor = "oklch(66.6% 0.179 58.318)"
	Amber700 TailwindColor = "oklch(55.5% 0.163 48.998)"
	Amber800 TailwindColor = "oklch(47.3% 0.137 46.201)"
	Amber900 TailwindColor = "oklch(41.4% 0.112 45.904)"
	Amber950 TailwindColor = "oklch(27.9% 0.077 45.635)"

	Yellow50  TailwindColor = "oklch(98.7% 0.026 102.212)"
	Yellow100 TailwindColor = "oklch(97.3% 0.071 103.193)"
	Yellow200 TailwindColor = "oklch(94.5% 0.129 101.54)"
	Yellow300 TailwindColor = "oklch(90.5% 0.182 98.111)"
	Yellow400 TailwindColor = "oklch(85.2% 0.199 91.936)"
	Yellow500 TailwindColor = "oklch(79.5% 0.184 86.047)"
	Yellow600 TailwindColor = "oklch(68.1% 0.162 75.834)"
	Yellow700 TailwindColor = "oklch(55.4% 0.135 66.442)"
	Yellow800 TailwindColor = "oklch(47.6% 0.114 61.907)"
	Yellow900 TailwindColor = "oklch(42.1% 0.095 57.708)"
	Yellow950 TailwindColor = "oklch(28.6% 0.066 53.813)"

	Lime50  TailwindColor = "oklch(98.6% 0.031 120.757)"
	Lime100 TailwindColor = "oklch(96.7% 0.067 122.328)"
	Lime200 TailwindColor = "oklch(93.8% 0.127 124.321)"
	Lime300 TailwindColor = "oklch(89.7% 0.196 126.665)"
	Lime400 TailwindColor = "oklch(84.1% 0.238 128.85)"
	Lime500 TailwindColor = "oklch(76.8% 0.233 130.85)"
	Lime600 TailwindColor = "oklch(64.8% 0.2 131.684)"
	Lime700 TailwindColor = "oklch(53.2% 0.157 131.589)"
	Lime800 TailwindColor = "oklch(45.3% 0.124 130.933)"
	Lime900 TailwindColor = "oklch(40.5% 0.101 131.063)"
	Lime950 TailwindColor = "oklch(27.4% 0.072 132.109)"

	Green50  TailwindColor = "oklch(98.2% 0.018 155.826)"
	Green100 TailwindColor = "oklch(96.2% 0.044 156.743)"
	Green200 TailwindColor = "oklch(92.5% 0.084 155.995)"
	Green300 TailwindColor = "oklch(87.1% 0.15 154.449)"
	Green400 TailwindColor = "oklch(79.2% 0.209 151.711)"
	Green500 TailwindColor = "oklch(72.3% 0.219 149.579)"
	Green600 TailwindColor = "oklch(62.7% 0.194 149.214)"
	Green700 TailwindColor = "oklch(52.7% 0.154 150.069)"
	Green800 TailwindColor = "oklch(44.8% 0.119 151.328)"
	Green900 TailwindColor = "oklch(39.3% 0.095 152.535)"
	Green950 TailwindColor = "oklch(26.6% 0.065 152.934)"

	Emerald50  TailwindColor = "oklch(97.9% 0.021 166.113)"
	Emerald100 TailwindColor = "oklch(95% 0.052 163.051)"
	Emerald200 TailwindColor = "oklch(90.5% 0.093 164.15)"
	Emerald300 TailwindColor = "oklch(84.5% 0.143 164.978)"
	Emerald400 TailwindColor = "oklch(76.5% 0.177 163.223)"
	Emerald500 TailwindColor = "oklch(69.6% 0.17 162.48)"
	Emerald600 TailwindColor = "oklch(59.6% 0.145 163.225)"
	Emerald700 TailwindColor = "oklch(50.8% 0.118 165.612)"
	Emerald800 TailwindColor = "oklch(43.2% 0.095 166.913)"
	Emerald900 TailwindColor = "oklch(37.8% 0.077 168.94)"
	Emerald950 TailwindColor = "oklch(26.2% 0.051 172.552)"

	Teal50  TailwindColor = "oklch(98.4% 0.014 180.72)"
	Teal100 TailwindColor = "oklch(95.3% 0.051 180.801)"
	Teal200 TailwindColor = "oklch(91% 0.096 180.426)"
	Teal300 TailwindColor = "oklch(85.5% 0.138 181.071)"
	Teal400 TailwindColor = "oklch(77.7% 0.152 181.912)"
	Teal500 TailwindColor = "oklch(70.4% 0.14 182.503)"
	Teal600 TailwindColor = "oklch(60% 0.118 184.704)"
	Teal700 TailwindColor = "oklch(51.1% 0.096 186.391)"
	Teal800 TailwindColor = "oklch(43.7% 0.078 188.216)"
	Teal900 TailwindColor = "oklch(38.6% 0.063 188.416)"
	Teal950 TailwindColor = "oklch(27.7% 0.046 192.524)"

	Cyan50  TailwindColor = "oklch(98.4% 0.019 200.873)"
	Cyan100 TailwindColor = "oklch(95.6% 0.045 203.388)"
	Cyan200 TailwindColor = "oklch(91.7% 0.08 205.041)"
	Cyan300 TailwindColor = "oklch(86.5% 0.127 207.078)"
	Cyan400 TailwindColor = "oklch(78.9% 0.154 211.53)"
	Cyan500 TailwindColor = "oklch(71.5% 0.143 215.221)"
	Cyan600 TailwindColor = "oklch(60.9% 0.126 221.723)"
	Cyan700 TailwindColor = "oklch(52% 0.105 223.128)"
	Cyan800 TailwindColor = "oklch(45% 0.085 224.283)"
	Cyan900 TailwindColor = "oklch(39.8% 0.07 227.392)"
	Cyan950 TailwindColor = "oklch(30.2% 0.056 229.695)"

	Sky50  TailwindColor = "oklch(97.7% 0.013 236.62)"
	Sky100 TailwindColor = "oklch(95.1% 0.026 236.824)"
	Sky200 TailwindColor = "oklch(90.1% 0.058 230.902)"
	Sky300 TailwindColor = "oklch(82.8% 0.111 230.318)"
	Sky400 TailwindColor = "oklch(74.6% 0.16 232.661)"
	Sky500 TailwindColor = "oklch(68.5% 0.169 237.323)"
	Sky600 TailwindColor = "oklch(58.8% 0.158 241.966)"
	Sky700 TailwindColor = "oklch(50% 0.134 242.749)"
	Sky800 TailwindColor = "oklch(44.3% 0.11 240.79)"
	Sky900 TailwindColor = "oklch(39.1% 0.09 240.876)"
	Sky950 TailwindColor = "oklch(29.3% 0.066 243.157)"

	Blue50  TailwindColor = "oklch(97% 0.014 254.604)"
	Blue100 TailwindColor = "oklch(93.2% 0.032 255.585)"
	Blue200 TailwindColor = "oklch(88.2% 0.059 254.128)"
	Blue300 TailwindColor = "oklch(80.9% 0.105 251.813)"
	Blue400 TailwindColor = "oklch(70.7% 0.165 254.624)"
	Blue500 TailwindColor = "oklch(62.3% 0.214 259.815)"
	Blue600 TailwindColor = "oklch(54.6% 0.245 262.881)"
	Blue700 TailwindColor = "oklch(48.8% 0.243 264.376)"
	Blue800 TailwindColor = "oklch(42.4% 0.199 265.638)"
	Blue900 TailwindColor = "oklch(37.9% 0.146 265.522)"
	Blue950 TailwindColor = "oklch(28.2% 0.091 267.935)"

	Indigo50  TailwindColor = "oklch(96.2% 0.018 272.314)"
	Indigo100 TailwindColor = "oklch(93% 0.034 272.788)"
	Indigo200 TailwindColor = "oklch(87% 0.065 274.039)"
	Indigo300 TailwindColor = "oklch(78.5% 0.115 274.713)"
	Indigo400 TailwindColor = "oklch(67.3% 0.182 276.935)"
	Indigo500 TailwindColor = "oklch(58.5% 0.233 277.117)"
	Indigo600 TailwindColor = "oklch(51.1% 0.262 276.966)"
	Indigo700 TailwindColor = "oklch(45.7% 0.24 277.023)"
	Indigo800 TailwindColor = "oklch(39.8% 0.195 277.366)"
	Indigo900 TailwindColor = "oklch(35.9% 0.144 278.697)"
	Indigo950 TailwindColor = "oklch(25.7% 0.09 281.288)"

	Violet50  TailwindColor = "oklch(96.9% 0.016 293.756)"
	Violet100 TailwindColor = "oklch(94.3% 0.029 294.588)"
	Violet200 TailwindColor = "oklch(89.4% 0.057 293.283)"
	Violet300 TailwindColor = "oklch(81.1% 0.111 293.571)"
	Violet400 TailwindColor = "oklch(70.2% 0.183 293.541)"
	Violet500 TailwindColor = "oklch(60.6% 0.25 292.717)"
	Violet600 TailwindColor = "oklch(54.1% 0.281 293.009)"
	Violet700 TailwindColor = "oklch(49.1% 0.27 292.581)"
	Violet800 TailwindColor = "oklch(43.2% 0.232 292.759)"
	Violet900 TailwindColor = "oklch(38% 0.189 293.745)"
	Violet950 TailwindColor = "oklch(28.3% 0.141 291.089)"

	Purple50  TailwindColor = "oklch(97.7% 0.014 308.299)"
	Purple100 TailwindColor = "oklch(94.6% 0.033 307.174)"
	Purple200 TailwindColor = "oklch(90.2% 0.063 306.703)"
	Purple300 TailwindColor = "oklch(82.7% 0.119 306.383)"
	Purple400 TailwindColor = "oklch(71.4% 0.203 305.504)"
	Purple500 TailwindColor = "oklch(62.7% 0.265 303.9)"
	Purple600 TailwindColor = "oklch(55.8% 0.288 302.321)"
	Purple700 TailwindColor = "oklch(49.6% 0.265 301.924)"
	Purple800 TailwindColor = "oklch(43.8% 0.218 303.724)"
	Purple900 TailwindColor = "oklch(38.1% 0.176 304.987)"
	Purple950 TailwindColor = "oklch(29.1% 0.149 302.717)"

	Fuchsia50  TailwindColor = "oklch(97.7% 0.017 320.058)"
	Fuchsia100 TailwindColor = "oklch(95.2% 0.037 318.852)"
	Fuchsia200 TailwindColor = "oklch(90.3% 0.076 319.62)"
	Fuchsia300 TailwindColor = "oklch(83.3% 0.145 321.434)"
	Fuchsia400 TailwindColor = "oklch(74% 0.238 322.16)"
	Fuchsia500 TailwindColor = "oklch(66.7% 0.295 322.15)"
	Fuchsia600 TailwindColor = "oklch(59.1% 0.293 322.896)"
	Fuchsia700 TailwindColor = "oklch(51.8% 0.253 323.949)"
	Fuchsia800 TailwindColor = "oklch(45.2% 0.211 324.591)"
	Fuchsia900 TailwindColor = "oklch(40.1% 0.17 325.612)"
	Fuchsia950 TailwindColor = "oklch(29.3% 0.136 325.661)"

	Pink50  TailwindColor = "oklch(97.1% 0.014 343.198)"
	Pink100 TailwindColor = "oklch(94.8% 0.028 342.258)"
	Pink200 TailwindColor = "oklch(89.9% 0.061 343.231)"
	Pink300 TailwindColor = "oklch(82.3% 0.12 346.018)"
	Pink400 TailwindColor = "oklch(71.8% 0.202 349.761)"
	Pink500 TailwindColor = "oklch(65.6% 0.241 354.308)"
	Pink600 TailwindColor = "oklch(59.2% 0.249 0.584)"
	Pink700 TailwindColor = "oklch(52.5% 0.223 3.958)"
	Pink800 TailwindColor = "oklch(45.9% 0.187 3.815)"
	Pink900 TailwindColor = "oklch(40.8% 0.153 2.432)"
	Pink950 TailwindColor = "oklch(28.4% 0.109 3.907)"

	Rose50  TailwindColor = "oklch(96.9% 0.015 12.422)"
	Rose100 TailwindColor = "oklch(94.1% 0.03 12.58)"
	Rose200 TailwindColor = "oklch(89.2% 0.058 10.001)"
	Rose300 TailwindColor = "oklch(81% 0.117 11.638)"
	Rose400 TailwindColor = "oklch(71.2% 0.194 13.428)"
	Rose500 TailwindColor = "oklch(64.5% 0.246 16.439)"
	Rose600 TailwindColor = "oklch(58.6% 0.253 17.585)"
	Rose700 TailwindColor = "oklch(51.4% 0.222 16.935)"
	Rose800 TailwindColor = "oklch(45.5% 0.188 13.697)"
	Rose900 TailwindColor = "oklch(41% 0.159 10.272)"
	Rose950 TailwindColor = "oklch(27.1% 0.105 12.094)"
)

var V4Families = []TailwindFamily{
	Slate,
	Gray,
	Zinc,
	Neutral,
	Stone,
	Mauve,
	Olive,
	Mist,
	Taupe,
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

var V4Colors = []TailwindColor{
	Slate50, Slate100, Slate200, Slate300, Slate400, Slate500, Slate600, Slate700, Slate800, Slate900, Slate950,
	Gray50, Gray100, Gray200, Gray300, Gray400, Gray500, Gray600, Gray700, Gray800, Gray900, Gray950,
	Zinc50, Zinc100, Zinc200, Zinc300, Zinc400, Zinc500, Zinc600, Zinc700, Zinc800, Zinc900, Zinc950,
	Neutral50, Neutral100, Neutral200, Neutral300, Neutral400, Neutral500, Neutral600, Neutral700, Neutral800, Neutral900, Neutral950,
	Stone50, Stone100, Stone200, Stone300, Stone400, Stone500, Stone600, Stone700, Stone800, Stone900, Stone950,
	Mauve50, Mauve100, Mauve200, Mauve300, Mauve400, Mauve500, Mauve600, Mauve700, Mauve800, Mauve900, Mauve950,
	Olive50, Olive100, Olive200, Olive300, Olive400, Olive500, Olive600, Olive700, Olive800, Olive900, Olive950,
	Mist50, Mist100, Mist200, Mist300, Mist400, Mist500, Mist600, Mist700, Mist800, Mist900, Mist950,
	Taupe50, Taupe100, Taupe200, Taupe300, Taupe400, Taupe500, Taupe600, Taupe700, Taupe800, Taupe900, Taupe950,
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

// Glossy first coerces a Tailwind V4 color to hex, and then returns a lipgloss.Color.
func (c TailwindColor) Glossy() lipgloss.Color {
	return lipgloss.Color(c.ToHex())
}

// ToHex coerces a Tailwind V4 color to a hex string.
func (c TailwindColor) ToHex() string {
	oklch, err := colors.ParseOklch(string(c))
	if err != nil {
		panic(err)
	}
	return oklch.ToRgb().ToHex()
}

// ToOklch parses a Tailwind V4 color as an OKLCH color struct.
func (c TailwindColor) ToOklch() *colors.OklchColor {
	oklch, err := colors.ParseOklch(string(c))
	if err != nil {
		panic(err) // does not occur for any V4 color
	}
	return oklch
}

// ForEachShadeByFamily iterates over each shade of each family and calls the provided function.
func ForEachShadeByFamily(fn func(family TailwindFamily, color TailwindColor, shade string)) {
	for f, family := range V4Families {
		for s := range 11 {
			fn(family, V4Colors[f*11+s], palette.ShadeOf(s))
		}
	}
}
