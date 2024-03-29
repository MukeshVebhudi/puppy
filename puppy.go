package puppy

import (
	"fmt"

	dog "github.com/MukeshVebhudi/Dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("Hi I am from version 1.1.0")
}

func From12() {
	fmt.Println("Hi I am from version 1.2.0")
}
