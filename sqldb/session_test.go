package sqldb

import (
	"testing"
)

func TestSession_Query(t *testing.T) {

	s := CreateNySqlSession("mysql:P@ssw0rd@tcp(106.75.63.248:3306)/mysql")
	s.Query("select * from tang_test_123;")
}
