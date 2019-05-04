package main

import (
	"fmt"
	"math"
)
//Two benefits of using pointer receiver 
//1. In which the method can modify the value that its receiver points to
//2. To avoid copy of value on each method call. Can be more proficient to use on large struct

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Scale: %v\n", v ,v.Abs())
}