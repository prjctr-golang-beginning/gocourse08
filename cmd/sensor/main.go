package main

import (
	"fmt"
	"time"
)

// GenericChannel - універсальний тип для каналу
type GenericChannel[T any] chan SensorData[T]

// TemperatureSensor - симуляція датчика температури
func ProcessTemperatureSensor(ch GenericChannel[float64], interval time.Duration) {
	var temperatureSensor Sensor[float64] = TemperatureSensor{}

	for {
		ch <- temperatureSensor.ReadData()
		time.Sleep(interval)
	}
}

// PulseSensor - симуляція датчика пульсу
func ProcessPulseSensor(ch GenericChannel[int], interval time.Duration) {
	var pulseSensor Sensor[int] = PulseSensor{}

	for {
		ch <- pulseSensor.ReadData()
		time.Sleep(interval)
	}
}

// SensorData - універсальний тип для даних датчика
type SensorData[T any] struct {
	Value     T
	Timestamp time.Time
}

// Sensor - універсальний інтерфейс для медичного датчика
type Sensor[T any] interface {
	ReadData() SensorData[T]
}

// TemperatureSensor - датчик температури
type TemperatureSensor struct{}

func (TemperatureSensor) ReadData() SensorData[float64] {
	// Логіка для читання даних з датчика температури
	return SensorData[float64]{Value: 36.6, Timestamp: time.Now()}
}

// PulseSensor - датчик пульсу
type PulseSensor struct{}

func (PulseSensor) ReadData() SensorData[int] {
	// Логіка для читання даних з датчика пульсу
	return SensorData[int]{Value: 72, Timestamp: time.Now()}
}

func main() {
	temperatureChannel := make(GenericChannel[float64])
	pulseChannel := make(GenericChannel[int])

	go ProcessTemperatureSensor(temperatureChannel, 2*time.Second)
	go ProcessPulseSensor(pulseChannel, 3*time.Second)

	// Читаємо дані з каналів у головній горутині
	for {
		select {
		case tempData := <-temperatureChannel:
			fmt.Printf("Temperature: %.2f at %v\n", tempData.Value, tempData.Timestamp)
		case pulseData := <-pulseChannel:
			fmt.Printf("Pulse: %d at %v\n", pulseData.Value, pulseData.Timestamp)
		}
	}
}
