package factory

type AvailabilityMap struct {
	Status map[int]bool
}


// CreateAvailabilityMap returns a map for each second in a 24 hour period,
// with an associated boolean value representing if at least one of the calendars
// is occupied (not available). The map returned to this function sets each entry
// as true (available) on initialization.
func CreateAvailabilityMap() AvailabilityMap {
	availabilityMap := make(map[int]bool)
	for i := 0; i < 60*60*24; i++ {
		availabilityMap[i] = true
	}
	return AvailabilityMap{Status: availabilityMap}
}

func (a AvailabilityMap) FindOpenIntervals() [][]int {
	freeTimes := make([][]int, 0)
	iAnchor := 0
	for i := 0; i < len(a.Status); i++ {
		if a.Status[i] {
			if !a.Status[i-1] {
				iAnchor = i
			}
			if a.Status[i+1] {
				i++
			} else {
				if !a.Status[i+1] {}
				freeTimes = append(freeTimes, []int{iAnchor, i})
				continue
			}
		} else {
			continue
		}
	}
	return freeTimes
}


