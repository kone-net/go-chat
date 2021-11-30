package errors

type error struct {
	msg string
}

func (e error) Error() string {
	return e.msg
}

func New(msg string) error {
	return error{
		msg: msg,
	}
}
