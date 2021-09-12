package calendar

import (
	"fmt"
	"strconv"
	"strings"
)

const hoursInDay = 24
const minutesInHour = 60
const secondsInMinute = 60

func GetSecondsInDay() int {
	return hoursInDay * minutesInHour * secondsInMinute
}

type Status = string

const (
	Available    Status = "available"
	NotAvailable Status = "booked"
)

func MakeAvailabilityMap() map[int]Status {
	availabilityMap := make(map[int]Status, GetSecondsInDay())
	for i := 1; i < GetSecondsInDay(); i++ {
		availabilityMap[i] = Available
	}
	return availabilityMap
}

func convertMinutesToSeconds(minute string) (int, error) {
	sec, err := strconv.Atoi(minute)
	if err != nil {
		fmt.Printf("got error converting minutes to seconds %s", err)
		return 0, err
	}
	return sec * secondsInMinute, nil
}

func convertHoursToSeconds(hour string) (int, error) {
	hr, err := strconv.Atoi(hour)
	if err != nil {
		fmt.Sprintf("got err converting hours to seconds %s", err)
		return 0, err
	}
	minutes := hr * minutesInHour
	return convertMinutesToSeconds(strconv.Itoa(minutes))

}

func createSecondToAvailableMap() map[int]bool {
	var secondToAvailabilityMap map[int]bool
	for i := 1; i < GetSecondsInDay(); i++ {
		secondToAvailabilityMap[i] = true
	}
	return secondToAvailabilityMap
}

func ConvertSecondToStringTime(second int) string {
	min := second / secondsInMinute
	tallier := 0
	var numberOfHours string
	var numberOfMinutes string
	for i := 1; i < 100000; i++ {
		tallier = tallier + minutesInHour
		if tallier == min {
			numberOfHours = strconv.Itoa(i)
			numberOfMinutes = "00"
			break
		} else if tallier > min {
			numberOfHours = strconv.Itoa(i - 1)
			minuteInt := 60 * (i -1)
			numberOfMinutes = strconv.Itoa(min - minuteInt)
			fmt.Printf("number of minutes %s", numberOfMinutes)
			break
		}
	}
	return fmt.Sprintf("%s:%s", numberOfHours, numberOfMinutes)

}

func ConvertStringTimeToSecond(time string) int {
	timeSplit := strings.Split(time, ":")
	hour := timeSplit[0]
	min := timeSplit[1]
	if hour == "00" {
		seconds, err := convertMinutesToSeconds(min)
		if err != nil {
			panic("nill")
		}
		fmt.Sprintf("created seconds in minutes")
		return seconds
	} else {
		hourFirstChar := hour[0]
		if string(hourFirstChar) == "0" {
			hour = string(hour[1])
		} else {
			hour = hour
		}
		int, err := convertMinutesToSeconds(min)
		if err != nil {
			fmt.Sprintf("failed to convert minutes to seconds %s", err)
		}
		hourInt, err := convertHoursToSeconds(hour)
		if err != nil {
			fmt.Sprintf("faield to convert hours to seconds %s", err)
		}
		return hourInt + int
	}
}