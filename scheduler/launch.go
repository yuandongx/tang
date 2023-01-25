package scheduler

import "time"

const (
	RUNNING = iota
	STOPPING
	PENDDING
	STOPED
)

type Runner struct {
	Status int
}

func (r *Runner) Run() {
	for {
		if r.Status == RUNNING {
			go run()
		} else {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func run() {
	for i, task := range AllTasks {
		print("Task ", i, task.Name, "is runnings!")
		task.Run()
	}
}
