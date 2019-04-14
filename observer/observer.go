package main

type Observer interface {
	Update(tmp, humidity, pressure float64)
}
