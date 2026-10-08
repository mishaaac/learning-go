package main

import "fmt"

func main() {
	stations := []struct {
		Name     string
		Readings []int
	}{
		{Name: "North", Readings: []int{12, 0, 38}},
		{Name: "East", Readings: []int{-1, 45, 52}},
		{Name: "South", Readings: []int{67, 91, 20}},
		{Name: "West", Readings: []int{22, 44}},
	}

	completedStations := 0
	rejectedStations := 0
	lowCount := 0
	moderateCount := 0
	highCount := 0

auditLoop:
	for _, station := range stations {
		stationReadingCount := 0

		for _, reading := range station.Readings {
			if reading == 0 {
				continue
			}

			if reading < 0 {
				rejectedStations++
				fmt.Printf("Station %s: rejected at reading %d\n", station.Name, reading)
				continue auditLoop
			}

			if reading >= 90 {
				fmt.Printf("Emergency: station=%s reading=%d\n", station.Name, reading)
				break auditLoop
			}

			tier := ""
			switch {
			case reading < 30:
				tier = "low"
				lowCount++
			case reading < 60:
				tier = "moderate"
				moderateCount++
			default:
				tier = "high"
				highCount++
			}

			stationReadingCount++
			fmt.Printf("Station %s: reading=%d tier=%s\n", station.Name, reading, tier)
		}

		completedStations++
		fmt.Printf("Station %s: completed readings=%d\n", station.Name, stationReadingCount)
	}

	fmt.Printf("Station totals: completed=%d rejected=%d\n", completedStations, rejectedStations)
	fmt.Printf("Tier totals: low=%d moderate=%d high=%d\n", lowCount, moderateCount, highCount)

	// Los totales del audit se declaran fuera de ambos loops para que sigan disponibles al final.
	// stationReadingCount pertenece solo a la iteración de la estación actual.
}
