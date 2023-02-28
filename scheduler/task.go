package scheduler

import (
	"encoding/json"
	"sync"
)

type Scheduler struct {
	Handlers map[string]Handler
	Tasks    sync.Map
}

type Task struct {
	Name string
	Trigger
	Args    []Any
	Result  Any
	Handler string
}

func NewScheduler() *Scheduler {
	return &Scheduler{Handlers: make(map[string]Handler), Tasks: sync.Map{}}
}

func (scheduler *Scheduler) Run() {
	scheduler.Tasks.Range(func(key, value any) bool {
		if task, ok := value.(Task); ok {
			if task.Trigger.TimeIsUp() {
				if run, ok := scheduler.Handlers[task.Handler]; ok {
					go run(task.Args...)
				} else {
					print("==> No task named " + task.Name + "!")
				}
			} else {
				print("==> Time is not up to run the task <" + task.Name + ">!")
			}
		}
		return true
	})

}

// 添加或者更新一个任务
func (scheduler *Scheduler) AddOrUpdate(t *Task) {
	scheduler.Tasks.Store(t.Name, t)
}

// 删除一个任务
func (scheduler *Scheduler) Delete(name string) (ok bool) {
	_, ok = scheduler.Tasks.LoadAndDelete(name)
	return
}

// 查找一个任务
func (scheduler Scheduler) Get(name string) (Task, bool) {
	if value, ok := scheduler.Tasks.Load(name); ok {
		task, flag := value.(Task)
		return task, flag
	}
	return Task{}, false
}

func LoadTask(data map[string]string) (t *Task) {
	t = &Task{}
	if name, ok := data["Name"]; ok {
		t.Name = name
	}
	if args, ok := data["Args"]; ok {
		var v []Any
		e := json.Unmarshal([]byte(args), &v)
		if e == nil {
			t.Args = v
		}
	}

	if h, ok := data["Handler"]; ok {
		t.Handler = h
	}

	trigger := Trigger{}
	trigger.JsonLoad(data)
	t.Trigger = trigger
	return
}
func (t *Task) Dump() map[string]string {
	dump := t.Trigger.JsonDump()
	if b, e := json.Marshal(t.Args); e == nil {
		dump["Args"] = string(b)
	}
	// dump["Args"] = strings.Join(t.Args, ",")
	dump["Name"] = t.Name
	dump["Handler"] = t.Handler
	return dump
}
