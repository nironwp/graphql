package helpers

func PointerToString(s *string) string {
	if s == nil {
		return "" // return empty string if the pointer is nil
	}
	return *s // dereference the pointer to get the string value
}
