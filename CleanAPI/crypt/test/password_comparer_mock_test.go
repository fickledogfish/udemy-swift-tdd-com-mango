package crypttest

import "fmt"

func ExampleNewPasswordComparerMock() {
	passwordComparer := NewPasswordComparerMock(func(hash []byte) bool {
		fmt.Printf("Hash is: %s\n", string(hash))
		return false
	})

	passwordComparer.MatchesHash([]byte("some hash"))

	// Output: Hash is: some hash
}
