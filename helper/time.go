package helper

import "time"

func GetTimeUTCNow() time.Time {
	now := time.Now().UTC()
	return now
}
