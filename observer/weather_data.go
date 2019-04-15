package main

// WeatherData implement Subject
type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	wd := &WeatherData{}

	wd.observers = []Observer{}
	return wd
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	observers := make([]Observer, len(wd.observers)-1, len(wd.observers)-1)
	for _, obs := range wd.observers {
		if obs != o {
			observers = append(observers, obs)
		}
	}
	wd.observers = observers
}

func (wd *WeatherData) NotifyObservers() {
	for _, obs := range wd.observers {
		obs.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.MeasurementsChanged()
}

func (wd *WeatherData) GetTemperature() float64 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float64 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float64 {
	return wd.pressure
}
