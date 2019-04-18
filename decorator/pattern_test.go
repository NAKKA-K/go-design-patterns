package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	beverage := interface{}(NewEspresso()).(Beverage)
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())

	beverage2 := interface{}(NewDarkRoast()).(Beverage)
	beverage2 = interface{}(NewMocha(beverage2)).(Beverage)
	beverage2 = interface{}(NewMocha(beverage2)).(Beverage)
	beverage2 = interface{}(NewWhip(beverage2)).(Beverage)
	fmt.Printf("%s $%.2f\n", beverage2.GetDescription(), beverage2.Cost())

	beverage3 := interface{}(NewHouseBland()).(Beverage)
	beverage3.SetSize(VENTI)
	beverage3 = interface{}(NewSoy(beverage3)).(Beverage)
	beverage3 = interface{}(NewMocha(beverage3)).(Beverage)
	beverage3 = interface{}(NewWhip(beverage3)).(Beverage)
	fmt.Printf("%s $%.2f\n", beverage3.GetDescription(), beverage3.Cost())
}
