package main

import (
	"fmt"
	"os"
	"strings"
)

func sayHello(name string) string {
	return "Hello, " + name
}

func main() {
	who := "World!"

	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:]," ")
	}
	fmt.Printf("Greeting: %s\n",sayHello(who))
}
