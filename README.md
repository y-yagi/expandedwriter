# Expandedwriter

Expandedwriter outputs data the same as  PostgreSQL's expanded table formatting.

## Example

```go
package main

import (
	"os"

	"github.com/y-yagi/expandedwriter"
)

func main() {
	w := expandedwriter.NewWriter(os.Stdout)

	w.SetFields([]string{"id", "email"})

	w.Append([]string{"1", "test1@example.com"})
	w.Append([]string{"2", "test2@example.com"})

	w.Render()

	// --[ Data 1 ]-------------
	// id    | 1
	// email | test1@example.com
	// --[ Data 2 ]-------------
	// id    | 2
	// email | test2@example.com
}
