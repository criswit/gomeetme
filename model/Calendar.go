package model

import (
	"gomeetme/timeconverters"
	"math"
	"time"
)

type Calendar struct {
	Meetings []Meeting `json:"meetings"`
}

// Meeting is a structure that represents a general meeting. Meetings have start times and end times,
// and as such those are respectively indicated at each pole of the []string array. Each Meeting can be
// can be assumed to be a list of length two, where the value at index 1 represents the start time of the meeting,
// and the value at start time two represents the end of hte meeting.
//TODO: handle scenarios where a meeting is incorrectly passed in with greater than two entries
type Meeting struct {
	Meeting []string `json:"meeting"`
}

type MeetingTime struct {
	TimeOfDay string `json:"time_of_day"`
}

func NewCalendar(meetings []Meeting) Calendar {
	return Calendar{Meetings: meetings}
}

func NewMeeting(meeting []string) Meeting {
	return Meeting{Meeting: meeting}
}

func NewMeetingStart(startTime string) MeetingTime {
	return MeetingTime{TimeOfDay: startTime}
}

func (m MeetingTime) CreateMeeting(duration time.Duration) Meeting {
	meetingSeconds := timeconverters.ConvertTimeToSecond(m.TimeOfDay)
	durationSeconds := int(math.Round(duration.Seconds()))
	endMeeting := timeconverters.FormatTime(meetingSeconds+durationSeconds)
	return Meeting{Meeting: []string{m.TimeOfDay, endMeeting}}

}
