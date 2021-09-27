package timeconverters

import (
	"log"
	"strconv"
	"strings"
)

type TimeInterval string

const (
	Hour   TimeInterval = "hour"
	Minute TimeInterval = "minute"
	Second TimeInterval = "second"
)

// ConvertTimeToSecond is a function which converts a stringified representation of the time into its corresponding second value.
// This function expects the input string to be of the format HH::MM::SS, as well as being in military time. The program will panic if the
// hour is provided in an incorrect format, and we will fix this
// TODO: throw a more friendly error when we encounter an incorrectly formated input string instead of just panicing
func ConvertTimeToSecond(in string) int {
	split := strings.Split(in, ":")
	hours, mins, secs := parseSplit(split)
	log.Printf("hour: %s, minute %s, second: %s", hours, mins, secs)
	return inputToSecond(hours, Hour) + inputToSecond(mins, Minute) + inputToSecond(secs, Second)
}

func parseSplit(in []string) (hour string, min string, sec string) {
	return in[0], in[1], in[2]
}

func inputToSecond(in string, timeInterval TimeInterval) int {
	time, err := strconv.Atoi(in)
	if err != nil {
		log.Printf("got error %s", err)
		return -1
	}
	switch timeInterval {
	case Hour:
		return time * 60 * 60
	case Minute:
		return time * 60
	case Second:
		return time
	default:
		return -1

	}
}
