package manufacture

import "gomeetme/model"

func MakePerson(cal1 []string, cal2 []string, cal3 []string, bound[]string) *model.Person {
	return &model.Person{
		Calendar: model.Calendar{[][]string{cal1, cal2, cal3}},
		Bound: model.Bound{Bound: []string{bound[0], bound[1]}},
	}
}