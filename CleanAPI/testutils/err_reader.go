package testutils

type readFunc func([]byte) (int, error)
type closeFunc func() error

type ReaderCloserMock struct {
	ReadWith  readFunc
	CloseWith closeFunc
}

func NewReaderCloserMock(
	readWith readFunc,
	closeWith closeFunc,
) ReaderCloserMock {
	return ReaderCloserMock{
		ReadWith:  readWith,
		CloseWith: closeWith,
	}
}

// Implementing io.Reader -----------------------------------------------------

func (m ReaderCloserMock) Read(data []byte) (int, error) {
	return m.ReadWith(data)
}

// Implementing io.Closer -----------------------------------------------------

func (m ReaderCloserMock) Close() error {
	return m.CloseWith()
}
