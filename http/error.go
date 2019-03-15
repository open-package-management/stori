package http

type FormatError struct {
	mediaType string
}

func (e *FormatError) Error() string {
	return "invalid mediatype"
}
