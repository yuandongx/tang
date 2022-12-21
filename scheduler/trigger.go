package scheduler

import (
	"fmt"
	"time"
)

const (
	Sunday int = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type WorkTime struct {
	start time.Time
	stop  time.Time
}

// func getWorkTime(t string) time.Time {
// 	// 2006-01-02T15:04:05Z07:00
// 	tm, _ := time.Parse(time.RFC3339, t)
// 	tm, _ := time.Parse(time.RFC3339, t)
// 	return tm
// }

// s: 2006-01-02
func workTime(s string, flag int) WorkTime {
	start := fmt.Sprintf("%sT09:25:00Z08:00", s)
	stop := fmt.Sprintf("%sT11:30:00Z08:00", s)
	if flag == 1 {
		start = fmt.Sprintf("%sT13:00:00Z08:00", s)
		stop = fmt.Sprintf("%sT15:00:00Z08:00", s)
	}
	s1, e1 := time.Parse(time.RFC3339, start)
	s2, e2 := time.Parse(time.RFC3339, stop)
	w := WorkTime{}
	if e1 == nil {
		w.start = s1
	}
	if e2 == nil {
		w.start = s2
	}
	return w
}

type Trigger struct {
	Morning     WorkTime
	Afternoon   WorkTime
	Deviation   time.Duration //时间偏差
	LastRunTime time.Time
	NextRuntime time.Time
	Interval    time.Duration
	SkippedDays []int // 那些天跳过
	JustWorkDay bool
}

func NewTrigger() *Trigger {
	return &Trigger{}
}

func ToInt(w Any) (int, bool) {
	v, ok := w.(int)
	return v, ok
}

func IntContain(array []int, a int) bool {
	for _, v := range array {
		if v == a {
			return true
		}
	}
	return false
}
