### heatmap/demo

![Animation](https://vhs.charm.sh/vhs-3z0GW9eXD4Q5Dqpa9d7hCd.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap. Each sub-array represents a row in the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the labels for the X and Y axes of the heatmap.
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Create a heatmap with the defined data and axis labels, and enable RGB colors.
	// Then render the heatmap.
	pterm.DefaultHeatmap.WithAxisData(headerData).WithData(data).WithEnableRGB().Render()
}
```

</details>

### heatmap/custom_colors

![Animation](https://vhs.charm.sh/vhs-f1QlokaZY828iOfDUnPS7.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message
	pterm.Info.Println("The following table has no rgb (supported by every terminal), no axis data and a legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options, and render it
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(headerData).
		WithLegend(false).
		WithColors(pterm.BgBlue, pterm.BgRed, pterm.BgGreen, pterm.BgYellow).
		WithLegend().
		Render()
}
```

</details>

### heatmap/custom_legend

![Animation](https://vhs.charm.sh/vhs-1we9LwHEoQ8K1KSEpaigx4.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the header data for the heatmap
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a custom legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options
	// Options are chained in a single line for simplicity
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(headerData).
		WithEnableRGB().
		WithLegendLabel("custom").
		WithLegendOnlyColoredCells().
		Render() // Render the heatmap
}
```

</details>

### heatmap/custom_rgb

![Animation](https://vhs.charm.sh/vhs-5K4eVK7S1b0sZZiUO2R2JD.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap.
	axisLabels := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message.
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a legend.")
	pterm.Println()

	// Define the color range for the heatmap.
	rgbRange := []pterm.RGB{
		pterm.NewRGB(0, 0, 255),
		pterm.NewRGB(255, 0, 0),
		pterm.NewRGB(0, 255, 0),
		pterm.NewRGB(255, 255, 0),
	}

	// Create and render the heatmap.
	pterm.DefaultHeatmap.
		WithData(data).
		WithBoxed(false).
		WithAxisData(axisLabels).
		WithEnableRGB().
		WithRGBRange(rgbRange...).
		Render()
}
```

</details>

### heatmap/no_grid

![Animation](https://vhs.charm.sh/vhs-6lkwrF7EPJmYKCm9gJJPog.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis data for the heatmap.
	axisData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Print an informational message.
	pterm.Info.Println("The following table has rgb (not supported by every terminal), axis data and a legend.")
	pterm.Println()

	// Create the heatmap with the defined data and options, then render it.
	pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(axisData).WithEnableRGB().WithLegend().WithGrid(false).Render()
}
```

</details>

### heatmap/separated

![Animation](https://vhs.charm.sh/vhs-45oSPWmnHW2mOeCo3hCccL.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the heatmap.
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	// Define the axis labels for the heatmap.
	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
	}

	// Create the heatmap with the specified data and options, and render it.
	pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithLegend(false).Render()
}
```

</details>

