package main

import "fmt"

// ForecastDisplay implement Observer & DisplayElement
type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64
	weatherData     *WeatherData
}

func NewForecastDisplay(weatherData *WeatherData) *ForecastDisplay {
	fd := &ForecastDisplay{}
	fd.currentPressure = 29.92

	fd.weatherData = weatherData
	weatherData.RegisterObserver(fd)
	return fd
}

func (fd *ForecastDisplay) Update(temperature, humidity, pressure float64) {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = pressure

	fd.Display()
}

func (fd *ForecastDisplay) Display() {
	fmt.Print("Forecast: ")
	if fd.currentPressure > fd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fd.currentPressure == fd.lastPressure {
		fmt.Println("More of the same")
	} else if fd.currentPressure < fd.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather")
	}
}
