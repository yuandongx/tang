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
