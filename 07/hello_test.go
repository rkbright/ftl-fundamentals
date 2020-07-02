// Declare package name
package hello_test

import (
	"fmt"
	"hello"
	"testing"
)

func TestHello(t *testing.T) {
	want1 := "Hi there yourself!"
	pass := "Hi there"
	got := hello.ReturnGreeting(pass)
	if want1 != got {
		t.Errorf("want: %q, got: %q", want1, got)
	}
	fmt.Println(got)
}
