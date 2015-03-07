package main

import (
	"bytes"
	"fmt"
)

func main() {
	bt := []byte{'a','b','c'}

	hw := []byte(" Hello, goodbye, etc! ")
	
	bfr := bytes.NewBuffer(bt)

	fmt.Printf("ToUpper %s\n",string(bytes.ToUpper(bt)))
	fmt.Printf("TrimSapce '%s'\n", string(bytes.TrimSpace(hw)))
	fmt.Printf("Length %d\n",bfr.Len())


}
