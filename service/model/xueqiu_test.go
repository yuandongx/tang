package com

import "testing"

func TestGetXueqiuRank(t *testing.T) {
	q := make(map[string]string)
	res, ok := GetXueqiuRank(q)
	t.Log(res, ok)
}
