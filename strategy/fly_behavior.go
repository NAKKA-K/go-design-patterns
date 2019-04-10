package main

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (fw *FlyWithWings) Fly() {
	fmt.Println("I'm flying!!")
}

type FlyNoWay struct{}

func (fn *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

type FlyRocketPowered struct{}

func (fr *FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket")
}
