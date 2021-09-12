package calendar

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const minute = "00:55"
const expectedSeconds = "3300"

func Test_ConvertMinutesToSeconds(t *testing.T) {
	actual := ConvertStringTimeToSecond(minute)
	assert.Equal(t, 3300, actual)

}

func Test_ConvertHoursToSeconds(t *testing.T) {
	time := "06:15"
	expectedSeconds, err  := strconv.Atoi("22500")
	if err != nil {
		panic(err)
	}
	actual := ConvertStringTimeToSecond(time)
	assert.Equal(t, expectedSeconds, actual)
}

func Test_ConvertSecondsToStringTime(t *testing.T)  {
	expectedTime := "6:15"
	actual := ConvertSecondToStringTime(29699)
	assert.Equal(t, expectedTime, actual)

}