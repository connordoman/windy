# Windy

Tailwind CSS color definitions for your Go applications.

<img width="1439" height="1407" alt="Screenshot 2026-02-28 at 23 53 24" src="https://github.com/user-attachments/assets/ccee561c-d821-489f-91b1-e92544e0089b" />

Windy supports Tailwind V3 and V4, making it easier than ever to kickstart your app's theme.

> [!NOTE]
> You can use both V3 and V4 colors directly in your full color terminal apps. Windy derives its colors directly from `tailwindcss/colors` and provides convenient format conversions.

## Windy :heart: [Lip Gloss](https://github.com/charmbracelet/lipgloss)

```go
style := lipgloss.NewStyle().Foreground(windy.Rose500.Glossy()).Bold(true)

fmt.Println(style.Render("It's like Windy is made for Lip Gloss"))
```

> [!IMPORTANT]
> When using V4 colors with Lip Gloss, `oklch` colors are converted to their closest hex counterpart.

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

<img width="1439" height="1527" alt="Screenshot 2026-03-01 at 00 33 13" src="https://github.com/user-attachments/assets/f1ebc4fd-edc5-41e3-93fc-84705250efeb" />


Since Windy is just a registry of strings, you can retrieve Tailwind V4 colors exactly as they're used on the frontend. If your context doesn't support `oklch`, all V4 colors support best-match conversion to hex colors. All you have to do is import `windy4`:

```go
import "github.com/connordoman/windy/windy4"

mauve500 := windy4.Mauve500 // oklch(54.2% 0.034 322.5)

style := lipgloss.NewStyle().Background(
          mauve500.Glossy() // lipgloss.Color("#79697b")
        )
```

When using V4 colors with Lip Gloss (or converting to hex in general) there are 57 colors that overlap exactly between V3 and V4:

<img width="1437" height="290" alt="Screenshot 2026-03-01 at 00 55 19" src="https://github.com/user-attachments/assets/87356cd2-9a7f-4dda-beff-91912b1de62e" />


---

_Windy is not affiliated with or endorsed by [Tailwind Labs](https://github.com/tailwindlabs)._
