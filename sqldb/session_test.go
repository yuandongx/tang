package sqldb

import (
	"testing"
)

func TestSession_Query(t *testing.T) {

	s := CreateMySqlSession("mysql:P@ssw0rd@tcp(localhost:3306)/mysql")
	d, e := s.Query("select * from tang_test_123;")
	t.Log(d)
	t.Log(e)
}
