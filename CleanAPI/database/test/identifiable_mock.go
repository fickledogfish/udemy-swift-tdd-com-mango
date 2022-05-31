package databasetest

type IdentifiableMock struct {
	Identifier string
}

func NewIdentifiableMock(id string) IdentifiableMock {
	return IdentifiableMock{
		Identifier: id,
	}
}

// Implementing Identifiable --------------------------------------------------

func (im IdentifiableMock) Id() string {
	return im.Identifier
}
