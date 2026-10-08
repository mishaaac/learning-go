package main

import "fmt"

func main() {
	jobStatus := "pending"
	result := "processed"

	if jobStatus == "pending" {
		result = "skipped"
	}

	fmt.Println(jobStatus, result)
}
