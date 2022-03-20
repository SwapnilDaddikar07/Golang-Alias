package main

import (
	"alias/Animals"
	"alias/Interfaces"
)

func main() {
	l := Animals.Lion{}
	t := Animals.Tiger{}

	//This will fail as soon as we change the interface to return the new error.
	test(l)
	test(t)
}

func test(animal Interfaces.Animal) {
	animal.MakeNoise()
}
