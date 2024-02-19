package puppy

import (
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
