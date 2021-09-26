package christime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputToSecond(t *testing.T) {
	tests := []struct {
		name   string
		TimeInterval TimeInterval
		input  string
		output int
		wantErr bool
	}{
		{
			name:   "minimal",
			TimeInterval: Hour,
			input:  "10",
			output: 36000,
			wantErr: false,
		},
		{
			name: "hour starts with 0",
			TimeInterval: Hour,
			input: "09",
			output: 32400,
			wantErr: false,
		},
		{
			name: "minute minimal",
			TimeInterval: Minute,
			input: "34",
			output: 2040,
			wantErr: false,
		},
		{
			name: "minute starts with 0",
			TimeInterval: Minute,
			input: "01",
			output: 60,
			wantErr: false,
		},
		{
			name: "second minimal",
			TimeInterval: Second,
			input: "00",
			output: 0,
			wantErr: false,
		},
		{
			name: "invalid string input",
			TimeInterval: Hour,
			input: "h5",
			output: -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inputToSecond(tt.input, tt.TimeInterval)
			if !tt.wantErr && tt.output < 0 {
				t.Errorf("test %s failed, wanted %d, got %d", tt.name, tt.output, got)
			}
			assert.Equal(t, got, tt.output)
		})
	}
}