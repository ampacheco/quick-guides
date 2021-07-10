package main


//
// chapter 14 titled, Stacks and Pointer Mechanics, for more.

import (
	"fmt"
)

func main() {

	// Declare variable of type int with a value of 10.
	count := 10
	// To get the address of a value, use the & operator.

	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass a copy of the "value of" count (what’s in the box)
	// to the increment1 function.
	increment1(count)

	// Print out the "value of" and "address of" count.
	// The value of count will not change after the function call.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	// Pass a copy of the "address of" count (where is the box)
	// to the increment2 function. This is still considered a pass by
	// value and not a pass by reference because addresses are values.

	increment2(&count)
	// Print out the "value of" and "address of" count.
	// The value of count has changed after the function call.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}

// increment1 declares the function to accept its own copy of
// and integer value.
func increment1(inc int) {
	// Increment the local copy of the caller’s int value.
	inc++
	fmt.Println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 declares the function to accept its own copy of
// an address that points to an integer value.
// Pointer variables are literal types and are declared using *.
func increment2(inc *int) {
	// Increment the caller’s int value through the pointer.
	*inc++
	fmt.Println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tPoints To[", *inc, "]")
}
