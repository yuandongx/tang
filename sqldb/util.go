package sqldb

type _error struct {
	e string
}

func (e _error) Error() string {
	return e.e
}
func error_(msg string) *_error {
	return &_error{e: msg}
}
