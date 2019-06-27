package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("localfile.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
