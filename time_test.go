package example

import (
	"testing"
	"time"
)

func TestCustomDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2020, 11, 23, 10, 0, 0, 0, time.UTC),
			want: "20 jun 2020 at 12:00",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "ticket event",
			tm:   time.Date(2020, 12, 25, 10, 0, 0, 0, time.FixedZone("ticket event", 3*60*60)),
			want: "20 jun 2020 at 12:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := customDate(tt.tm)
			if got != tt.want {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}
