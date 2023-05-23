package factory

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type Bbject struct {
		A int
		B string
		C float64
		D bool
	}
	a := Bbject{1, "b", 3.121, false}
	b := &Bbject{1, "b", 3.121, false}
	decode(a)
	decode(b)

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
