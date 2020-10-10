// +build arduino

package main

import (
	"machine"
	"time"
)

const pomp = machine.D2
const lamp = machine.D3

var analogInPin = machine.ADC{Pin: machine.ADC3}
var analogInPinTemp = machine.ADC{Pin: machine.ADC2}
var fotoSensorPin = machine.ADC{Pin: machine.ADC5}

var (
	temperature  = uint16(0)
	humidity     = uint16(0)
	illumination = uint16(0)
	minHumidity  = uint16(200)
	minLighting  = uint16(3300)
)

func init() {
	machine.InitADC()

	lamp.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pomp.Configure(machine.PinConfig{Mode: machine.PinOutput})

	analogInPin.Configure()
	analogInPinTemp.Configure()
	fotoSensorPin.Configure()
}

func main() {
	for {
		temperature = analogInPin.Get()
		humidity = analogInPinTemp.Get()
		illumination = fotoSensorPin.Get()

		println("temp = ", temperature)
		println("illumination = ", illumination)
		println("humidity = ", humidity)

		if illumination < minLighting {
			lamp.High()
		} else {
			lamp.Low()
		}

		if temperature < minHumidity {
			pomp.High()
		} else {
			pomp.Low()
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
