package model

// Bound represents ranges of times during which an individual IS or IS NOT willing to accept meetings.
// A time that falls within the range of a person's NotWillingToAccept bound is equivalent to them already having been
// booked in another meeting. They are not available.
type Bound struct {
	Bound []string
}
