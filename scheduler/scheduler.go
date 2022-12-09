package scheduler

type Any interface{}

type Handler = func(a ...Any) Any

type Task struct {
	Handler
	Args []Any
}

func (t Task) Run() Any {
	return t.Handler(t.Args...)
}

type Scheduler struct {
	started  bool
	taskList []Task
}

func New() *Scheduler {
	return &Scheduler{
		started:  false,
		taskList: make([]Task, 0),
	}
}

func (s *Scheduler) Started() bool {
	return s.started
}

func (s *Scheduler) AddTask(t Task) {
	s.taskList = append(s.taskList, t)
}

func (s *Scheduler) GetTaskList() []Task {
	return s.taskList
}

func (s *Scheduler) Start() {
	for {
		for _, task := range s.taskList {
			go task.Run()
		}
	}
}
