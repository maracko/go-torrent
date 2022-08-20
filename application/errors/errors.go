package errors

type SaveError struct {
	err string
}

type OpenError struct {
	err string
}

func NewSaveError(err string) SaveError {
	return SaveError{err}
}

func NewOpenError(err string) OpenError {
	return OpenError{err}
}

func (t SaveError) Error() string {
	return t.err
}

func (t OpenError) Error() string {
	return t.err
}
