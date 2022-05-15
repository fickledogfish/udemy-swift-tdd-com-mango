package crypttest

type hasherCompletion func(string) ([]byte, error)

type passwordHasher struct {
	completeWith hasherCompletion
}

func NewPasswordHasher(completion hasherCompletion) passwordHasher {
	return passwordHasher{
		completeWith: completion,
	}
}

// Implementing IPasswordHasher -----------------------------------------------

func (h passwordHasher) Hash(password string) ([]byte, error) {
	return h.completeWith(password)
}
