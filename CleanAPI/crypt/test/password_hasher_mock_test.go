package crypttest

import "fmt"

func ExampleNewPasswordHasherMock() {
	passwordHasher := NewPasswordHasherMock(
		func(password string) ([]byte, error) {
			fmt.Printf("Password is: %s\n", password)
			return []byte{}, nil
		},
	)

	passwordHasher.Hash("some password")

	// Output: Password is: some password
}
