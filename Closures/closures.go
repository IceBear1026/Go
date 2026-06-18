/*
Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.

This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i, to form a closure.
*/
package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // nextInt is now a function that we can call to get the next integer in the sequence.

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()    // newInts is a new function that has its own closure over a different variable i, which starts at 0 again.
	fmt.Println(newInts()) // 1
}

/*
This is a really important Go concept. The confusing part is that this function is returning another function, not just returning a normal value like int or string.

*** Starting with function signature
func intSeq() func() int

intSeq is a function that returns another function.

And the returned function has this type func() int, which means it's a function that takes no arguments and returns an int.

So this func intSeq() func() int means:
intSeq takes no input
intSeq returns a function
That returned function takes no input and returns an int

*** Now what happens INSIDE intSeq
func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

Inside intSeq, we create a variable:
i := 0

Then we return an anonymous function:
return func() int {
	i++
	return i
}

this anonymous function does not have a name. It just says:
- when this function is called:
1. increment i by 1
2. return i

So the returned function is basically:
func() int {
	i++
	return i
}
*/

/*
If a local variable is still needed after the parent function returns, Go does NOT let it die with the stack frame. Go moves/preserves it so the returned function can keep using it.

What is the closure.
return func() int {
	i++
	return i
}

This means return a function value: func() int
However the parent function intSeq() runs ONCE here.
nextInt := intSeq()

And when this line runs, intSeq() executes:
i := 0
return func() int {
	i++
	return i
}
So nextInt now holds the returned function.
At this point intSeq is done.

But the returned function still references i.
So Go keeps i alive.

So then where does i live? Stack or Heap?
Normally local variables live in a function's stack frame:
func normal() int {
	i := 0
	i++
	return i
}
When normal() finishes and returns, it stack frame goes away, i dies.

But with this:
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
The inner function still needs i after intSeq() finishes and returns.
So Go says:
"I can not store i only in the temporary stack frame, because the returned function will need it later."
So Go may move i to the heap. This is called escape analysis.

nextInt ───> function
              |
              v
              captured i = 0

Then each time you call:

nextInt()
the function modifies the same captured i.

===

So the parent function intSeq() DOES close. Its stack frame is gone.
But the i variable does not disapper because the returned function captured it.

intSeq stack frame: gone
captured variable i: still alive
returned function: still alive because nextInt points to it

===
So does this cause stale memory?
Not in the dangerous C/C++ sense.
In C/C++, returning a pointer/reference to a local stack variable can be dangerous because the stack memory becomes invalid.

In Go, the compiler prevents that kind of unsafe behavior by moving the captured variable somewhere safe if needed.
*/
