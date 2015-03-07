package main

import (
	"fmt"
	"math"
)

type cartesianPoint struct {
	x, y float64
}

type polarPoint struct {
	r, t float64
}

type Point interface {
	Printer // interface composition
	X() float64
	Y() float64
}

type Printer interface {
	Print()
}

func (c cartesianPoint) X() float64 {
	return c.x
}

func (c cartesianPoint) Y() float64 {
	return c.y
}

func (p polarPoint) X() float64 {
	return p.r * math.Cos(p.t)
}

func (p polarPoint) Y() float64 {
	return p.r * math.Sin(p.t)
}

func (self cartesianPoint) Print() {
	fmt.Printf("(x: %f, y: %f)\n",self.x,self.y)
}

func (self polarPoint) Print() {
	fmt.Printf("(r: %f, t: %f)\n", self.r,self.t)
}

func test(point Point) {
	point.Print()
}

func main() {
	cartesian := cartesianPoint{12.12,13.13}
	polar := polarPoint{21.21,31.31}

	cartesian.Print()
	polar.Print()

	test(cartesian)
	test(polar)
}
