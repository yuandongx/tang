package sqldb

import (
	"fmt"
	"testing"
	"time"
)

func TestSession_CreatePostgresSession(t *testing.T) {
	type TangTest struct {
		A int
		B string
		C time.Time
		D bool
		E float32
		F float64
	}
	src := GetPostgresSource()
	t.Log("database source is : ", src)
	s := CreatePostgresSession(src)
	tt := TangTest{A: 1, B: "B", C: time.Now()}
	s.Create(tt)
	n, er := s.Save(tt)
	if n == 1 && er == nil {
		t.Log("Test Pass!")
	} else {
		t.Fatal("Test Failed!")
	}
	d, e := s.Query("select * from tang_test;")
	if len(d) > 0 && e == nil {
		t.Log(d)
		t.Log("Test Pass!")
	} else {
		t.Fatal("Test Failed!")
	}
	s.Init()
	s.Register(tt)
	data, err := s.Query(fmt.Sprintf("select * from %s_register;", s.name_space))
	t.Log(data)
	t.Log(err)
}
