package scheduler

import "encoding/json"

var (
	HandlerList map[string]Handler
	AllTasks    map[string]*Task
)

type Task struct {
	Name string
	Trigger
	Args   []Any
	Result Any
}

func init() {
	HandlerList = make(map[string]Handler)
	AllTasks = make(map[string]*Task)
}

func (t *Task) Run() {
	if t.Trigger.TimeIsUp() {
		if run, ok := HandlerList[t.Name]; ok {
			run(t.Args...)
		} else {
			print("==> No task named " + t.Name + "!")
		}
	} else {
		print("==> Time is not up to run the task <" + t.Name + ">!")
	}
}

func AddOrUpdateTask(t *Task, h Handler) {
	AllTasks[t.Name] = t
	HandlerList[t.Name] = h
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
	return dump
}
