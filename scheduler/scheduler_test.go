package scheduler

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tm := time.Time{}
	now := time.Now()
	t.Log(tm.IsZero())
	t.Log(tm.After(now))
}
