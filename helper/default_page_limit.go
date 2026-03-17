package helper

func DefaultPageLimit(page *int, limit *int) {
	if *page <= 0 {
		*page = 1
	}

	if *limit <= 0 {
		*limit = 10
	}
}
