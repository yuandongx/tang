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

// 任务触发器接口
type TriggerInterface interface {
	NextRuntime() time.Time
	InRange() bool
	UpdateLastTime(t *time.Time)
}

// 时间范围
type TimeRange struct {
	start time.Time // 开始时间
	stop  time.Time // 结束时间
}

// 在时间范围内
func (t TimeRange) InRange() bool {
	now := time.Now()
	// 开始时间为 0 或者 开始时间之后 当前时间可能有效
	f1 := t.start.IsZero() || now.After(now) || now.Equal(t.start)
	// 结束时间为 0 或者 结束时间之前 当前时间可能有效
	f2 := t.stop.IsZero() || now.Before(now) || now.Equal(t.stop)
	// 开始 和结束时间同时有效 该时间范围才有效
	return f1 && f2
}

type Interval struct {
	TimeRange
	LastRunTime time.Time
	Interval    time.Duration // 时间间隔
	Deviation   time.Duration // 偏差容忍度
}

// i: 时间间隔
// d： 偏差容忍
// t: 有效期
func NewInterval(i time.Duration, d time.Duration, t TimeRange) *Interval {
	return &Interval{
		TimeRange:   t,
		LastRunTime: TIMEZERO,
		Interval:    i,
		Deviation:   d,
	}
}

func (i Interval) InRange() bool {
	return i.TimeRange.InRange()
}

func (d *Interval) UpdateLastTime(t time.Time) {
	d.LastRunTime = t
}
func (i *Interval) NextRuntime() time.Time {
	if i.InRange() {
		if i.LastRunTime.IsZero() {
			return time.Now()
		} else {
			return i.LastRunTime.Add(i.Interval)
		}
	}
	return TIMEZERO
}

type DateTime struct {
	RunAt       time.Time
	Deviation   time.Duration
	LastRunTime time.Time
}

func NewDateTime(runat time.Time, d time.Duration) *DateTime {
	return &DateTime{
		RunAt:       runat,
		Deviation:   d,
		LastRunTime: TIMEZERO,
	}
}

func (d DateTime) InRange() bool {
	now := time.Now()
	return d.RunAt.Sub(now) < d.Deviation
}

func (d *DateTime) UpdateLastTime(t time.Time) {
	d.LastRunTime = t
}

func (d DateTime) NextRuntime() time.Time {
	if !d.InRange() {
		return TIMEZERO
	}
	return d.RunAt
}

type Corn struct {
	TimeRange
	Years       []int
	Months      []int
	Days        []int
	Hours       []int
	Minutes     []int
	Seconds     []int
	Weekdays    []int
	Deviation   time.Duration
	LastRunTime time.Time
}

func (c Corn) InRange() bool {
	return c.TimeRange.InRange()
}

func (c *Corn) UpdateLastTime(t time.Time) {
	c.LastRunTime = t
}

func (c *Corn) NextRuntime() time.Time {
	now := time.Now()
	year := now.Year()
	month, _ := ToInt(now.Month())
	day := now.Day()
	weekday, _ := ToInt(now.Weekday())
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	y := len(c.Years) == 0 || IntContain(c.Years, year)
	m := len(c.Months) == 0 || IntContain(c.Months, month)
	d := len(c.Days) == 0 || IntContain(c.Days, day)
	w := len(c.Weekdays) == 0 || IntContain(c.Weekdays, weekday)
	h := len(c.Hours) == 0 || IntContain(c.Hours, hour)
	// n := len(c.Minutes) == 0 || IntContain(c.Minutes, minute)

	if y && m && d && w && h {
		startMinut := 0
		if len(c.Minutes) > 0 {
			for _, m := range c.Minutes {
				if minute <= m {
					startMinut = m
					break
				}
			}
			if startMinut == 0 {
				startMinut = c.Minutes[0]
			}
		}
		startSecond := 0
		if len(c.Seconds) > 0 {

			for _, s := range c.Seconds {
				if second < s {
					startSecond = s
					break
				}
			}
			if startSecond == 0 {
				startSecond = c.Seconds[0]
			}

		}
		return time.Date(year, now.Month(), day, hour, startMinut, startSecond, 0, TimeLoaction())
	}
	return TIMEZERO
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
