package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/connordoman/windy"
	"github.com/connordoman/windy/pkg/windy4"
)

var (
	boxStyle       = lipgloss.NewStyle().Padding(1, 0, 2).Align(lipgloss.Center).Width((11 * 15)) //.Border(lipgloss.RoundedBorder())
	underlineStyle = lipgloss.NewStyle().Underline(true).Bold(true)
)

func textBox(text string) string {
	return boxStyle.Render(text)
}

func main() {

	v3Strs := []string{}
	v3Lines := []string{}
	windy.ForEachShadeByFamily(func(family windy.TailwindFamily, color windy.TailwindColor, shade string) {
		fg := windy.Neutral100.Glossy()
		if s, err := strconv.Atoi(shade); err == nil && s < 500 {
			fg = windy.Neutral700.Glossy()
		}
		str := lipgloss.NewStyle().Padding(1, 2).Background(color.Glossy()).Foreground(fg).Width(15).Height(3).Align(lipgloss.Center).Render(fmt.Sprintf("%s %s", family, shade))
		v3Strs = append(v3Strs, str)
		if shade == "950" {
			v3Lines = append(v3Lines, lipgloss.JoinHorizontal(lipgloss.Center, v3Strs...))
			v3Strs = []string{}
		}
	})
	fmt.Println(textBox(underlineStyle.Render("Tailwind V3")))
	fmt.Println(strings.Join(v3Lines, "\n"))

	fmt.Println()

	v4Strs := []string{}
	v4Lines := []string{}
	windy4.ForEachShadeByFamily(func(family windy4.TailwindFamily, color windy4.TailwindColor, shade string) {
		fg := windy4.Neutral100.Glossy()
		if s, err := strconv.Atoi(shade); err == nil && s < 500 {
			fg = windy4.Neutral700.Glossy()
		}
		str := lipgloss.NewStyle().Padding(1, 2).Background(color.Glossy()).Foreground(fg).Width(15).Height(3).Align(lipgloss.Center).Render(fmt.Sprintf("%s %s", family, shade))
		v4Strs = append(v4Strs, str)
		if shade == "950" {
			v4Lines = append(v4Lines, lipgloss.JoinHorizontal(lipgloss.Center, v4Strs...))
			v4Strs = []string{}
		}
	})
	fmt.Println(textBox(underlineStyle.Render("Tailwind V4") + "\n\n" + strings.Join(v4Lines, "\n")))

	overlaps := []string{}
	windy.ForEachShadeByFamily(func(family windy.TailwindFamily, color windy.TailwindColor, shade string) {
		windy4.ForEachShadeByFamily(func(v4Family windy4.TailwindFamily, v4Color windy4.TailwindColor, v4Shade string) {
			if string(family) == string(v4Family) && shade == v4Shade && string(color) == v4Color.ToHex() {
				fg := windy.Neutral100.Glossy()
				if s, err := strconv.Atoi(v4Shade); err == nil && s < 500 {
					fg = windy.Neutral700.Glossy()
				}
				colorized := lipgloss.NewStyle().Background(lipgloss.Color(v4Color.ToHex())).Foreground(fg).Padding(0, 1).Width(15).Align(lipgloss.Center).Render(fmt.Sprintf("%s%s", family, shade))
				overlaps = append(overlaps, colorized)
			}
		})
	})

	fmt.Println()

	overlapRow := []string{}
	overlapLines := []string{}
	for i, overlap := range overlaps {
		overlapRow = append(overlapRow, overlap)
		if (i+1)%11 == 0 {
			overlapLines = append(overlapLines, strings.Join(overlapRow, ""))
			overlapRow = []string{}
		}
	}
	fmt.Println(textBox(underlineStyle.Render(fmt.Sprintf("%d Overlaps (after oklch → hex)", len(overlaps))) + "\n\n" + strings.Join(overlapLines, "\n")))

}
