package crypttest

type comparerCompletion func([]byte) bool

type passwordComparerMock struct {
	completeWith comparerCompletion
}

func NewPasswordComparerMock(
	completion comparerCompletion,
) passwordComparerMock {
	return passwordComparerMock{
		completeWith: completion,
	}
}

// Implementing IPasswordComparer ---------------------------------------------

func (c passwordComparerMock) MatchesHash(hash []byte) bool {
	return c.completeWith(hash)
}
