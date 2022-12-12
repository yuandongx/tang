package scheduler

import (
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

type Trigger struct {
	StartTime   time.Time
	StopTime    time.Time
	Deviation   time.Duration //时间偏差
	LastRunTime time.Time
	NextRuntime time.Time
	Interval    time.Duration
	SkippedDays []int // 那些天跳过
	JustWorkDay bool
}

// 获取下次执行时间
func (c *Trigger) GetNextRunTime() time.Time {
	if c.Expired() {
		return TIMEZERO
	}
	if c.Pending() {
		return c.StartTime
	}
	now := time.Now()

	// 跳过那些天
	s := len(c.SkippedDays) == 0 || !IntContain(c.SkippedDays, now.Day())

	return TIMEZERO
}

// 任务是否在有效期内
func (t *Trigger) Ongoing() bool {
	now := time.Now()
	// 0 < now 所以当前时间必须在start time 之后
	// 如果结束时间不为 0 则结束时间应在当前时间之后
	return t.StartTime.Before(now) && (t.StopTime.IsZero() || t.StopTime.After(now))
}

// 任务等待中
func (t *Trigger) Pending() bool {
	now := time.Now()
	return t.StartTime.IsZero() && now.Before(t.StartTime)
}

// 过期的
func (t *Trigger) Expired() bool {
	now := time.Now()
	return t.StopTime.IsZero() && now.After(t.StopTime)
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
