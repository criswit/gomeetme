package main

import (
	"fmt"
	"gomeetme/calendar"
	"gomeetme/data"
)

var trackingIndex int

func main() {
	trackingIndex = 1
	availabilityMap := calendar.MakeAvailabilityMap()
	calendar1 := data.MakeCalendar1()
	calendar2 := data.MakeCalendar2()
	mergedCalendars := mergeCalendars(calendar1, calendar2)
	for _, entry := range mergedCalendars {
		updateAvailabilityMapFromCalendarBlock(availabilityMap, entry)
	}
	transitionPoints := getTransitionIndicies(availabilityMap);
	for _, point := range transitionPoints {
		fmt.Printf("transition point: %d\n", point)
	}
}

func getTransitionIndicies(availabilityMap map[int]calendar.Status) []int {
	transitionPoints := make([]int, 0)
	var endPoint int
	trackingIndex = 1
	for i := trackingIndex; i <= calendar.GetSecondsInDay(); i++ {
		if availabilityMap[i] == calendar.Available {
			i++
		}
		if availabilityMap[i] == calendar.NotAvailable {
			endPoint = i -1
			break
		}
	}
	transitionPoints = append(transitionPoints, trackingIndex, endPoint)
	return transitionPoints
}

func findAvailableBlocks(availabilityMap map[int]calendar.Status) [][]string {
	availabilityBlocks := make([][]string, 0)
	for i := trackingIndex; i <= calendar.GetSecondsInDay(); i++ {
		if availabilityMap[i] == calendar.Available {
			i++
		}
		if availabilityMap[i] == calendar.NotAvailable {
			calendarStart := calendar.ConvertSecondToStringTime(trackingIndex)
			calendarEnd := calendar.ConvertSecondToStringTime(i - 1)
			availabilityBlocks = append(availabilityBlocks, []string{calendarStart, calendarEnd})
			trackingIndex = i
		}
	}
	return availabilityBlocks

}

func updateAvailabilityMapFromCalendarBlock(availabilityMap map[int]calendar.Status, block []string) map[int]calendar.Status {
	startBlockSeconds := calendar.ConvertStringTimeToSecond(block[0])
	endBlockSeconds := calendar.ConvertStringTimeToSecond(block[1])
	fmt.Printf("current entry starts at %s seconds and ends at %s seconds", startBlockSeconds, endBlockSeconds)
	for i := startBlockSeconds; i <= endBlockSeconds; i++ {
		availabilityMap[i] = calendar.NotAvailable
	}
	return availabilityMap
}

func mergeCalendars(cal1 [][]string, cal2 [][]string) [][]string {
	var mergedCalendars [][]string
	for _, element := range cal1 {
		mergedCalendars = append(mergedCalendars, element)
	}
	for _, element := range cal2 {
		mergedCalendars = append(mergedCalendars, element)
	}
	return mergedCalendars
}