package scheduler

import "log"

type Any interface{}

type Handler = func(a ...Any) Any

var logger = log.Default()

func print(args ...Any) {
	logger.Println(args)
}
