package sqldb

import (
	"testing"
)

func TestSession_CreatePostgresSession(t *testing.T) {
	type TangTest struct {
		A int    `type:"integer"`
		B string `type:"text"`
	}
	src := GetPostgresSource()
	t.Log("database source is : ", src)
	s := CreatePostgresSession(src)
	tt := TangTest{A: 1, B: "B"}
	s.Create(tt)
	n, er := s.Save(tt)
	if n == 1 && er == nil {
		t.Log("Test Pass!")
	} else {
		t.Fatal("Test Failed!")
	}
	d, e := s.Query("select * from tang_test_int;")
	if len(d) > 0 && e == nil {
		t.Log(d)
		t.Log("Test Pass!")
	} else {
		t.Fatal("Test Failed!")
	}
}
