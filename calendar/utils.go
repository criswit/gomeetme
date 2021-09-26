package calendar

// CreateAvailabilityMap returns a map for each second in a 24 hour period,
// with an associated boolean value representing if at least one of the calendars
// is occupied (not available). The map returned to this function sets each entry
// as true (available) on initialization.
func CreateAvailabilityMap() map[int]bool {
	availabilityMap := make(map[int]bool)
	for i := 0; i < 60*60*24; i++ {
		availabilityMap[i] = true
	}
	return availabilityMap
}

