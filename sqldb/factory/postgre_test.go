package factory

import (
	"testing"
)

func TestPostgresModel_CreateCaseInt(t *testing.T) {
	t1 := TangTestCaseInt{}
	m1 := Parse(t1, POSTGRES, "")
	got, got1 := m1.Create()
	want1 := "CREATE TABLE public.tang_test_case_int(aint integer NOT NULL, bint16 smallint NOT NULL, cint32 integer NOT NULL, dint8 smallint NOT NULL, eint64 bigint NOT NULL, id serial, PRIMARY KEY(id));"
	if want1 == got {
		t.Log(got, got1)
	} else {
		t.Errorf("want1 %s got: %s", want1, got)
	}
}
func TestPostgresModel_CreateCaseFloate(t *testing.T) {
	case1 := TangTestCaseFloat{}
	m := Parse(case1, POSTGRES, "tang")
	got, got1 := m.Create()
	t.Log(got, got1)
}
