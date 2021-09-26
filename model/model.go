package model

type Person struct {
	Calendar
	Bound
}

type Calendar struct {
	Calendar [][]string
}

type Bound struct {
	Bound []string
}
