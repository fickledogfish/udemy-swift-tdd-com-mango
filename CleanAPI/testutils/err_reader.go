package testutils

type readFunc func([]byte) (int, error)
type closeFunc func() error

type ReadCloserMock struct {
	ReadWith  readFunc
	CloseWith closeFunc
}

func NewReadCloserMock(
	readWith readFunc,
	closeWith closeFunc,
) ReadCloserMock {
	return ReadCloserMock{
		ReadWith:  readWith,
		CloseWith: closeWith,
	}
}

// Implementing io.Reader -----------------------------------------------------

func (m ReadCloserMock) Read(data []byte) (int, error) {
	return m.ReadWith(data)
}

// Implementing io.Closer -----------------------------------------------------

func (m ReadCloserMock) Close() error {
	return m.CloseWith()
}
