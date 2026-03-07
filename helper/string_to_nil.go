package helper

func EmptyStringToNil(s *string) *string {
	if s == nil || *s == "" {
		return nil
	}
	return s
}
