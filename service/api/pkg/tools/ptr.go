package tools

func StringPtr(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
