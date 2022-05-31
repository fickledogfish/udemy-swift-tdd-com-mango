package databasetest

import "fmt"

func ExampleNewIdentifiableMock() {
	mock := NewIdentifiableMock("some_string")

	fmt.Printf("%v", mock.Id())
	// Output: some_string
}
