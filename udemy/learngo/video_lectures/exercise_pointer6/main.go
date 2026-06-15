package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("I am", d.first, "and I'm walking")
}

func (pd *dog) run() {
	pd.first = "Bob"
	fmt.Println("my name is", pd.first, "and I'm running")
}

type dogint interface {
	walk()
	run()
}

func dogfunc(i dogint) {
	i.run()
}

func main() {

	d1 := dog{"milo"}
	d1.run() // just calling the method by itself is fine. Also notice how it returned the first name of the dog that was in the func and
	// ignored the "milo" name above. This is because your dog non-pointer/value type got promoted to a pointer since it called the method directly
	//and therefore the pointer method "run()" was able to modify the value and return it. Not sure how it automatically got converted back to a
	//value type (i.e. automatic type promotion) but that's pretty cool magic.

	//dogfunc(d1) - this won't run. The linter warning language is chef's kiss: "cannot use d1 (variable of struct type dog) as dogint value in argument to dogfunc:
	// dog does not implement dogint (method run has pointer receiver)compilerInvalidIfaceAssign"

	d2 := &dog{"daisy"}
	dogfunc(d2) //the key is remembering to pass the pointer (i.e. hex address) of d2 into the func that expects a pointer address,
	// then the func body is calculated, which then calls and passes the d2 pointer off again to the run() method. Run method is attached to a pointer
	//(i.e. its receive is type pointer (*dog) so it will run the method as intended
	// By the way - you could've also passed a dereferenced variable of a dog type to dog func, but passing a derefernced type value makes more sense logically/semantically

	// d1.walk()
	// d1.run() //if we were calling the types methods directly, then this would work fine (promotion to pointer, per the Bodner book), though we are going
	// through an interface to show the limitations with value instances when they call pointer methods through func that take interface input - no worky

	// d2 := dog{"daisy"}
	// d2.run()

}
