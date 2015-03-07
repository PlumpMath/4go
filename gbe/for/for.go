package main

import "fmt"

func main() {
	// With single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
