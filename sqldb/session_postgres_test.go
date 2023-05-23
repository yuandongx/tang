package sqldb

import (
	"testing"
)

func TestSession_Query(t *testing.T) {
	type TangTest struct {
		A int    `type:"integer"`
		B string `type:"text"`
	}
	src := GetPostgresSource()
	t.Log("database source is : ", src)
	s := CreatePostgresSession(src)
	s.Create(TangTest{})
	d, e := s.Query("select * from tang_test;")
	t.Log(d)
	t.Log(e)
}
