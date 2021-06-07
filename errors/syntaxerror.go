package errors

type SyntaxError struct {
	error
	Message string
	Offset  int64
}

func NewSyntaxError(msg string, offset int64) *SyntaxError {
	return &SyntaxError{
		Message: msg,
		Offset:  offset,
	}
}

func (syntax *SyntaxError) Error() string {
	return syntax.Message
}
