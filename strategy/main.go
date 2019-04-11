package main

import "fmt"

func main() {
	ducks := []Duck{
		NewMallardDuck(),
		NewRubberDuck(),
		NewDecoyDuck(),
		NewModelDuck(),
		NewRedHeadDuck(),
	}

	for i, duck := range ducks {
		fmt.Printf("[%d duck] =====\n", i+1)
		duck.Display()
		duck.Swim()
		duck.PerformFly()
		duck.PerformQuack()
	}
}
