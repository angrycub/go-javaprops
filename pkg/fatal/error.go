package fatal

type FatalError struct {
	err error
	ec  int
}

func NewError(err error, ec int) FatalError {
	return FatalError{
		err: err,
		ec:  ec,
	}
}

func (f FatalError) Unwrap() error {
	return f.err
}

func (f FatalError) Error() string {
	return f.err.Error()
}

func (f FatalError) ExitCode() int {
	return f.ec
}
