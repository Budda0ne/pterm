### table/demo

![Animation](https://vhs.charm.sh/vhs-4NMo5yXw7v0wtjOqrOeBgt.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the first table
	tableData1 := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

	pterm.Println() // Blank line

	// Define the data for the second table
	tableData2 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Create another table with a header and the defined data, then render it
	pterm.DefaultTable.WithHasHeader().WithData(tableData2).Render()
}
```

</details>

### table/alternate-row-style

![Animation](https://vhs.charm.sh/vhs-5dyBpzlQwOYMzISaI5fn97.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	// Each inner slice represents a row in the table.
	// The first row is considered as the header of the table.
	alternateStyle := pterm.NewStyle(pterm.BgDarkGray)

	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with the defined data.
	// The table has a header and is boxed.
	// Finally, render the table to print it.
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
}
```

</details>

### table/boxed

![Animation](https://vhs.charm.sh/vhs-7qXGEJfrGZ6fMTwxqk5W55.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	// Each inner slice represents a row in the table.
	// The first row is considered as the header of the table.
	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with the defined data.
	// The table has a header and is boxed.
	// Finally, render the table to print it.
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
}
```

</details>

### table/multiple-lines

![Animation](https://vhs.charm.sh/vhs-7dXD3ndtumyMAF9QJafxCK.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	data := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul\n\nNewline", "Dean", "augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "nunc.sed@est.com\nNewline"},
		{"Libby", "Camacho", "lobortis@semper.com"},
		{"张", "小宝", "zhang@example.com"},
	}

	// Create and render the table.
	// The options are chained in a single line for simplicity.
	// The table has a header, a row separator, and a header row separator.
	pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithHeaderRowSeparator("-").WithData(data).Render()
}
```

</details>

### table/right-alignment

![Animation](https://vhs.charm.sh/vhs-2mDOj18uXaOOtISlScg69v.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Define the data for the table.
	// Each inner slice represents a row in the table.
	// The first row is considered as the header.
	tableData := pterm.TableData{
		{"Firstname", "Lastname", "Email", "Note"},
		{"Paul", "Dean", "augue@velitAliquam.co.uk", ""},
		{"Callie", "Mckay", "nunc.sed@est.com", "这是一个测试, haha!"},
		{"Libby", "Camacho", "lobortis@semper.com", "just a test, hey!"},
		{"张", "小宝", "zhang@example.com", ""},
	}

	// Create a table with the defined data.
	// The table has a header and the text in the cells is right-aligned.
	// The Render() method is used to print the table to the console.
	pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(tableData).Render()
}
```

</details>

