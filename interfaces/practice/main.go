package main

import (
	"fmt"
)

type blln struct {
	rate     int
	duration int
	balloon  int
}
type leas struct {
	rate     int
	duration int
}
type rateCalculator interface {
	calcRate() int
}

func (b blln) calcRate() int {
	return b.rate*b.duration + b.balloon
}
func (l leas) calcRate() int {
	return l.rate * l.duration
}

func totalRate(r []rateCalculator) {
	tot := 0
	for _, v := range r {
		tot = tot + v.calcRate()

	}
	fmt.Println("The total Rate for is: ", tot)

}
func main() {
	c1 := blln{
		rate:     4000,
		duration: 12,
		balloon:  1000,
	}
	c2 := blln{
		rate:     895,
		duration: 14,
		balloon:  3500,
	}
	c3 := leas{
		rate:     780,
		duration: 6,
	}
	sum := []rateCalculator{c1, c2, c3}

	totalRate(sum)

}
