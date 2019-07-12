package main

import "fmt"

type Meower interface {
	meow()
}

type Animal struct {
	size int
}

func (*Animal) domesticate() {
	fmt.Println("New pet!")
}

type Cat struct {
	// Note that the interfaces are no where in the type definition
	animal       Animal
	meowStrength int
}

func (*Cat) meow() {
	fmt.Println("Meow!")
}

type Tiger struct {
	cat                   Cat
	wildingHuntingAbility int
}

func (*Tiger) meow() {
	fmt.Println("Rawr!")
}

func pet(meower Meower) {
	meower.meow()
}

func main() {
	tigerPtr := &Tiger{}
	pet(tigerPtr)

	tigerVal := *tigerPtr // Take value of pointer
	pet(tigerVal)         // Invalid

	tigerVal.cat.animal.domesticate() // Gets verbose with a large amount of composition

	fmt.Println(tigerVal.cat.meowStrength)
}
