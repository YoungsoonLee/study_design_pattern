package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("goodsTrain: Landing")
	} else {
		fmt.Println("goodsTrain: Waiting")
	}
}

func (g *goodsTrain) departure() {
	fmt.Println("goodsTrain: Leaving")
	g.mediator.notifyFree()
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("goodsTrain: Arrival Permitted. Landing")
}
