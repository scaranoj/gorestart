package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {

	//what would be a good test? Greet returns a string. What about an input error, like an int?
	//well it returns a string, so start by assigning it to something
	s := Greet("Bob")
	if s != "Real good there, Bob" {
		t.Error("Expected:", "Real good there, Bob", "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Lipshitz!"))
	//Output:
	//Real good there, Lipshitz!
}

func BenchmarkGreet(b *testing.B) {
	//old way
	// for i := 0; i < b.N; i++ {
	//new way
	for b.Loop() {
		Greet("Jack!")
	}
}
