package main

import "fmt"

// CurrentConditionsDisplay implement Observer & DisplayElement
type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	cd := &CurrentConditionsDisplay{}

	cd.weatherData = weatherData
	weatherData.RegisterObserver(cd)
	return cd
}

func (cd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	cd.temperature = temperature
	cd.humidity = humidity
	cd.Display()
}

func (cd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %f F degrees and %f 5 humidity\n", cd.temperature, cd.humidity)
}
