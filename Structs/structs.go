package main

import (
	"fmt"
)

// Go does not have class and only have struct. Struct is a collection of fields.
// "type" is the keyword to define a new struct type.
// we are naving this struct type "person" with two fields: name and age.
// this does not create a new person struct, it just creates a type/blueprint.
type person struct {
	name string
	age  int
}

// newPerson is a function that takes a name as an argument (string) and returns a pointer to a person struct.
// what does it mean to send a pointer to a struct? It means that we are sending the address of the struct in memory, rather than a copy of the struct itself. This is useful when we want to modify the original struct.
func newPerson(name string) *person {
	// we are creating a new person struct and initializing the name field with the value of the name argument.
	// we are then setting the age field to 42.

	// These are not fundamentally different. The code is just showing two different ways to set struct fields.
	// p := person{name: name}
	// means create a person struct. Set its name field to the value of the name parameter. Leave age as the zero value for now.

	// p.age = 42
	// this changes 0 to 42.

	/*
		Then we end up with:
		person{
			name: "some name",
			age:  42,
		}

		I could have just done this as well:

		func newPerson(name string) *person {
		p := person{
			name: name,
			age:  42,
		}

		return &p
		}

		I could even do this:
		func newPerson(name string) *person {
			return &person{
				name: name,
				age:  42,
			}
		}
	*/

	p := person{}
	fmt.Println(p.name) // ""
	fmt.Println(p.age)  // 0
	// Also another important thing is that I had to instantiate the struct first before I can set the age field as well.
	p.name = name
	p.age = 42

	// why is the pointer used with * but when we dereference it why do we use &?
	// the * is used to define a pointer type, while the & is used to get the address of a variable.

	fmt.Println(p.name)
	fmt.Println(p.age)
	return &p
}

func main() {
	// we are creating a copy of Person Struct and giving is value.
	// we don't need to specify the value as long as it's in the correct order of the struct fields.
	fmt.Println(person{"Bob", 20}) // {Bob 20}

	// we are explicitly giving it the actual struct field names which is another way to create a copy of Person Struct.
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

	// we are creating another copy of the person struct type but this time only mentioning the name and the other value will just be 0 because it's not given a value.
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	// this means to create a new person value with name Ann and age 40. Then return its address.
	/*
		So this:
		person{name: "Ann", age: 40}
		creates the struct value.

		And this:
		&person{name: "Ann", age: 40}
		creates the struct value and gives you a pointer to it.

		Output:
		&{Ann 40}

		It just creates one temporary person value and prints its address/value.
	*/
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	fmt.Println(newPerson("Jon")) // &{Jon 42}

	// & means: give me the address of this value.
	// So: &p means address of p.

	/*
		Why use pointer then?

		A pointer lets you refer to the same struct value without copying the whole thing.

		Example:

		p1 := person{name: "Ann", age: 40}
		p2 := p1

		p2.age = 99

		fmt.Println(p1) // {Ann 40}
		fmt.Println(p2) // {Ann 99}

		Here, p2 := p1 copies the struct. Changing p2 does not change p1.

		But with a pointer:
		p1 := person{name: "Ann", age: 40}
		p2 := &p1

		p2.age = 99

		fmt.Println(p1) // {Ann 99}

		===
		Go automatically lets you write:
		p2.age = 99
		instead of forcing you to write:
		(*p2).age = 99
	*/

	// =========================================================//

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean

	sp := &s            // sp is a pointer to s
	fmt.Println(sp.age) // 50

	// either way works to set the age field of the struct that sp points to. Go automatically dereferences the pointer when you use the dot notation.
	sp.age = 51
	fmt.Println(s.age) // 51

	(*sp).age = 52
	fmt.Println(s.age) // 52

	// you can also create an anonymous struct which is a struct without a name. This is useful when you want to create a struct that is only used in one place and you don't want to define a new type for it.

	// this {}{} is a simple way to provide the values for the fields of the struct easily.
	// in this case we have "Rex" and true as the values for the name and isGood fields respectively. Order matters.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) // {Rex true}

}
