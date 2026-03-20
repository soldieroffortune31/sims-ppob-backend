package helper

import "time"

func ConvertDateToUTCRange(dateStr string) (time.Time, time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")

	t, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).UTC()
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, loc).UTC()

	return start, end, nil
}
