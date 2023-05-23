package sqldb

import (
	"fmt"
	"testing"
)

func TestSession_Query(t *testing.T) {
	type TangTest struct {
		A int    `type:"integer"`
		B string `type:"text"`
	}
	src := GetPostGresSource()
	t.Log("database source is : ", src)
	s := CreatePostgresSession(src)
	i, e := s.Create(TangTest{})
	fmt.Println(i, e)
	d, e := s.Query("select * from public.tang_test;")
	t.Log(d)
	t.Log(e)
}
