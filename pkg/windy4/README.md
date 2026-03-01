# Windy for Tailwind V4

Windy also supports Tailwind V4. Colors are stored as CSS `oklch` strings, and can be coerced on-the-fly to RGB hex values. Just import `windy4`:

```go
import "github.com/connordoman/windy/windy4"

mauve500 := windy4.Mauve500 // oklch(54.2% 0.034 322.5)

style := lipgloss.NewStyle().Background(
          mauve500.Glossy() // lipgloss.Color("#79697b")
        )
```
