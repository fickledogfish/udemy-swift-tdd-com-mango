package crypttest

type hasherCompletion func(string) ([]byte, error)

type PasswordHasherMock struct {
	CompleteWith hasherCompletion
}

func NewPasswordHasherMock(completion hasherCompletion) PasswordHasherMock {
	return PasswordHasherMock{
		CompleteWith: completion,
	}
}

// Implementing PasswordHasher ------------------------------------------------

func (h PasswordHasherMock) Hash(password string) ([]byte, error) {
	return h.CompleteWith(password)
}
