package crypttest

type hasherCompletion func(string) ([]byte, error)

type passwordHasherMock struct {
	completeWith hasherCompletion
}

func NewPasswordHasherMock(completion hasherCompletion) passwordHasherMock {
	return passwordHasherMock{
		completeWith: completion,
	}
}

// Implementing IPasswordHasher -----------------------------------------------

func (h passwordHasherMock) Hash(password string) ([]byte, error) {
	return h.completeWith(password)
}
