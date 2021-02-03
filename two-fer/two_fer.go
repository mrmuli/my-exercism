// Package twofer demonstrates Control flow.
package twofer

import "fmt"

// ShareWith returns a named output given a name.
func ShareWith(name string) string {
	var resp string
	switch {
	case len(name) > 0:
		resp = fmt.Sprintf("One for %v, one for me.", name)
	default:
		resp = fmt.Sprintf("One for you, one for me.")
	}
	return resp
}
