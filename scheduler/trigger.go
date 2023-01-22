package scheduler

import (
	"fmt"
	"strconv"
	"strings"
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

func (w WorkTime) String() string {
	return w.start.Format(time.RFC3339) + "," + w.stop.Format(time.RFC3339)
}

func LoadWorkTime(s string) WorkTime {
	w := WorkTime{}
	tmps := strings.Split(s, ",")
	if len(tmps) == 2 {
		if t, err := time.Parse(time.RFC3339, tmps[0]); err == nil {
			w.start = t
		}
		if t, err := time.Parse(time.RFC3339, tmps[1]); err == nil {
			w.stop = t
		}
	}
	return w
}

func getWorkTime(flag int) WorkTime {
	// 2006-01-02T15:04:05Z07:00
	now := time.Now()

	start := fmt.Sprintf("%d-%d-%dT09:25:00+08:00", now.Year(), now.Month(), now.Day())
	stop := fmt.Sprintf("%d-%d-%dT11:30:00+08:00", now.Year(), now.Month(), now.Day())

	if flag == 1 {
		start = fmt.Sprintf("%d-%d-%dT13:00:00+08:00", now.Year(), now.Month(), now.Day())
		stop = fmt.Sprintf("%d-%d-%dT15:00:00+08:00", now.Year(), now.Month(), now.Day())
	}

	tm1, err1 := time.Parse(time.RFC3339, start)
	tm2, err2 := time.Parse(time.RFC3339, stop)

	wt := WorkTime{}

	if err1 == nil {
		wt.start = tm1
	}

	if err2 == nil {
		wt.stop = tm2
	}

	return wt
}

// 上午工作时间
func Morning() WorkTime {
	return getWorkTime(0)
}

// 下午工作时间
func Afternoon() WorkTime {
	return getWorkTime(1)
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
	Index       string //索引
}

func NewTrigger() *Trigger {
	return &Trigger{}
}

func (t *Trigger) NextRunTime() time.Time {
	now := time.Now()
	if t.NextRuntime.After(now) {
		return t.NextRuntime
	}
	if t.LastRunTime.IsZero() {
		t.NextRuntime = now.Add(t.Interval)
		return t.NextRuntime
	}
	next := t.LastRunTime
	for {
		if next.After(now) {
			break
		} else {
			next = next.Add(t.Interval)
		}
	}
	t.NextRuntime = next
	return t.NextRuntime
}

// 将trigger 转换成map类型的数据以便存储
func (t *Trigger) JsonDump() map[string]string {
	result := make(map[string]string)
	result["Morning"] = t.Morning.String()
	result["Afternoon"] = t.Afternoon.String()
	result["Deviation"] = t.Deviation.String()
	result["LastRunTime"] = t.LastRunTime.Format(time.RFC3339)
	result["NextRuntime"] = t.NextRuntime.Format(time.RFC3339)
	result["Interval"] = t.Interval.String()
	result["SkippedDays"] = intArrayString(t.SkippedDays)
	if t.JustWorkDay {
		result["JustWorkDay"] = "true"
	} else {
		result["JustWorkDay"] = "false"
	}
	result["Index"] = t.Index
	return result

}

func intArrayString(d []int) string {
	s := ""
	for _, x := range s {
		s += fmt.Sprintf("%d", x)
	}
	return s
}

// 将从数据库中获取到的map类型的数据转换成Trigger
func (t *Trigger) JsonLoad(data map[string]string) *Trigger {
	trigger := &Trigger{}

	if v, ok := data["Morning"]; ok {
		trigger.Morning = LoadWorkTime(v)
	}

	if v, ok := data["Afternoon"]; ok {
		trigger.Afternoon = LoadWorkTime(v)
	}

	if v, ok := data["Deviation"]; ok {
		if i, err := time.ParseDuration(v); err == nil {
			trigger.Deviation = i
		}
	}

	if v, ok := data["LastRunTime"]; ok {
		if i, err := time.Parse(time.RFC3339, v); err == nil {
			trigger.LastRunTime = i
		}
	}

	if v, ok := data["NextRuntime"]; ok {
		if i, err := time.Parse(time.RFC3339, v); err == nil {
			trigger.NextRuntime = i
		}
	}

	if v, ok := data["Interval"]; ok {
		if i, err := time.ParseDuration(v); err == nil {
			trigger.Interval = i
		}
	}

	if skips, ok := data["SkippedDays"]; ok {
		ss := strings.Split(skips, ",")
		skipdays := make([]int, 0)
		for _, s := range ss {
			if i, err := strconv.Atoi(s); err == nil {
				skipdays = append(skipdays, i)
			}
		}
		trigger.SkippedDays = skipdays
	}

	if just, ok := data["JustWorkDay"]; ok {
		if just == "false" {
			trigger.JustWorkDay = false
		} else if just == "true" {
			trigger.JustWorkDay = true
		}
	}

	if index, ok := data["Index"]; ok {
		trigger.Index = index
	}
	return trigger
}

// 今天是有效的一天，即可工作的一天
// 今天既不是周末也不是节假日，即要跳过的日子
func (t *Trigger) TodayIsVaild() bool {
	now := time.Now()
	yday := now.YearDay()
	flag1 := now.Weekday() == time.Saturday || now.Weekday() == time.Sunday
	flag2 := false
	for _, day := range t.SkippedDays {
		flag2 = yday == day
		if flag2 {
			break
		}
	}
	return !flag1 && !flag2
}

// 更新工作时间
func (t *Trigger) UpdateWorkTime() {
	t.Morning = Morning()
	t.Afternoon = Afternoon()
}

// 设置路过时间
func (t *Trigger) SetSkips(skip []int) {
	t.SkippedDays = skip
}

// 设置时间偏差
func (t *Trigger) SetDeviation(td time.Duration) {
	t.Deviation = td
}

// 更新时间间隔
func (t *Trigger) SetInterval(ti time.Duration) {
	t.Interval = ti
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
