package phonenumbers

type NumberParseError struct {
	msg string
}

func (e *NumberParseError) Error() string {
	return e.msg
}
