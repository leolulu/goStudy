package helloword

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld()
	want := "Hello World!"

	if result != want {
		t.Errorf("HelloWorld() = %v want %v", result, want)
	}
}
