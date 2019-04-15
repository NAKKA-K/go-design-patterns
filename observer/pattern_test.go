package main

import "testing"

func TestObserver(t *testing.T) {
	weatherData := NewWeatherData()

	data := interface{}(weatherData).(Subject)

	_ = NewCurrentConditionsDisplay(data)
	_ = NewStatisticsDisplay(weatherData)
	_ = NewForecastDisplay(weatherData)
	_ = NewHeatIndexDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
