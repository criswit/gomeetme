package data

func MakeCalendar1() [][]string {
	return [][]string{
		{"10:00", "12:00"},
		{"12:30", "14:00"},
	}
}

func MakeCalendar2() [][]string {
	return [][]string{
		{"8:15", "9:30"},
		{"14:00", "16:00"},
	}

}

func MakeBound1() []string {
	return []string{"8:00", "14:30"}
}

func MakeBound2() []string {
	return []string{"8:15", "17:00"}

}