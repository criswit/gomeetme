package main

import (
	"fmt"
	"gomeetme/factory"
	"gomeetme/model"
	"gomeetme/timeconverters"
	"time"
)

const hoursInDay = 24

func main() {
	availabilityCalendar := factory.CreateAvailabilityMap()
	duration := time.Duration(time.Hour)
	meetingStart := model.NewMeetingStart("10:30:00")
	meeting1 := meetingStart.CreateMeeting(duration)
	meeting2 := model.NewMeetingStart("12:30:00").CreateMeeting(time.Minute *15)
	calendar := model.NewCalendar([]model.Meeting{meeting1, meeting2})
	for _, m := range calendar.Meetings {
		start := timeconverters.ConvertTimeToSecond(m.Meeting[0])
		end := timeconverters.ConvertTimeToSecond(m.Meeting[1])
		for i := start; i <= end; i++ {
			availabilityCalendar.Status[i] = false
		}
	}
	openSeconds := availabilityCalendar.FindOpenIntervals()
	for i:=0; i<len(openSeconds);i++ {
		start := timeconverters.FormatTime(openSeconds[i][0])
		end := timeconverters.FormatTime(openSeconds[i][1])
		fmt.Printf("calendar is open from %s to %s\n", start, end)
	}
	fmt.Print(openSeconds)

}

