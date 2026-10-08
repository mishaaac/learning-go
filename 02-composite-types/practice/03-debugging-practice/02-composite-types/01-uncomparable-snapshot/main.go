package main

import "fmt"

type DeploymentSnapshot struct {
	Environment string
	Services    []string
}

func main() {
	previous := DeploymentSnapshot{
		Environment: "production",
		Services:    []string{"api", "worker", "scheduler"},
	}
	current := DeploymentSnapshot{
		Environment: "production",
		Services:    []string{"api", "worker", "scheduler"},
	}

	environmentsMatch := previous.Environment == current.Environment
	servicesMatch := [3]string(previous.Services) == [3]string(current.Services)
	snapshotsMatch := environmentsMatch && servicesMatch

	fmt.Printf("Snapshots match: %t\n", snapshotsMatch)
}
