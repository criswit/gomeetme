package manufacture

import (
	"github.com/stretchr/testify/assert"
	"gomeetme/model"
	"testing"
)

func TestMakePerson(t *testing.T) {
	tests := []struct {
		name           string
		inputCalender1 []string
		inputCalender2 []string
		inputCalender3 []string
		inputBound     []string
		expectedPerson *model.Person
	}{
		{
			name:           "basic",
			inputCalender1: []string{"10:00", "12:00"},
			inputCalender2: []string{"20:00", "22:00"},
			inputCalender3: []string{"10:00", "12:00"},
			inputBound:     []string{"9:00", "12:00"},
			expectedPerson: &model.Person{
				Calendar: model.Calendar{Calendar: [][]string{{"10:00", "12:00"}, {"20:00", "22:00"}, {"10:00", "12:00"}}},
				Bound:    model.Bound{Bound: []string{"9:00", "12:00"}},
			},
		},
	}
	for _, tt := range tests {
		got := MakePerson(tt.inputCalender1, tt.inputCalender2, tt.inputCalender3, tt.inputBound)
		assert.Equal(t, tt.expectedPerson, got)

	}
}
