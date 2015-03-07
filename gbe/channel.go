package gbe

import "fmt"

func main() {

	fmt.Println("Jitendra Takalkar")

	go func(name string) {
		fmt.Println("From inside anonymous function using goroutine")
	}("Jitendra")

	fmt.Println("Printing from main function")
}
