package main

import (
	"fmt"
	"time"
)

const (
	//Max defines how many loops are going to run
	Max int = 5
	//MaxBigger defines a bigger number to loop throufh
	MaxBigger = 10000
	//Time defines how many seconds we will wait
	Time time.Duration = time.Duration(1)
)

func main() {
	fmt.Println("Hello")
}
