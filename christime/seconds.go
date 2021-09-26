package christime

import (
	"fmt"
	"strconv"
)

const secondsInHour = 3600
const secondsInMinute = 60

// 11:43:00
// 01:00:00
func FormatTime(in int) string {
	hours, minutes, seconds := parseIntervals(in)
	return fmt.Sprintf("%s:%s:%s", addLeading0(int(hours)), addLeading0(int(minutes)), addLeading0(int(seconds)))

}

// parseIntervals is a function that ....
func parseIntervals(in int) (int64, int64, int64) {
	hours := int64(float64(in / secondsInHour))

	minutes := int64(float64((in % secondsInHour) / secondsInMinute))

	seconds := int64(in % secondsInHour % secondsInMinute)
	return hours, minutes, seconds
}

func addLeading0(in int) string {
	init := strconv.Itoa(in)
	if len(init) == 1 {
		return fmt.Sprintf("0%s", init)
	}
	return init
}