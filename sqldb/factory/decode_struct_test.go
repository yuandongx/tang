package factory

import (
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	//	type Bbject struct {
	//		A int
	//		B string
	//		C float64
	//		D bool
	//	}
	//	a := Bbject{1, "b", 3.121, false}
	//	b := &Bbject{1, "b", 3.121, false}
	//	Decode(a)
	//	//update(a, "a", 12)
	//	update(b, "a-b-d-13223", 12)
	//
	a := "asd-ASDF-sdsf"
	b := ""
	tmps := strings.Split(a, "-")
	for _, item := range tmps {
		tmp := make([]rune, len(item))
		for i, c := range item {
			if i == 0 && c >= 'a' && c <= 'z' {
				tmp[0] = c - 32
			} else {
				tmp[i] = c
			}
		}
		b += string(tmp)
	}
	t.Log(b)
}

func TestCreate(t *testing.T) {
	type A struct {
		A int    `type:"INT"`
		B bool   `type:"BOOL"`
		C string `type:"TEXT"`
	}
	a := A{A: 0, B: true, C: "CCC"}
	m := Parse(a, "mysql", "")
	s, _ := m.Create()
	t.Log(s)
}

func Test_lower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "azAZ", args: args{s: "azAZ"}, want: "az_a_z"},
		{name: "AdsfSDdfs", args: args{s: "AdsfSDdfs"}, want: "adsf_s_ddfs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lower(tt.args.s); got != tt.want {
				t.Errorf("lower() = %v, want %v", got, tt.want)
			}
		})
	}
}
