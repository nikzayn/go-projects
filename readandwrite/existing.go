package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	myData := []byte("Versace on the floor")

	err := ioutil.WriteFile("existing.data", myData, 0777)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("existing.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	f, err := os.OpenFile("existing.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString("new data is inserted to the file.\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("existing.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
