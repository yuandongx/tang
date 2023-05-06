package scheduler

import (
	"time"
	"fmt"
)

var TIMEZERO = time.Time{}

func TimeLoaction() *time.Location {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	return beijing
}

func wait(seconds int64) {
	d := time.Second * time.Duration(seconds)
	// time.Sleep(d)
	fmt.Println(d)
	time.Sleep(1)
}
