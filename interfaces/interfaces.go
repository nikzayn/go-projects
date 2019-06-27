package main

import "fmt"

type Guitarist interface {
	playGuitar()
}

type BassGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BassGuitarist) playGuitar() {
	fmt.Println("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) playGuitar() {
	fmt.Println("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	var player BassGuitarist
	player.Name = "Nikhil"
	player.playGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Danial"
	player2.playGuitar()
}
