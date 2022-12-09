package scheduler

import "time"

var TIMEZERO = time.Time{}

func TimeLoaction() *time.Location {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	return beijing
}
