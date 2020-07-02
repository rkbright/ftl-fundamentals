// Package hello
package hello

// Declare greeting function
func ReturnGreeting(s string) string {
	if s == "Hi there" {
		return "Hi there yourself!"
	}
	return "Hi Gophers!"
}
