package crypttest

type comparerCompletion func([]byte) bool

type PasswordComparerMock struct {
	CompleteWith comparerCompletion
}

func NewPasswordComparerMock(
	completion comparerCompletion,
) PasswordComparerMock {
	return PasswordComparerMock{
		CompleteWith: completion,
	}
}

// Implementing IPasswordComparer ---------------------------------------------

func (c PasswordComparerMock) MatchesHash(hash []byte) bool {
	return c.CompleteWith(hash)
}
