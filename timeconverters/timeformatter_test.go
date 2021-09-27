package timeconverters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatTime(t *testing.T) {
	tests := []struct {
		name  string
		input int
		time  string
	}{
		{
			name:  "minimal",
			input: 3665,
			time:  "01:01:05",
		},
		{
			name:  "military",
			input: 72615,
			time:  "20:10:15",
		},
		{
			name:  "dad",
			input: 3600,
			time:  "01:00:00",
		},
	}
	for _, tt := range tests {
		got := FormatTime(tt.input)
		assert.Equal(t, tt.time, got)
	}
}
