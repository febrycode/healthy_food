package util

import "time"

const (
	WIB string = "Asia/Jakarta"
	UTC string = "UTC"
)

func GetTimeNow() time.Time {
	return time.Now().In(GetLocation())
}

// GetLocation - get location wib
func GetLocation() *time.Location {
	return time.FixedZone(WIB, 7*3600)
}
