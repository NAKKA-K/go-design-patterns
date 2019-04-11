package main

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}

type MuteQuack struct{}

func (mq *MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

type Squeak struct{}

func (sq *Squeak) Quack() {
	fmt.Println("Squeak")
}

type FakeQuack struct{}

func (fq *FakeQuack) Quack() {
	fmt.Println("Qwak")
}
