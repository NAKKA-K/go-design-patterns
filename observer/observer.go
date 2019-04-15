package main

type Observer interface {
	Update(temperature, humidity, pressure float64)
}
