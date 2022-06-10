package testutils

type readFunc func([]byte) (int, error)
type closeFunc func() error

type ReadCloserMock struct {
	ReadWith  readFunc
	CloseWith closeFunc
}

// Implementing io.Reader -----------------------------------------------------

func (m ReadCloserMock) Read(data []byte) (int, error) {
	if m.ReadWith != nil {
		return m.ReadWith(data)
	} else {
		return 0, nil
	}
}

// Implementing io.Closer -----------------------------------------------------

func (m ReadCloserMock) Close() error {
	if m.CloseWith != nil {
		return m.CloseWith()
	} else {
		return nil
	}
}
