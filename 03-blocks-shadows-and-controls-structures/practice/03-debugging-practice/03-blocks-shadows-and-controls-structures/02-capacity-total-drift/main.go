package main

import "fmt"

func main() {
	capacity := 10
	approved := 0
	requests := []int{4, 7, 3}

	for _, request := range requests {
		if approved+request > capacity {
			fmt.Println("rejected", request)
			continue
		}

		approved = approved + request
		fmt.Println("accepted", request, approved)
	}

	fmt.Println("total", approved)
}
