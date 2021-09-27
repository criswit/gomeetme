package main

import (
	"fmt"
	"gomeetme/factory"
)

const hoursInDay = 24

func main() {
	fmt.Print(factory.GenerateRandomTime())
}
