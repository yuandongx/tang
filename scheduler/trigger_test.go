package scheduler

import (
	"testing"
	"time"
)

func TestTrigger(t *testing.T) {
	trigger := Trigger{
		Morning:     LoadWorkTime("2023-01-23T09:25:00+08:00,2023-01-23T11:25:00+08:00"),
		Afternoon:   LoadWorkTime("2023-01-23T13:00:00+08:00,2023-01-23T15:25:00+08:00"),
		Deviation:   1 * time.Second,
		LastRunTime: time.Time{},
		NextRuntime: time.Time{},
		Interval:    10 * time.Second,
		SkippedDays: []int{23, 34, 45},
		JustWorkDay: false,
		Index:       "12-1-2-34-5",
	}
	next := trigger.NextRunTime()
	t.Log(next)
	ts := trigger.JsonDump()
	t.Log(ts)
	w := LoadWorkTime("2023-01-23T09:25:00+08:00,2023-01-23T11:25:00+08:00")
	t.Log(w.String())
}
