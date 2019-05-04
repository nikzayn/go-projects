package main

import "fmt"

type person struct {
	fname string
	lname string
}

type innovater struct {
	person
	chanceToCreate bool
}

type human interface {
	speaks()
}

func (p person) speaks() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning, James". `)
}

func (io innovater) speaks() {
	fmt.Println(io.fname, io.lname, `inovates, "Photosynthesis into reality". `)
}

func saySomething(hu human) {
	hu.speaks()
}

func main() {
	xi := []int{2, 3, 4, 5, 6}
	fmt.Println(xi)

	m := map[string]int{
		"Nik":    22,
		"Shanki": 18,
	}
	fmt.Println(m)

	p := person{"Nikhil", "Vaidyar"}

	io := innovater{
		person{"ALbert",
			"Einstien",
		}, true,
	}

	saySomething(p)
	saySomething(io)

}
