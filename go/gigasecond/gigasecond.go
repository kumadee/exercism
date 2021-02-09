// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the time value after adding a gigasecond to the given time value.
func AddGigasecond(t time.Time) time.Time {
	// gigasecond defined in nanoseconds
	const gigasecond = time.Duration(1000000000000000000)
	return t.Add(gigasecond)
}