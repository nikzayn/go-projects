package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	myData := []byte("Nobody can drag me down.\n")

	err := ioutil.WriteFile("file.data", myData, 0777)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("file.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
