package scheduler

func Run() {
	s := NewScheduler()
	for {
		s.Run()
		wait(1)
	}
}
