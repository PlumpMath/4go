package main

import "fmt"

type Address struct {
	street string
	pincode int
}

type Employee struct {	
	name string
	address Address
	age int
}

type Customer struct {
	name string
	address Address
	age int
}

func (e *Employee) updateName(name string) {
	// pass by reference
	e.name = name
}

func (e Employee) updateN(name string) {
	// pass by value - copy - immutable
	sampleFunction()
	fmt.Printf("updateN: Before : [%s]\n",e.name)
	e.name=name
	fmt.Printf("updateN: After : [%s]\n",e.name)
}

func (c *Customer) updateName(name string) {
	c.name = name
}

func sampleFunction(){
	fmt.Printf("Sample Function - does nothing \n")
}


func main() {
	emp := Employee { "jitendra",
		Address { "goadbundher road", 400507}, 
		30}

	fmt.Printf("Employee: Before Update [%s]\n",emp.name)

	emp.updateName("Jitendra")

	fmt.Printf("Employee: After Update [%s]\n", emp.name)

	emp.updateN("Rajendra")

	fmt.Printf("Employee: After update [%s]\n", emp.name)
}

