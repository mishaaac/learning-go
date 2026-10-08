package main

import "fmt"

type SensorReading struct {
	Sensor string
	Value  int
}

func main() {
	incoming := [3]SensorReading{
		{Sensor: "alpha", Value: 12},
		{Sensor: "beta", Value: 0},
		{Sensor: "gamma", Value: 27},
	}
	accepted := make([]SensorReading, len(incoming))

	for index, reading := range incoming {
		accepted[index] = reading
	}

	fmt.Println("Accepted:", len(accepted))
	fmt.Println(accepted)
}
