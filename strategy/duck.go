package main

import "fmt"

type Duck interface {
	PerformFly()
	PerformQuack()
	Swim()
	Display()
}

type duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *duck) setFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *duck) PerformFly() {
	d.flyBehavior.Fly()
}

func (d *duck) setQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

func (d *duck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d *duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

type MallardDuck struct {
	duck
}

func NewMallardDuck() Duck {
	duck := &MallardDuck{}
	duck.setFlyBehavior(&FlyWithWings{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (md *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard duck")
}

type RubberDuck struct {
	duck
}

func NewRubberDuck() Duck {
	duck := &RubberDuck{}
	duck.setFlyBehavior(&FlyNoWay{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (rd *RubberDuck) Display() {
	fmt.Println("I'm a rubber duckie")
}

type DecoyDuck struct {
	duck
}

func NewDecoyDuck() Duck {
	duck := &DecoyDuck{}
	duck.setFlyBehavior(&FlyNoWay{})
	duck.setQuackBehavior(&MuteQuack{})
	return duck
}

func (dd *DecoyDuck) Display() {
	fmt.Println("I'm a duck Decoy")
}

type ModelDuck struct {
	duck
}

func NewModelDuck() Duck {
	duck := &ModelDuck{}
	duck.setFlyBehavior(&FlyNoWay{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (md *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

type RedHeadDuck struct {
	duck
}

func NewRedHeadDuck() Duck {
	duck := &RedHeadDuck{}
	duck.setFlyBehavior(&FlyWithWings{})
	duck.setQuackBehavior(&Quack{})
	return duck
}

func (rd *RedHeadDuck) Display() {
	fmt.Println("I'm a real Red Headed duck")
}
