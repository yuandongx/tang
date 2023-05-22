package sqldb

import (
	"testing"
)

func TestSession_Query(t *testing.T) {
	src := GetDbSource()
	s := CreateMySqlSession(src)
	d, e := s.Query("select * from tang_test_123;")
	t.Log(d)
	t.Log(e)
}
