package helper

import "math"

func GetOffset(page int, limit int) int {
	return (page - 1) * limit
}

func GetTotalPage(total int, limit int) int {
	if limit == 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(limit)))
}
