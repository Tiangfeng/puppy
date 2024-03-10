package puppy

import (
	"fmt"

	"github.com/Tiangfeng/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func From11() {
	fmt.Println("From 1.1.0")
}

func From12() {
	fmt.Println("From 1.2.0")
}

func From13() {
	fmt.Println("From 1.3.0")
}
