// Package gigasecond computes the moment after a gigasecond.
package gigasecond

import (
	"time"
)

// AddGigasecond determines the moment after a gigasecond
func AddGigasecond(t time.Time) time.Time {
	// challenge is what to do with 10^9 seconds and how to add the two values.
	// Found out there's an Add() function: https://pkg.go.dev/time#Time.Add
	// Could also be:
	// moment := float64(1000000000) equating to 1e+09
	// t.Add(1e+09 * time.Second) https://play.golang.org/p/C76o1EI_70Z

	return t.Add(1000000000 * time.Second)
}
