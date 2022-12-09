package scheduler

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	tm := time.Time{}
	t.Log(tm.IsZero())
}
