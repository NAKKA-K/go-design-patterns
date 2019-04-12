package main

import (
	"fmt"
	"testing"
)

func TestStrategyPattern(t *testing.T) {
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
