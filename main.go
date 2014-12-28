package main

import (
	"fmt"
	"github.com/jittakal/4go/aws3"
	"github.com/jittakal/4go/easy"
	"github.com/jittakal/4go/gplp"
)

func main() {
	// Hello 4go
	fmt.Println("Hello 4go.")

	// Easy
	fmt.Println(easy.SayHello("Jitendra"))

	// GPLP
	gplp.PathMain()

	// AWS3
	aws3.ListBuckets()
}
