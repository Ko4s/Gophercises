package main

import (
	"blackjack/deck"
	"fmt"
)

func main() {
	d := deck.New()
	fmt.Println(d[0])
}
