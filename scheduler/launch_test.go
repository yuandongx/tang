package scheduler

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	a := []Any{1, "1wq", 43.343, false}
	b, e := json.Marshal(a)
	fmt.Println(string(b))
	fmt.Println(e)
	var c []Any
	json.Unmarshal(b, &c)
	fmt.Println("1112111", c)
}

func TestRun(t *testing.T) {
	Run()
}
