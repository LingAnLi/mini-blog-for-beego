package syserror

type Error interface {
	Code() int
	Error()string
	ReasonError()error
}

func NewErr(msg string,reason error) Error {
	return UnKnowError{msg,reason}
}