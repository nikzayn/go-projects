package main

import "fmt"

type I Interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() { 
	fmt.Println(t.S)
}

type F float64

func (f F) M(){
	
}