package main

import "fmt"

// HeatIndexDisplay implement Observer & DisplayElement
type HeatIndexDisplay struct {
	heatIndex   float64
	weatherData *WeatherData
}

func NewHeatIndexDisplay(weatherData *WeatherData) *HeatIndexDisplay {
	hd := &HeatIndexDisplay{}
	hd.weatherData = weatherData

	weatherData.RegisterObserver(hd)
	return hd
}

func (hd *HeatIndexDisplay) Update(temperature, humidity, pressure float64) {
	hd.heatIndex = hd.computeHeatIndex(temperature, humidity)
	hd.Display()
}

func (hd *HeatIndexDisplay) computeHeatIndex(t, rh float64) float64 {
	index := ((16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) +
		(0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) +
		(0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) +
		(0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 *
		(rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) +
		(0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) +
		0.000000000843296*(t*t*rh*rh*rh)) -
		(0.0000000000481975 * (t * t * t * rh * rh * rh)))
	return index
}

func (hd *HeatIndexDisplay) Display() {
	fmt.Printf("Heat index is %f \n", hd.heatIndex)
}
