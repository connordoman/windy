# Windy

Tailwind CSS color definitions for your Go applications.

With support for Tailwind V3 and V4 colors, Windy makes it easy to choose your palette.

> [!NOTE]
> You can use both V3 and V4 colors directly in your full color terminal apps. Windy derives its colors directly from `tailwindcss/colors` and provides convenient format conversions.

## Windy :heart: [Lip Gloss](https://github.com/charmbracelet/lipgloss)

```go
style := lipgloss.NewStyle().Foreground(windy.Rose500.Glossy()).Bold(true)

fmt.Println(style.Render("It's like Windy is made for Lip Gloss"))
```

> [!IMPORTANT]
> When using V4 colors with Lip Gloss, `oklch` colors are converted to their closest hex counterpart. Using this method, there are 62 colors that overlap exactly between V3 and V4.

## Preview Windy

See how all the colors will look in your own terminal:

```shell
go run github.com/connordoman/windy/preview@latest
```

## Installation

```shell
go get github.com/connordoman/windy
```

## Tailwind V4

Since Windy is just a registry of strings, you can retrieve Tailind V4 colors exactly as they're used on the frontend. If your context doesn't support `oklch`, all V4 colors support best-match conversion to hex colors. All you have to do is import `windy4`:

```go
import "github.com/connordoman/windy/windy4"

mauve500 := windy4.Mauve500 // oklch(54.2% 0.034 322.5)

style := lipgloss.NewStyle().Background(
          mauve500.Glossy() // lipgloss.Color("#79697b")
        )
```

---

_Windy is not affiliated with or endorsed by [Tailwind Labs](https://github.com/tailwindlabs)._
