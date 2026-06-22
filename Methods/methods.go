package main

import (
	"fmt"
)

// Go supports methods defined on struct types.

type rect struct {
	width, height int
}

// In Go, a method is just a function with a special parameter before the name:
// this (r rect) or (r *rect) pointe ror not either of these will make the function into a method. The parameter is called a receiver, and it allows the method to access the fields of the struct.

// means define a method called "area" that belongs to *rect
// the receiver is (r *rect), which means r is a pointer to a rect.
// so inside the method, r points to an existing rect.
/*
Important thing to notice here is that this method only reads with returning the r.width * r.height. It's not actually making assignments or modifying like: r.width = 100.

So the better explanation is:
area uses a pointer receiver.
That means r points to the original rect value instead of receiving a copy.
This method only reads the fields, but pointer receivers can also modify fields.
*/
func (r *rect) area() int {
	return r.width * r.height
}

// What does this mean without the pointer?
// Define a method called perim that belongs to rect. But for this one it creates a copy of the rect value. So we are not actually reading the value that is coming in.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// technically do I need to use a pointer for either of these input parameters?
// not sure. Don't they just take the value of each of the inputted variable?
// why is the first one has a pointer (area()) and why is perim() doesn't?

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

/*
This area method has a receiver type of *rect, which means it takes a pointer to a rect struct. This allows the method to modify the original struct in memory, rather than a copy of it.

Methods can be defined for either pointer or value receiver types.

You can also see that above we are calling 2 methods defined for our struct.
Usually, in C++, we would have a class or struct and a function defined inside of it. In Go, we can define a function outside of the struct and then use a receiver type to associate the function with the struct.

Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.

What does that mean by preventing or avodig modifying their methods? Is this similar to how we consider structs and classes where there are public and private where it will be allowed to modify the struct or not?

No. Public vs Private is more about capitalizing the name vs lowercase name.

Example:

type rect struct {
    width  int
    height int
}

Because rect, width, and height are lowercase, they are unexported. Other packages cannot access them directly.

But this:

type Rect struct {
    Width  int
    Height int
}

has exported names. Other packages can access them.

*/
