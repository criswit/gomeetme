package christime

import (
	"log"
	"strconv"
	"strings"
)

type TimeInterval string

const (
	Hour TimeInterval = "hour"
	Minute TimeInterval = "minute"
	Second TimeInterval = "second"
)

// ConvertTimeToSecond will take an input string formatted like HH::MMM:SS and convert it to a second
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
		return time*60*60
	case Minute:
		return time*60
	case Second:
		return time
	default:
		return -1

	}
}