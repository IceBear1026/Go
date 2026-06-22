package main

import (
	"fmt"
	"math"
)

/*
Interface let you wrote code based on behavior instead of specific struct type.

Without an interface, you might need separate functions:

func measureRect(r rect) {
    fmt.Println(r.area())
    fmt.Println(r.perim())
}

func measureCircle(c circle) {
    fmt.Println(c.area())
    fmt.Println(c.perim())
}

That gets repetitive.

With an interface, you can write one function:

func measure(g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

Now measure() can accept anything that passes the geometry checklist:

measure(rect{width: 3, height: 4})
measure(circle{radius: 5})
measure(triangle{base: 3, height: 4})

So the reason is not only “verify the struct.” It is more like:

“I want this function to accept anything that can do these behaviors.”

Interface lets me say: “I don’t care what struct you are. If your type has these required methods, I can use you here.”
*/

// we are creating a data type called geometry. It is an interface type. An interface type is a set of method signatures.
// interfaces are named collections of method signatures.
type geometry interface {
	// this means that any type that has these two methods (area and perim) is considered to implement the geometry interface.
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// we are defining the area method for rect. This means that rect implements the geometry interface because it has the area and perim methods defined.
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	// this is basically saying whatever type comes in, as long as it meets the geometry interface requirement, do these.
	// print its own field values.
	fmt.Println(g)
	// run its respective area and perim methods. This is polymorphism. The same function call can do different things based on the type of the object that it is called on.
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// *** GREAT way to check and also if IT IS TRUE do something
func detectCircle(g geometry) {
	// this if statement puts value of g into c and checks if it is of type circle. And just like we would check the data type for example for string we would do g.(string) or if we were checking integer g.(int). For this case we have a struct type called circle hence g.(circle). And if ok is true after ";" we print out this statement.
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

/*
The big idea is an interface describes behavior, not data.
So geometry does not care whether something is rectangle, circle, triangle, etc. It only cares:
- Can this thing calculate area?
- Can this thing calculate perimeter?

type geometry interface {
    area() float64
    perim() float64
}

So for our case, it means that WHAT methods a type must POSSESS before Go allows it to be used as geometry.

type geometry interface {
    area() float64
    perim() float64
}

means
To count as geometry, a type must have BOTH methods:
area() float64
AND
perim() float64

Think of it like a checklist.

Can rect be used as geometry?

Does rect have area() float64?  yes
Does rect have perim() float64? yes

Then rect satisfies geometry.

For circle:

Can circle be used as geometry?

Does circle have area() float64?  yes
Does circle have perim() float64? yes

Then circle satisfies geometry.

===
So when you call:
measure(r)

Go checks:
measure wants a geometry.
r is a rect.
Does rect have area() and perim()? yes
Okay, allow it.

*/
