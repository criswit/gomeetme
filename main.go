package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"gomeetme/factory"
	"gomeetme/model"
	timeconv"gomeetme/timeconverters"
	"time"
)


func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Person", "Meetings")
	tbl2 := table.New("Availabile Slots")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, person := range initPersona() {
		tbl.AddRow(person.Name, person.Meetings)
	}
	avails := calcAvailabilities(initPersona())
	tbl2.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl2.AddRow(avails)

	tbl.Print()
	fmt.Sprintf("\n")
	tbl2.Print()
	//availabilityCalendar := factory.CreateAvailabilityMap()
	//duration := time.Duration(time.Hour)
	//meetingStart := model.NewMeetingStart(factory.GenerateRandomTime())
	//meeting1 := meetingStart.CreateMeeting(duration)
	//meeting2 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Minute *15)
	//calendar := model.NewCalendar([]model.Meeting{meeting1, meeting2})
	//for _, m := range calendar.Meetings {
	//	start := timeconverters.ConvertTimeToSecond(m.Meeting[0])
	//	end := timeconverters.ConvertTimeToSecond(m.Meeting[1])
	//	for i := start; i <= end; i++ {
	//		availabilityCalendar.Status[i] = false
	//	}
	//}
	//openSeconds := availabilityCalendar.FindOpenIntervals()
	//for i:=0; i<len(openSeconds);i++ {
	//	start := timeconverters.FormatTime(openSeconds[i][0])
	//	end := timeconverters.FormatTime(openSeconds[i][1])
	//	fmt.Printf("calendar is open from %s to %s\n", start, end)
	//}
	//fmt.Print(openSeconds)

}

func initPersona() []model.Person {
	people := make([]model.Person, 0)
	m1 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Minute * 45)
	m2 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Hour * 1)
	m3 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Minute * 500)
	m4 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Minute * 30)
	m5 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Hour * 1)
	m6 := model.NewMeetingStart(factory.GenerateRandomTime()).CreateMeeting(time.Minute * 500)
	p1 := model.Person{
		Name: "George",
		Calendar: model.Calendar{Meetings: []model.Meeting{m1, m2, m3}},
	}
	p2 := model.Person{
		Name: "Ted",
		Calendar: model.Calendar{Meetings: []model.Meeting{m4, m5, m6}},
	}
	people = append(people, p1, p2)
	return people
}

func calcAvailabilities(persons []model.Person) [][]string  {
	availabilityCalendar := make([][]string, 0)
	availabilities := factory.CreateAvailabilityMap()
	for _, person := range persons {
		for _, meeting := range person.Calendar.Meetings {
			timeBlock := meeting.Meeting
			for i := timeconv.ConvertTimeToSecond(timeBlock[0]); i < timeconv.ConvertTimeToSecond(timeBlock[1]); i++ {
				availabilities.Status[i] = false
			}
		}
	}

	availableSeconds := availabilities.FindOpenIntervals()
	for i := 0; i < len(availableSeconds); i++ {
		start := timeconv.FormatTime(availableSeconds[i][0])
		end := timeconv.FormatTime(availableSeconds[i][1])
		availabilityCalendar = append(availabilityCalendar, []string{fmt.Sprintf("\"%s\", \"%s\"", start, end)})
	}
	return availabilityCalendar
}