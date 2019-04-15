package main

import "fmt"

// StatisticsDisplay implement Observer & DisplayElement
type StatisticsDisplay struct {
	maxTemp     float64
	minTemp     float64
	tempSum     float64
	numReadings int64
	weatherData *WeatherData
}

func NewStatisticsDisplay(weatherData *WeatherData) *StatisticsDisplay {
	sd := &StatisticsDisplay{}
	sd.weatherData = weatherData

	sd.weatherData.RegisterObserver(sd)
	return sd
}

func (sd *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
	sd.tempSum = sd.tempSum + temperature
	sd.numReadings = sd.numReadings + 1

	if temperature > sd.maxTemp {
		sd.maxTemp = temperature
	}
	if temperature < sd.minTemp {
		sd.minTemp = temperature
	}

	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	fmt.Printf("Avg/Max/Min temperature = %f/%f/%f\n", (sd.tempSum / float64(sd.numReadings)), sd.maxTemp, sd.minTemp)
}
